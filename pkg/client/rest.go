package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kv/pkg/rest"
	"kv/pkg/watch"
	"log"
	"net/http"
	"net/url"
)

type RestClient struct {
	url    string
	client http.Client
}

func NewRest(host string) KV {
	return &RestClient{
		url:    fmt.Sprintf("http://%s", host),
		client: http.Client{},
	}
}

func (kv *RestClient) Get(ctx context.Context, key string) (interface{}, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", makeKeyUrl(kv.url, key), nil)
	if err != nil {
		return nil, err
	}

	resp, err := kv.client.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	update := rest.GetResponse{}
	err = json.Unmarshal(respBytes, &update)
	return update.Value, err
}

func (kv *RestClient) Put(ctx context.Context, key string, val interface{}) error {
	b, err := json.Marshal(rest.PutRequest{Key: key, Value: val})
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(b)

	req, err := http.NewRequestWithContext(ctx, "POST", makeKeyUrl(kv.url, key), body)
	if err != nil {
		return err
	}

	resp, err := kv.client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}

func (kv *RestClient) Delete(ctx context.Context, key string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", makeKeyUrl(kv.url, key), nil)
	if err != nil {
		return err
	}

	resp, err := kv.client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}

func (kv *RestClient) Watch(ctx context.Context, key string, operation watch.Operation) (chan watch.Update, error) {
	watchReq := watch.WatchRequest{
		Key:       key,
		WatchType: operation,
	}

	b, err := json.Marshal(watchReq)
	if err != nil {
		return nil, err
	}

	watchClient := http.Client{
		Timeout: 0, // no timeout
	}

	watchChan := make(chan watch.Update)

	go func() {
		defer close(watchChan)
		for {
			bytesBuff := bytes.NewBuffer(b)
			req, err := http.NewRequestWithContext(ctx, "POST", makeWatchUrl(kv.url), bytesBuff)
			if err != nil {
				break
			}

			resp, err := watchClient.Do(req)
			if err != nil {
				log.Printf("watch request failed: %+v\n", err)
				return
			} else if resp.StatusCode != 200 {
				log.Printf("watch request failed: %s\n", resp.Status)
				return
			}

			respBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				break
			}
			resp.Body.Close()

			update := watch.Update{}
			err = json.Unmarshal(respBytes, &update)
			if err != nil {
				log.Printf("decoding json failed: %+v\n", err)
				break
			}
			watchChan <- update
		}
	}()

	return watchChan, nil
}

func makeKeyUrl(requestUrl, key string) string {
	return requestUrl + "/kv/" + url.PathEscape(key)
}

func makeWatchUrl(requestUrl string) string {
	return requestUrl + "/watch"
}

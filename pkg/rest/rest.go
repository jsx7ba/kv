package rest

type PutRequest struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type GetResponse struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type GetRequest struct {
	Key string `json:"key"`
}

type DeleteRequest struct {
	Key string `json:"key"`
}

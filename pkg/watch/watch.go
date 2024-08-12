package watch

import (
	"errors"
	"kv/internal/gen"
)

type Operation int

const (
	All Operation = iota + 1
	Put
	Delete
)

type Update struct {
	Op    Operation   `json:"op"`
	Key   string      `json:"key"`
	Value interface{} `json:"value,omitempty"`
}

func OperationFromString(name string) (op Operation, err error) {
	switch name {
	case "put":
		op = Put
	case "delete":
		op = Delete
	case "all":
		op = All
	default:
		err = errors.New("unknown watch type: " + name)
	}
	return
}

func OperationFrom(optype gen.OpType) Operation {
	var o Operation
	switch optype {
	case gen.OpType_PUT:
		o = Put
	case gen.OpType_DELETE:
		o = Delete
	default:
		panic("unknown gen.OpType type")
	}
	return o
}

func (o Operation) Convert() gen.OpType {
	var optype gen.OpType
	switch o {
	case Put:
		optype = gen.OpType_PUT
	case Delete:
		optype = gen.OpType_DELETE
	case All:
		optype = gen.OpType_ALL
	default:
		panic("unknown Operation type")
	}
	return optype
}

type WatchRequest struct {
	Key       string    `json:"key"`
	WatchType Operation `json:"watchType"`
}

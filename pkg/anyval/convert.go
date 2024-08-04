package anyval

import (
	"errors"
	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"kv/internal/gen"
	"reflect"
)

// ConvertableType defines all the types the key value store will work with.
type ConvertableType interface {
	constraints.Ordered | []string | string |
		~[]int | ~[]int64 | ~[]int32 | ~[]int16 | ~[]int8 |
		~[]float32 | ~[]float64
}

// Marshal convert interface{} values to an anypb.Any value.
func Marshal(val interface{}) (*anypb.Any, error) {
	var m proto.Message

	switch v := val.(type) {
	case string:
		m = wrapperspb.String(v)
	case float32:
		m = wrapperspb.Float(v)
	case float64:
		m = wrapperspb.Double(v)
	case int32:
		m = wrapperspb.Int32(v)
	case int16:
		m = wrapperspb.Int32(int32(v))
	case int8:
		m = wrapperspb.Int32(int32(v))
	case uint32:
		m = wrapperspb.UInt32(v)
	case uint16:
		m = wrapperspb.UInt32(uint32(v))
	case uint8:
		m = wrapperspb.UInt32(uint32(v))
	case int:
		m = wrapperspb.Int64(int64(v))
	case int64:
		m = wrapperspb.Int64(v)
	case uint:
		m = wrapperspb.UInt64(uint64(v))
	case uint64:
		m = wrapperspb.UInt64(v)
	case bool:
		m = wrapperspb.Bool(v)
	case []byte:
		m = wrapperspb.Bytes(v)
	case []string:
		m = &gen.StringSliceWrapper{Value: v}
	case []int32:
		m = &gen.Int32SliceWrapper{Value: v}
	case []int64:
		m = &gen.Int64SliceWrapper{Value: v}
	case []float32:
		m = &gen.Float32SliceWrapper{Value: v}
	case []float64:
		m = &gen.Float64SliceWrapper{Value: v}
	default:
		return nil, errors.New("unrecognized type: " + reflect.TypeOf(val).Name())
	}

	return anypb.New(m)
}

// UnmarshalType converts anypb.Any values to the value requested by the caller.
func UnmarshalType[T ConvertableType](anyVal *anypb.Any) (T, error) {
	v, err := unmarshal(anyVal)
	var tempVal T
	if err != nil {
		return tempVal, err
	}

	// This block casts to the type the caller asks for.  Al little extra work is need as protocol buffers don't
	// support the smaller types and native sized types.
	// int and uint are encoded as int64 and uint64 respectively.
	// int32, int16, int8 are encoded as int32.
	// uint32, uint16, uint8 are encoded as uint32
	var dynamicVal interface{} = tempVal

	switch dynamicVal.(type) {
	case int:
		var x = v.(int64)
		dynamicVal = int(x)
	case int16:
		var x = v.(int32)
		dynamicVal = int16(x)
	case int8:
		var x = v.(int32)
		dynamicVal = int8(x)
	case uint:
		var x = v.(uint64)
		dynamicVal = uint(x)
	case uint16:
		var x = v.(uint32)
		dynamicVal = uint16(x)
	case uint8:
		var x = v.(uint32)
		dynamicVal = uint8(x)
	default:
		var ok bool
		dynamicVal, ok = v.(T)
		if !ok {
			return tempVal, errors.New("cannot cast to message value to " + reflect.TypeOf(v).Name())
		}
	}

	return dynamicVal.(T), nil
}

func unmarshal(anyVal *anypb.Any) (interface{}, error) {
	m, err := anyVal.UnmarshalNew()
	if err != nil {
		return nil, err
	}

	var v interface{}
	switch m := m.(type) {
	case *wrapperspb.StringValue:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *wrapperspb.BytesValue:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *wrapperspb.UInt64Value:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *wrapperspb.UInt32Value:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *wrapperspb.BoolValue:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *wrapperspb.FloatValue:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *wrapperspb.DoubleValue:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *wrapperspb.Int64Value:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *wrapperspb.Int32Value:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *gen.StringSliceWrapper:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *gen.Int32SliceWrapper:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *gen.Int64SliceWrapper:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *gen.Float32SliceWrapper:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	case *gen.Float64SliceWrapper:
		err = anyVal.UnmarshalTo(m)
		v = m.GetValue()
	default:
		err = errors.New("unrecognized type: " + anyVal.TypeUrl)
	}
	return v, err
}

// Unmarshal converts anypb.Any values to interface{}.
func Unmarshal(anyVal *anypb.Any) (interface{}, error) {
	return unmarshal(anyVal)
}

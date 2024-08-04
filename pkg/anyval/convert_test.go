package anyval

import (
	"reflect"
	"testing"
)

func TestMarshalInt(t *testing.T) {
	expected := 42
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[int](anyVal)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected [%d] got [%d]", expected, actual)
	}
}

func TestConvertInt32(t *testing.T) {
	expected := int32(42)
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[int32](anyVal)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected [%d] got [%d]", expected, actual)
	}
}

func TestConvertUInt32(t *testing.T) {
	expected := uint32(42)
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[uint32](anyVal)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected [%d] got [%d]", expected, actual)
	}
}

func TestConvertUint8(t *testing.T) {
	expected := uint8(42)
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[uint8](anyVal)
	if err != nil {
		t.Fatal(err)
	}
	if expected != actual {
		t.Fatalf("expected [%d] got [%d]", expected, actual)
	}
}

func TestConvertString(t *testing.T) {
	expected := "test\nvalue"
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[string](anyVal)

	if actual != expected {
		t.Fatalf("expected [%s] got [%s]", expected, actual)
	}
}

func TestConvertStringSlice(t *testing.T) {
	expected := []string{"test", "value", "for", "testing"}
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[[]string](anyVal)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected [%s] got [%s]", expected, actual)
	}
}

func TestConvertInt32Slice(t *testing.T) {
	expected := []int32{0, 1, 1, 2, 3, 5, 8, 13}
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[[]int32](anyVal)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v got %+v\n", expected, actual)
	}
}

func TestConvertInt64Slice(t *testing.T) {
	expected := []int64{0, 1, 1, 2, 3, 5, 8, 13}
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[[]int64](anyVal)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v got %+v\n", expected, actual)
	}
}

func TestConvertFloat32Slice(t *testing.T) {
	expected := []float32{0.0, 1.0, 1.618, 3.141596}
	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[[]float32](anyVal)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v got %+v\n", expected, actual)
	}
}

func TestConvertFloat64Slice(t *testing.T) {
	expected := []float64{0.0, 1.0, 1.618, 3.141596}

	anyVal, err := Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := UnmarshalType[[]float64](anyVal)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v got %+v\n", expected, actual)
	}
}

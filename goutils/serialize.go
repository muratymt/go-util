package goutils

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

type SerializedCopy struct {
	data     []byte
	dataType reflect.Type
}

func Value2SerializedCopy(v interface{}) (*SerializedCopy, error) {
	data, err := Serialize(v)
	if err != nil {
		return nil, err
	}

	result := &SerializedCopy{
		data:     *data,
		dataType: reflect.TypeOf(v),
	}
	return result, nil
}

func SerializedCopy2Value(serializedCopy *SerializedCopy) (interface{}, error) {
	newValue := reflect.New(serializedCopy.dataType)
	err := Deserialize(&serializedCopy.data, newValue.Interface())
	if err != nil {
		return nil, err
	}
	return newValue.Elem().Interface(), nil
}

func Serialize(value interface{}) (*[]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return nil, err
	}

	result := buf.Bytes()

	return &result, nil
}

func Deserialize(data *[]byte, out interface{}) error {
	var buf bytes.Buffer
	buf.Write(*data)
	dec := gob.NewDecoder(&buf)
	return dec.Decode(out)
}

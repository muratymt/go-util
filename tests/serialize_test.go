package tests

import (
	"github.com/muratymt/go-utils/goutils/goutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestSerealizeStructIn struct {
	A int
	B *string
}

type TestSerealizeStruct struct {
	X *TestSerealizeStructIn
}

func TestSerializeInt(t *testing.T) {
	originalValue := int64(24)

	serializedCopy, err := goutils.Value2SerializedCopy(originalValue)

	assert.Nil(t, err)
	var copyValue interface{}

	copyValue, err = goutils.SerializedCopy2Value(serializedCopy)

	assert.Equal(t, originalValue, copyValue.(int64))
}

func TestSerializeString(t *testing.T) {
	originalValue := "Foreign Function Interface"

	serializedCopy, err := goutils.Value2SerializedCopy(originalValue)

	assert.Nil(t, err)
	var copyValue interface{}

	copyValue, err = goutils.SerializedCopy2Value(serializedCopy)

	assert.Equal(t, originalValue, copyValue.(string))
}

func TestSerializeStruct(t *testing.T) {
	a := 4567
	b := "Foreign Function Interface"

	structIn := TestSerealizeStructIn{
		A: a,
		B: &b,
	}

	structOut := TestSerealizeStruct{X: &structIn}

	serializedCopy, err := goutils.Value2SerializedCopy(structOut)

	assert.Nil(t, err)
	var copyValue interface{}

	copyValue, err = goutils.SerializedCopy2Value(serializedCopy)

	assert.Equal(t, structOut, copyValue.(TestSerealizeStruct))
}

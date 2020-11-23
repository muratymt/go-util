package tests

import (
	"github.com/muratymt/go-utils/goutils/goutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunkStr(t *testing.T) {
	data := []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7"}

	res := *goutils.ChunkStr(&data, 0, 3)

	assert.Len(t, res, 3)
	assert.Contains(t, res, "key1")
	assert.Contains(t, res, "key2")
	assert.Contains(t, res, "key3")

	res = *goutils.ChunkStr(&data, 3, 3)

	assert.Len(t, res, 3)
	assert.Contains(t, res, "key4")
	assert.Contains(t, res, "key5")
	assert.Contains(t, res, "key6")

	res = *goutils.ChunkStr(&data, 4, 3)

	assert.Len(t, res, 3)
	assert.Contains(t, res, "key5")
	assert.Contains(t, res, "key6")
	assert.Contains(t, res, "key7")

	res = *goutils.ChunkStr(&data, 5, 3)

	assert.Len(t, res, 2)
	assert.Contains(t, res, "key6")
	assert.Contains(t, res, "key7")

	res = *goutils.ChunkStr(&data, 7, 3)

	assert.Len(t, res, 0)

	res = *goutils.ChunkStr(&data, 10, 3)

	assert.Len(t, res, 0)
}

func TestDiffStr(t *testing.T) {
	data1 := []string{"key1", "key2", "key4", "key7"}
	data2 := []string{"key2", "key3", "key5"}

	diff := *goutils.DiffStr(&data1, &data2)

	assert.Len(t, diff, 3)
	assert.Contains(t, diff, "key1")
	assert.Contains(t, diff, "key4")
	assert.Contains(t, diff, "key7")
}

package tests

import (
	"github.com/muratymt/go-utils/goutils/goutils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGenerateRandomString(t *testing.T) {
	length := 1000
	charset := []rune("ABCD")
	rndString := goutils.GenerateRandomString(uint(length), &charset)

	assert.Equal(t, length, utf8.RuneCountInString(*rndString))
	assert.True(t, strings.Contains(*rndString, "A"))
	assert.True(t, strings.Contains(*rndString, "B"))
	assert.True(t, strings.Contains(*rndString, "C"))
	assert.True(t, strings.Contains(*rndString, "D"))
	assert.False(t, strings.Contains(*rndString, "E"))
}

package goutils

import "math/rand"

func GenerateRandomString(length uint, charset *[]rune) *string {
	byteResult := make([]int32, length)
	charsetLen := len(*charset)
	for i := range byteResult {
		byteResult[i] = (*charset)[rand.Intn(charsetLen)]
	}
	result := string(byteResult)

	return &result
}

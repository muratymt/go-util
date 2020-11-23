package goutils

func DiffStr(a, b *[]string) (diff *[]string) {
	mb := make(map[string]byte, len(*b))

	for _, value := range *b {
		mb[value] = 1
	}

	result := make([]string, 0)

	diff = &result

	for _, value := range *a {
		if _, ok := mb[value]; !ok {
			result = append(result, value)
		}
	}

	return
}

func ChunkStr(slice *[]string, offset int, count int) *[]string {
	keyCount := len(*slice)

	if offset > keyCount {
		offset = keyCount
	}

	var keyMax int
	if offset+count >= keyCount {
		keyMax = keyCount
	} else {
		keyMax = offset + count
	}

	result := (*slice)[offset:keyMax]

	return &result
}

package util

// Looks for the first occurence of the match bytes in order.
// Includes the matching bytes.
func BytesUntilMatch(bytes []byte, match []byte) []byte {
	le := len(match)
	for i := 0; i < len(bytes)-le; i++ {
		found := true
		for j, m := range match {
			if bytes[i+j] != m {
				found = false
			}
		}
		if found {
			return bytes[:i+le]
		}
	}
	return bytes
}

// Multiplies a byteslice into an integer. Expects slice to be
// ordered from least significant byte to most significant.
func ByteSliceMultiplication(bytes []byte, byteMaxValue int) int {
	integer := 0
	multiplier := 1
	for _, bit := range bytes {
		integer += int(bit) * multiplier
		multiplier *= byteMaxValue
	}
	return integer
}

func ReverseByteSlice(bytes []byte) []byte {
	rev := bytes
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return rev
}

func CompareByteSlices(b1 []byte, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i, a := range b1 {
		if a != b2[i] {
			return false
		}
	}
	return true
}

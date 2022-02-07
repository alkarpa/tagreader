package id3

import "github.com/alkarpa/tagreader/util"

var encoding_ISO byte = 0
var encoding_Unicode byte = 1

// ID3v2 frame specs:
// ISO strings are terminated with \x00
// Unicode strings are terminated with \x00\x00
func bytesUntilEncodingEnd(encoding byte, bytes []byte) []byte {
	var terminator []byte
	switch encoding {
	case encoding_ISO:
		terminator = []byte{0}
	case encoding_Unicode:
		terminator = []byte{0, 0}
	}
	return util.BytesUntilMatch(bytes, terminator)
}

// ID3v2 frame specs:
// Unicode strings start with the Unicode BOM (\xFF\xFE or \xFE\xFF)
func decodeStringBytes(encoding byte, bytes []byte) string {
	switch encoding {
	case encoding_Unicode:
		return string(bytes[2:])
	default:
		return string(bytes)
	}
}

func id3HeaderByteSize(bytes []byte) int {
	id3_header_sizebytes_max_byte_value := 128
	rev := util.ReverseByteSlice(bytes)
	return util.ByteSliceMultiplication(rev, id3_header_sizebytes_max_byte_value)
}
func id3FrameByteSize(bytes []byte) int {
	id3_frame_sizebytes_max_byte_value := 256
	rev := util.ReverseByteSlice(bytes)
	return util.ByteSliceMultiplication(rev, id3_frame_sizebytes_max_byte_value)
}

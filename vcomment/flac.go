package vcomment

import (
	"os"

	"github.com/alkarpa/tagreader/util"
)

// Reads metadata blocks from the file in search for the comments block
// identified by the 0x04 byte. Also determines whether further metadatablock
// reads are necessary.
// Returns comment byte data when found and a keep-reading-metadatablock bool
func readMetadatablock(f *os.File) ([]byte, bool) {
	dat := make([]byte, 4)
	f.Read(dat)

	block_header := dat[0]

	sevbit := block_header & (0xff >> 1)

	lengthbytes := util.ReverseByteSlice(dat[1:4])
	blocklength := util.ByteSliceMultiplication(lengthbytes, 256)
	if blocklength == 0 {
		return nil, false
	}

	if sevbit == 0x04 {
		comment := make([]byte, blocklength)
		f.Read(comment)
		return comment, false
	}
	f.Seek(int64(blocklength), 1)

	lastblock := (block_header&byte(128) == 0x01)

	return nil, !lastblock
}

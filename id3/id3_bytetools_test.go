package id3

import "testing"

func TestId3ByteSizes(t *testing.T) {

	t.Run("Header size, 7 bit zero (128), $00 00 02 01 = 257", func(t *testing.T) {
		tagsize := []byte{0x00, 0x00, 0x02, 0x01}
		received := id3HeaderByteSize(tagsize) // id3HeaderByteSize(tagsize, 128)
		expected := 257
		if received != expected {
			t.Fatalf("Expected %d, received %d", expected, received)
		}
	})

	t.Run("Frame size, full 8 bit (256), $00 00 02 01 = 513", func(t *testing.T) {
		tagsize := []byte{0x00, 0x00, 0x02, 0x01}
		received := id3FrameByteSize(tagsize) // id3HeaderByteSize(tagsize, 256)
		expected := 513
		if received != expected {
			t.Fatalf("Expected %d, received %d", expected, received)
		}
	})
}

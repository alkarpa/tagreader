package audiofile

import "testing"

func TestOpenFile(t *testing.T) {

	t.Run("No file at path", func(t *testing.T) {
		_, err := OpenAndReadFile("../../testdata/no_such_file")
		if err == nil {
			t.Fatalf("Opened a file that doesn't exist")
		}
	})

	t.Run("Empty file", func(t *testing.T) {
		_, err := OpenAndReadFile("../../testdata/empty_file")
		if err == nil {
			t.Fatalf("Read bytes from a file that is empty")
		}
	})

	t.Run("Unknown identifier", func(t *testing.T) {
		tags, err := OpenAndReadFile("../../testdata/unknown_identifier")
		if err != nil {
			t.Fatalf(err.Error())
		}
		if len(tags) > 0 {
			t.Fatalf("There should be nothing to read")
		}
	})

	t.Run("testdata/id3v2_3.mp3", func(t *testing.T) {
		ti, err := OpenAndReadFile("../../testdata/id3v2_3.mp3")
		if err != nil {
			t.Fatalf("Error opening file")
		}
		if ti["Album"] != "Test Files" {
			t.Fatalf("It's no use!")
		}
	})

	t.Run("testdata/vorbiscomment.ogg", func(t *testing.T) {
		ti, err := OpenAndReadFile("../../testdata/vorbiscomment.ogg")
		if err != nil {
			t.Fatalf("Error opening file")
		}
		if ti["ALBUM"] != "Test Files" {
			t.Fatalf("It's no use!")
		}
	})

}

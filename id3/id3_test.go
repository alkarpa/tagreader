package id3

import (
	"os"
	"testing"
)

func tagTest(t *testing.T, tags map[string]string, key, expected string) {
	received := tags[key]
	if received != expected {
		t.Fatalf("%s, expected '%s', received '%s'", key, expected, received)
	}
}

func TestId3(t *testing.T) {

	t.Run("Empty file", func(t *testing.T) {
		id3v2_3, err := os.Open("../testdata/empty_file")
		if err != nil {
			t.Fatal(err)
		}
		_, err = ReadId3ToMap(id3v2_3)
		if err == nil {
			t.Fatal("Reading an empty file")
		}
	})
	t.Run("Garbage header", func(t *testing.T) {
		id3v2_3, err := os.Open("../testdata/unknown_identifier")
		if err != nil {
			t.Fatal(err)
		}
		_, err = ReadId3ToMap(id3v2_3)
		if err == nil {
			t.Fatal("The header should have given bad instructions")
		}
	})

	t.Run("testdata/id3v2_3.mp3", func(t *testing.T) {
		id3v2_3, err := os.Open("../testdata/id3v2_3.mp3")
		if err != nil {
			t.Fatal(err)
		}

		trackinfo, err := ReadId3ToMap(id3v2_3)
		if err != nil {
			t.Fatal(err)
		}

		tagTest(t, trackinfo, "Title", "ID3v2.3.0")
		tagTest(t, trackinfo, "Album", "Test Files")
		tagTest(t, trackinfo, "Year", "2021")
		tagTest(t, trackinfo, "Artist", "alkarpa")

	})

}

package tagreader

import (
	"fmt"
	"os"
	"testing"

	"github.com/alkarpa/tagreader/util"
)

func TestTagReader(t *testing.T) {

	t.Run("File does not exist", func(t *testing.T) {
		received := GetTrackInfo("testdata/no_such_file")
		expected := TrackInfo{}
		if !expected.equals(received) {
			t.Fatalf("Expected %s, received %s", expected, received)
		}
	})

	t.Run("Id3", func(t *testing.T) {
		received := GetTrackInfo("testdata/id3v2_3.mp3")
		expected := TrackInfo{
			Album:   "Test Files",
			Artist:  "alkarpa",
			Title:   "ID3v2.3.0",
			TrackNo: "1",
			Year:    "2021",
		}
		if !expected.equals(received) {
			t.Fatalf("Expected %s, received %s", expected, received)
		}
	})
	t.Run("OggS", func(t *testing.T) {
		received := GetTrackInfo("testdata/vorbiscomment.ogg")
		expected := TrackInfo{
			Album:   "Test Files",
			Artist:  "alkarpa",
			Title:   "Vorbis Comment",
			TrackNo: "4",
			Year:    "2021",
		}
		if !expected.equals(received) {
			t.Fatalf("Expected %s, received %s", expected, received)
		}
	})
	t.Run("fLaC", func(t *testing.T) {
		received := GetTrackInfo("testdata/flaco.flac")
		expected := TrackInfo{
			Album:   "Test Files",
			Artist:  "alkarpa",
			Title:   "FLAC file",
			TrackNo: "2",
			Year:    "2021",
		}
		if !expected.equals(received) {
			t.Fatalf("Expected %s, received %s", expected, received)
		}
	})

	t.Run("!TrackInfo.equals", func(t *testing.T) {
		ti, ti2 := &TrackInfo{Title: "Test"}, &TrackInfo{Album: "tesT"}
		if ti.equals(*ti2) {
			t.Fatalf("%s should not be equal to %s", ti, ti2)
		}

	})

	t.Run("ImageData", func(t *testing.T) {
		expected, err := os.ReadFile("testdata/cover.png")
		if err != nil {
			t.Fatal(err)
		}
		t.Run("id3", func(t *testing.T) {
			received := GetTrackImage("testdata/id3v2_3.mp3")
			if !util.CompareByteSlices(received, expected) {
				t.Fatalf("ImageData differs, expected len %d, received len %d", len(expected), len(received))
			}
		})
		t.Run("ogg", func(t *testing.T) {
			received := GetTrackImage("testdata/vorbiscomment.ogg")
			fmt.Println("expected", expected)
			fmt.Println("received", received)
			if !util.CompareByteSlices(received, expected) {
				t.Fatalf("ImageData differs, expected len %d, received len %d", len(expected), len(received))
			}
		})

	})
}

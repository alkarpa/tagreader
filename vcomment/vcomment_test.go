package vcomment

import (
	"os"
	"testing"
)

// Duplicate --- should be moved to a tool package
func tagTest(t *testing.T, tags map[string]string, key, expected string) {
	if tags[key] != expected {
		t.Fatalf("%s, expected '%s', received '%s'", key, expected, tags[key])
	}
}

func TestVorbisCommentOgg(t *testing.T) {

	t.Run("Open vorbiscomment.ogg", func(t *testing.T) {
		f, err := os.Open("../testdata/vorbiscomment.ogg")
		if err != nil {
			t.Fatal(err)
		}
		tags, err := ReadVorbisCommentToMap(f, "OggS")
		if err != nil {
			t.Fatal(err)
		}

		//tags := vc.user_comments.comments

		t.Run("TITLE=Vorbis Comment", func(t *testing.T) {
			tagTest(t, tags, "TITLE", "Vorbis Comment")
		})
		t.Run("DATE=2021", func(t *testing.T) {
			tagTest(t, tags, "DATE", "2021")
		})
		t.Run("ALBUM=Test Files", func(t *testing.T) {
			tagTest(t, tags, "ALBUM", "Test Files")
		})
		t.Run("ARTIST=alkarpa", func(t *testing.T) {
			tagTest(t, tags, "ARTIST", "alkarpa")
		})
		t.Run("Comment is longer that 256 characters (over 1 byte length value)", func(t *testing.T) {
			if len(tags["COMMENT"]) <= 256 {
				t.Fatalf("Comment length under 1 byte value, was %d", len(tags["COMMENT"]))
			}
		})
		t.Run("Comment ends with '123456'", func(t *testing.T) {
			if tags["COMMENT"][len(tags["COMMENT"])-6:] != "123456" {
				t.Fatalf("Comment expected to end '123456', instead was %s", tags["COMMENT"][len(tags["COMMENT"])-7:])
			}
		})
	})

	t.Run("Open flaco.flac", func(t *testing.T) {
		f, err := os.Open("../testdata/flaco.flac")
		if err != nil {
			t.Fatal(err)
		}
		tags, err := ReadVorbisCommentToMap(f, "fLaC")
		if err != nil {
			t.Fatal(err)
		}

		t.Run("TITLE=FLAC file", func(t *testing.T) {
			tagTest(t, tags, "TITLE", "FLAC file")
		})
		t.Run("DATE=2021", func(t *testing.T) {
			tagTest(t, tags, "DATE", "2021")
		})
		t.Run("ALBUM=Test Files", func(t *testing.T) {
			tagTest(t, tags, "ALBUM", "Test Files")
		})
		t.Run("ARTIST=alkarpa", func(t *testing.T) {
			tagTest(t, tags, "ARTIST", "alkarpa")
		})
	})

	t.Run("Empty file", func(t *testing.T) {
		f, err := os.Open("../testdata/empty_file")
		if err != nil {
			t.Fatal(err)
		}
		t.Run("flac", func(t *testing.T) {
			flacVorbisComment(f)
		})
		f.Seek(0, 0)
		t.Run("ogg", func(t *testing.T) {
			oggVorbisComment(f)
		})
	})
	t.Run("Garbage file", func(t *testing.T) {
		f, err := os.Open("../testdata/unknown_identifier")
		if err != nil {
			t.Fatal(err)
		}
		t.Run("flac", func(t *testing.T) {
			flacVorbisComment(f)
		})
		f.Seek(0, 0)
		t.Run("ogg", func(t *testing.T) {
			oggVorbisComment(f)
		})
	})

}

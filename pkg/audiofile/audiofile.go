// Package audiofile provides a function to read tags from an audio file
package audiofile

import (
	"os"

	"github.com/alkarpa/tagreader/id3"
	"github.com/alkarpa/tagreader/vcomment"
)

// Opens the file at the specified path and attempts to read
// its metadata tags
func OpenAndReadFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	trackinfo, err := readTagsFromFile(f)
	if err != nil {
		return nil, err
	}

	f.Close()
	return trackinfo, nil
}

func readTagsFromFile(f *os.File) (map[string]string, error) {
	identifier := make([]byte, 4)
	if _, err := f.Read(identifier); err != nil {
		return nil, err
	}
	f.Seek(0, 0)

	switch string(identifier) {
	case "ID3\x03":
		return id3.ReadId3ToMap(f)
	case "OggS", "fLaC":
		return vcomment.ReadVorbisCommentToMap(f, string(identifier))
	}
	return make(map[string]string), nil
}

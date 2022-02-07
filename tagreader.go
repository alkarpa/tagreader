// Package tagreader provides a struct for storing audio tags
// that are interesting to a music player interface and a method
// to read the tags from a file.
package tagreader

import (
	"fmt"

	"github.com/alkarpa/tagreader/pkg/audiofile"
)

// Tries to read the audio tags from the file at the given path
func GetTrackInfo(path string) TrackInfo {
	tags, err := audiofile.OpenAndReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return newTrackInfo(tags)
}

func GetTrackImage(path string) []byte {
	tags, err := audiofile.OpenAndReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return newImageData(tags)
}

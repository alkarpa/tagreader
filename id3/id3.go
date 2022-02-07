// Package id3 contains code to read ID3v3.2.0 tags from a file
package id3

import (
	"os"
)

func ReadId3ToMap(f *os.File) (map[string]string, error) {
	id3, err := readId3v23(f)
	if err != nil {
		return nil, err
	}
	return id3.tags, nil
}

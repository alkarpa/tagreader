// Command-line program to read metadata from audio files.
// Expects a path to an audio file as an argument.
package main

import (
	"fmt"
	"os"

	"github.com/alkarpa/tagreader/pkg/audiofile"
)

func main() {
	args := os.Args[1:]
	for _, path := range args {
		ti, err := audiofile.OpenAndReadFile(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(ti)
	}
}

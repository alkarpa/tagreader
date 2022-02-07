// Package vcomment contains code to read Vorbis Comment tags from a file
package vcomment

import (
	"os"
)

func ReadVorbisCommentToMap(f *os.File, identifier string) (map[string]string, error) {
	var vc *vorbis_comment
	var err error
	switch identifier {
	case "OggS":
		vc, err = oggVorbisComment(f)
	case "fLaC":
		vc, err = flacVorbisComment(f)
	}
	if err != nil {
		return nil, err
	}
	return getCommentsMap(vc), nil // newTrackInfoFromVComment(vc)
}

func oggVorbisComment(f *os.File) (*vorbis_comment, error) {

	skipPage(f) // Vorbis comments are on the second page

	page, err := newPage(f)
	if err != nil {
		return nil, err
	}
	vc, err := findVorbisComment(page)
	if err != nil {
		return nil, err
	}

	return vc, nil
}

func findVorbisComment(p *page) (*vorbis_comment, error) {
	packt, err := findCommentPacket(p)
	if err != nil {
		return nil, err
	}
	vc := newVorbisComment(packt.body)
	return vc, nil
}

func flacVorbisComment(f *os.File) (*vorbis_comment, error) {

	_, err := f.Seek(4, 0)
	if err != nil {
		return nil, err
	}

	var vc *vorbis_comment

	var dat []byte
	keep_looping := true
	for keep_looping {
		dat, keep_looping = readMetadatablock(f)
		if len(dat) > 0 {
			vc = newVorbisComment(dat)
			break
		}
	}

	return vc, nil
}

package id3

import (
	"errors"
	"os"
)

var frameNames map[string]string = map[string]string{
	"TALB": "Album",
	"TIT2": "Title",
	"TPE1": "Artist",
	"TYER": "Year",
	"TRCK": "Track",
}

var header_bytes_length int = 10

type header struct {
	Id    string
	Size  int
	Flags [2]byte
}

func (h header) frameName() string {
	mapped := frameNames[h.Id]
	if mapped != "" {
		return mapped
	}
	return h.Id
}

type id3v23 struct {
	size int
	tags map[string]string
}

func (id3 *id3v23) readId3v2Size(f *os.File) error {
	identifier := make([]byte, 10)
	if _, err := f.Read(identifier); err != nil {
		return err
	}
	sizebytes := identifier[6:]
	id3.size = id3HeaderByteSize(sizebytes)
	return nil
}

func readId3v23(f *os.File) (id3v23, error) {
	id3 := &id3v23{
		tags: make(map[string]string),
	}

	err := id3.readId3v2Size(f)
	if err != nil {
		return *id3, err
	}

	tagbytes := make([]byte, id3.size)
	f.Read(tagbytes)

	var remaining_bytes []byte = tagbytes
	for len(remaining_bytes) > header_bytes_length {
		remaining_bytes, err = id3.readFrame(remaining_bytes)
		if err != nil {
			return *id3, err
		}
	}

	return *id3, nil
}

/*
func newTrackInfoFromId3v23(meta id3v23) *tagreader.TrackInfo {
	tags := meta.tags
	ti := &tagreader.TrackInfo{
		Album:   tags["Album"],
		Artist:  tags["Artist"],
		Title:   tags["Title"],
		TrackNo: tags["Track"],
		Year:    tags["Year"],
	}
	return ti
}
*/

// id3v2.3.0 informal standard 3.3
func (t id3v23) readTagHeader(bytes []byte) *header {
	frame := bytes[:4]
	size := id3FrameByteSize(bytes[4:8]) // id3ByteSizer(bytes[4:8], 256)
	flags := [2]byte{bytes[8], bytes[9]}
	return &header{Id: string(frame), Size: size, Flags: flags}
}

// Returns the remaining unread bytes
func (t id3v23) readFrame(bytes []byte) ([]byte, error) {
	header := t.readTagHeader(bytes)

	if header.Id == "\x00\x00\x00\x00" { // reached padding
		return make([]byte, 0), nil
	}

	frameSize := header_bytes_length + header.Size
	if frameSize > len(bytes) {
		return nil, errors.New("trying to read a file beyond its bounds")
	}

	frameSlice := bytes[header_bytes_length:frameSize]

	value := getFrameValue(header.Id, frameSlice)

	if len(value) > 0 {
		t.tags[header.frameName()] = value
	}

	return bytes[frameSize:], nil
}

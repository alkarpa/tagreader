package vcomment

import (
	"errors"
	"os"
)

type page struct {
	//bytes                   []byte
	capture                 []byte //4
	version                 byte
	header_type             byte
	granule_position        []byte // 8
	bitstream_serial_number []byte // 4
	page_sequence_number    []byte // 4
	checksum                []byte // 4
	page_segments           byte
	segment_table           []byte // page_segments
	body                    []byte
}

func readPageHeaderBytes(f *os.File) ([]byte, error) {
	b := make([]byte, 27)
	_, err := f.Read(b)
	if err != nil {
		return nil, err // errors.New("could not read 27 header bytes from file")
	}

	page_segments := b[26]

	segments_table := make([]byte, int(page_segments))
	_, err = f.Read(segments_table)
	if err != nil {
		return nil, errors.New("could not read segments table from file")
	}

	b = append(b, segments_table...)

	return b, nil
}
func readPageBodyBytes(f *os.File, header []byte) ([]byte, error) {
	segments_table := header[27:]
	segments_length_sum := 0
	for _, le := range segments_table {
		segments_length_sum += int(le)
	}

	b := make([]byte, segments_length_sum)
	_, err := f.Read(b)
	if err != nil {
		return nil, errors.New("could not read page body from file")
	}
	return b, nil
}

func newPage(f *os.File) (*page, error) {

	header, err := readPageHeaderBytes(f)
	if err != nil {
		return nil, err
	}
	body, err := readPageBodyBytes(f, header)
	if err != nil {
		return nil, err
	}

	page := &page{
		capture:                 header[:4],
		version:                 header[4],
		header_type:             header[5],
		granule_position:        header[6:14],
		bitstream_serial_number: header[14:18],
		page_sequence_number:    header[18:22],
		checksum:                header[22:26],
		page_segments:           header[26],
		segment_table:           header[27:],
		body:                    body,
	}

	return page, nil
}

func skipPage(f *os.File) {
	b, err := readPageHeaderBytes(f)
	if err != nil {
		return
	}
	readPageBodyBytes(f, b)
}

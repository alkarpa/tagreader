package vcomment

import (
	"errors"
	"fmt"
)

func findCommentPacket(page *page) (*packet, error) {

	segment_size := 0
	last_index := 0
	for i, s := range page.segment_table {
		segment_size += int(s)
		if s < 255 || i == len(page.segment_table)-1 {
			packet_bytes := page.body[last_index:segment_size]
			packet, err := newPacket(packet_bytes)
			if err != nil {
				return nil, err
			}

			if packet.header.packet_type == packet_type_comment {
				return packet, nil
			}

			last_index += segment_size
			segment_size = 0
		}
	}

	return nil, errors.New("could not find comment package")

}

type packet struct {
	header vorbis_common_packet_header
	body   []byte
}

func newPacket(b []byte) (*packet, error) {
	header, err := newVorbisCommonPacketHeader(b)
	if err != nil {
		return nil, err
	}

	return &packet{
		header: *header,
		body:   b[7:],
	}, nil
}

const packet_type_identification byte = 0x01
const packet_type_comment byte = 0x03
const packet_type_setup byte = 0x05

type vorbis_common_packet_header struct {
	packet_type byte    // 1: identification; 3: comment; 5: setup
	vorbis      [6]byte // expected [0x76 0x6f 0x72 0x62 0x69 0x73] (vorbis)
}

func (vcph vorbis_common_packet_header) validPacketType() bool {
	switch vcph.packet_type {
	case packet_type_identification, packet_type_comment, packet_type_setup:
		return true
	default:
		return false
	}
}

func (vcph vorbis_common_packet_header) validVorbisBytes() bool {
	expected_vorbis_bytes := [6]byte{0x76, 0x6f, 0x72, 0x62, 0x69, 0x73} // Vorbis I specification
	return vcph.vorbis == expected_vorbis_bytes
}

func newVorbisCommonPacketHeader(packet []byte) (*vorbis_common_packet_header, error) {
	header := &vorbis_common_packet_header{
		packet_type: packet[0],
	}
	copy(header.vorbis[:], packet[1:7])
	if !header.validPacketType() {
		return nil, fmt.Errorf("unexpected packet type %d in header", header.packet_type)
	}
	if !header.validVorbisBytes() {
		return nil, errors.New("packet header 'vorbis' not found")
	}

	return header, nil
}

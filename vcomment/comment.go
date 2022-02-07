package vcomment

import (
	"encoding/base64"
	"fmt"

	"github.com/alkarpa/tagreader/util"
)

type vorbis_comment struct {
	vendor        vorbis_comment_vendor
	user_comments vorbis_comment_user_comments
}

func newVorbisComment(packet []byte) *vorbis_comment {
	vcomment := &vorbis_comment{}

	pos := 0

	vcomment.vendor = *newVorbisCommentVendor(packet)
	pos += vcomment.vendor.fullLength()

	vcomment.user_comments = *newVorbisCommentUserComments(packet[pos:])

	return vcomment
}

type vorbis_comment_vendor struct {
	length [4]byte
	vendor []byte
}

func newVorbisCommentVendor(bytes []byte) *vorbis_comment_vendor {
	vcv := &vorbis_comment_vendor{}
	copy(vcv.length[:], bytes[:4])
	vendor_length := unsignedIntOf32bits(vcv.length[:])
	vcv.vendor = bytes[len(vcv.length) : len(vcv.length)+vendor_length]
	return vcv
}

func unsignedIntOf32bits(b []byte) int {
	return util.ByteSliceMultiplication(b, 256)
}

func (vcv vorbis_comment_vendor) fullLength() int {
	return len(vcv.length) + len(vcv.vendor)
}

type vorbis_comment_user_comments struct {
	comments map[string]string
}

func getCommentsMap(vc *vorbis_comment) map[string]string {
	return vc.user_comments.comments
}

func (vcuc vorbis_comment_user_comments) addComment(bytes []byte) int {
	length := bytes[:4]
	comment := bytes[4 : 4+unsignedIntOf32bits(length)]

	for i, b := range comment {
		if b == byte('=') {
			key := string(comment[:i])
			value := string(comment[i+1:])
			//if tagIsText(key) {
			vcuc.comments[key] = commentValueProcessing(key, value)
			//}
			break
		}
	}
	return len(length) + len(comment)
}

func commentValueProcessing(key, value string) string {
	if key == "METADATA_BLOCK_PICTURE" {
		return getMetadataBlockPictureBinaryDataString(value)
	}
	return value
}

func getMetadataBlockPictureBinaryDataString(value string) string {
	decoded, err := base64.StdEncoding.DecodeString(string(value))
	if err != nil {
		return ""
	}
	fmt.Println("decoded len", len(decoded))
	mime_len_bytes := util.ReverseByteSlice(decoded[4:8])
	mime_len := util.ByteSliceMultiplication(mime_len_bytes, 256)
	fmt.Println("mime len", mime_len)
	pos := 8 + mime_len
	desc_len_bytes := util.ReverseByteSlice(decoded[pos : pos+4])
	desc_len := util.ByteSliceMultiplication(desc_len_bytes, 256)
	pos += 4 + desc_len
	fmt.Println("desc len", desc_len)
	pos += 5 * 4
	fmt.Println("return len", len(decoded[pos:]))
	return string(decoded[pos:])
}

func newVorbisCommentUserComments(packet []byte) *vorbis_comment_user_comments {
	vcuc := &vorbis_comment_user_comments{
		comments: make(map[string]string),
	}
	pos := 0
	user_comment_list_length := packet[pos : pos+4]
	pos += len(user_comment_list_length)

	for i := 0; i < unsignedIntOf32bits(user_comment_list_length); i++ {
		pos += vcuc.addComment(packet[pos:])
	}

	return vcuc
}

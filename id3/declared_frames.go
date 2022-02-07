package id3

type declared_frame interface {
	list() []string
	read(bytes []byte) string
}

// id3v2.3.0 informal standard 4.2.1
type frame_text_information struct{}

func (fr frame_text_information) list() []string {
	return []string{
		"TALB", "TBPM", "TCOM", "TCON", "TCOP", "TDAT", "TDLY", "TENC",
		"TEXT", "TFLT", "TIME", "TIT1", "TIT2", "TIT3", "TKEY", "TLAN",
		"TLEN", "TMED", "TOAL", "TOFN", "TOLY", "TOPE", "TORY", "TOWN",
		"TPE1", "TPE2", "TPE3", "TPE4", "TPOS", "TPUB", "TRCK", "TRDA",
		"TRSN", "TRSO", "TSIZ", "TSRC", "TSSE", "TYER",
	}
}

func (fr frame_text_information) read(bytes []byte) string {
	encoding := bytes[0]
	valueBytes := bytesUntilEncodingEnd(encoding, bytes[1:])
	return decodeStringBytes(encoding, valueBytes)
}

// id3v2.3.0 informal standard 4.2.2
type frame_user_defined_text_information struct{}

func (fr frame_user_defined_text_information) list() []string {
	return []string{"TXXX"}
}

func (fr frame_user_defined_text_information) read(bytes []byte) string {
	encoding := bytes[0]
	descBytes := bytesUntilEncodingEnd(encoding, bytes[1:])
	valueBytes := bytesUntilEncodingEnd(encoding, bytes[1+len(descBytes):])
	return decodeStringBytes(encoding, valueBytes)
}

// id3v2.3.0 informal standard 4.3.1
type frame_url_link struct{}

func (fr frame_url_link) list() []string {
	return []string{
		"WCOM", "WCOP", "WOAF", "WOAR", "WOAS", "WORS", "WPAY", "WPUB",
	}
}

func (fr frame_url_link) read(bytes []byte) string {
	valueBytes := bytesUntilEncodingEnd(0, bytes)
	return decodeStringBytes(0x00, valueBytes)
}

// id3v2.3.0 informal standard 4.11
type frame_comments struct{}

func (fr frame_comments) list() []string {
	return []string{"COMM"}
}

func (fr frame_comments) read(bytes []byte) string {
	encoding := bytes[0]
	//language := string( bytes[1:4] )
	descBytes := bytesUntilEncodingEnd(encoding, bytes[4:])
	valueBytes := bytesUntilEncodingEnd(encoding, bytes[4+len(descBytes):])
	return decodeStringBytes(encoding, valueBytes)
}

// id3v2.3.0 informal standard 4.15
type frame_attached_picture struct{}

func (fr frame_attached_picture) list() []string {
	return []string{"APIC"}
}

func (fr frame_attached_picture) read(bytes []byte) string {

	encoding := bytes[0]
	mimeBytes := bytesUntilEncodingEnd(encoding, bytes[1:])
	pos := 1 + len(mimeBytes)
	//pictureByte := bytes[pos]
	pos += 1
	descriptionBytes := bytesUntilEncodingEnd(encoding, bytes[pos:])
	pos += len(descriptionBytes)
	imageData := bytes[pos:]

	return string(imageData)
	//fmt.Printf("MIME: %s; type: %d; description: '%s'\n", string(mimeBytes), pictureByte, string(descriptionBytes))

}

var declared_frames []declared_frame = []declared_frame{
	&frame_text_information{},
	&frame_user_defined_text_information{},
	&frame_url_link{},
	&frame_comments{},
	&frame_attached_picture{},
}

func getFrameValue(framename string, bytes []byte) string {
	//fmt.Println("Finding frame", framename)
	for _, t := range declared_frames {
		if arrayContains(t.list(), framename) {
			return t.read(bytes)
		}
	}
	return ""
}

func arrayContains(arr []string, match string) bool {
	for _, s := range arr {
		if s == match {
			return true
		}
	}
	return false
}

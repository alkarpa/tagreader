package tagreader

import (
	"strings"
)

type TrackInfo struct {
	Album   string
	Artist  string
	Title   string
	TrackNo string
	Year    string
}

func newTrackInfo(tags map[string]string) TrackInfo {
	for key, value := range tags {
		upperKey := strings.ToUpper(key)
		tags[upperKey] = removeNulls(value)
	}
	return TrackInfo{
		Album:   tags["ALBUM"],
		Artist:  tags["ARTIST"],
		Title:   tags["TITLE"],
		TrackNo: tagSynonyms(tags, "TRACK", "TRACKNUMBER"), // tags["TRACK"],
		Year:    tagSynonyms(tags, "YEAR", "DATE"),         //tags["YEAR"],
	}
}

func (ti TrackInfo) equals(ti2 TrackInfo) bool {
	if ti.Album != ti2.Album ||
		ti.Artist != ti2.Artist ||
		ti.Title != ti2.Title ||
		ti.TrackNo != ti2.TrackNo ||
		ti.Year != ti2.Year {
		return false
	}
	return true
}

func tagSynonyms(tags map[string]string, synonyms ...string) string {
	for _, synonym := range synonyms {
		if tags[synonym] != "" {
			return tags[synonym]
		}
	}
	return ""
}

func removeNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}

func newImageData(tags map[string]string) []byte {
	for key, value := range tags {
		upperKey := strings.ToUpper(key)
		tags[upperKey] = value
	}
	var imageData []byte
	imagestring := tagSynonyms(tags, "APIC", "METADATA_BLOCK_PICTURE")
	imageData = []byte(imagestring)
	return imageData
}

package main

import (
	"testing"
)
var torrent1 = "[Taka]_Naruto_Shippuuden_177_[720p][6EC1F800].mp4"
var torrent2 = "[BakedFish] Ansatsu Kyoushitsu (2015) - 06 [720p][AAC].mp4"
var torrent3 = "[Watashi]_Parasyte_-_the_maxim_-_04_[720p][A7E97910].mkv"

func TestStripEpisodeNum(t *testing.T){
	want1 := 177
	value := StripEpisodeNum(torrent1)
	want2 := 6
	value2 := StripEpisodeNum(torrent2)
	want3 := 4
	value3 := StripEpisodeNum(torrent3)
	if value != want1 || value2 != want2 || value3 != want3{
		t.Fail()
	}
}

func TestStripSeriesTitle(t *testing.T){
	want1 := "naruto shippuuden"
	value := StripSeriesTitle(torrent1)
	want2 := "ansatsu kyoushitsu"
	value2 := StripSeriesTitle(torrent2)
	want3 := "parasyte"
	value3 := StripSeriesTitle(torrent3)
	if value != want1 || value2 != want2 || value3 != want3{
		t.Fail()
	}
}

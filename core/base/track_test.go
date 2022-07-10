package base

import "testing"

func TestTrack(t *testing.T) {
	var artist = Artist{}
	var names = "FFFF, TTTT"
	var namesSlice = artist.namesToSlice(names)
	//
	var track = Track{}
	track.Artist = namesSlice
	track.Title = "Title"
	var searchable = track.ToSearchable()
	t.Log(searchable)
}

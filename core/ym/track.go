package ym

import (
	"github.com/oklookat/toanother/core/base"
)

// track.
type Track struct {
	Albums     []*Album  `json:"albums"`
	Artists    []*Artist `json:"artists"`
	DurationMs int       `json:"durationMs"`
	ID         string    `json:"id"`
	Title      string    `json:"title"`
}

func (p *Track) ToBase(playlistID int64) (baseTrack *base.Track, err error) {
	baseTrack = &base.Track{}
	baseTrack.Title = p.Title
	baseTrack.DurationMs = p.DurationMs
	baseTrack.PlaylistID = playlistID

	if p.Albums != nil && len(p.Albums) > 0 {
		baseTrack.AlbumTitle = &p.Albums[0].Title
	}

	var ar = Artist{}
	baseTrack.Artist, err = ar.collectNames(p.Artists)
	if err != nil {
		return
	}

	baseTrack.ID, err = baseTrack.AddToTable(dbConn)
	return
}

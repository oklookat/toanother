package spotify

import (
	"context"
	"errors"
	"math"

	"github.com/oklookat/toanother/core/base"
	"github.com/zmb3/spotify/v2"
)

type track struct {
	Searcher     searcher
	Adder        trackAdder
	OnProcessing func(current, total int)
	OnNotFound   func(item *base.Track)
}

// Find track id by base Track.
func (t *track) Find(track *base.Track) (found bool, id spotify.ID, err error) {
	if t.Searcher == nil {
		err = errors.New("nil searcher")
		return
	}
	if track == nil {
		err = errors.New("nil track")
		return
	}

	defer func() {
		if err != nil || found || t.OnNotFound == nil {
			return
		}
		t.OnNotFound(track)
	}()

	result, err := t.Searcher.Search(context.Background(),
		track.ToSearchable(), spotify.SearchTypeTrack)
	if err != nil {
		return
	}

	if result == nil || result.Tracks == nil || result.Tracks.Tracks == nil || result.Tracks.Total < 1 {
		return
	}

	for _, ft := range result.Tracks.Tracks {
		var foundDur = ft.Duration
		var trackDur = track.DurationMs
		var diff = math.Abs(float64(foundDur - trackDur))
		// if duration difference > 3 seconds - not our track
		if diff > 3000.0 {
			continue
		}
		found = true
		id = ft.ID
		break
	}
	return
}

// Add (like) tracks.
func (t *track) AddToLibrary(tracks []*base.Track) (err error) {
	if t.Adder == nil {
		err = errors.New("nil adder")
		return
	}
	if tracks == nil {
		err = errors.New("nil tracks")
		return
	}

	ids, err := findIds[*base.Track](tracks, t, t.OnProcessing)
	if err != nil {
		return
	}

	for counter := range ids {
		err = t.Adder.AddTracksToLibrary(context.Background(), ids[counter]...)
		if err != nil {
			return
		}
	}
	return
}

package spotify

import (
	"context"
	"errors"
	"math"

	"github.com/oklookat/toanother/core/base"
	"github.com/zmb3/spotify/v2"
)

type track struct {
	instance *Instance
}

func (t *track) Find(track *base.Track) (found bool, id spotify.ID, err error) {
	if t.instance == nil {
		err = errors.New("nil instance")
		return
	}
	if track == nil {
		err = errors.New("nil track")
		return
	}
	result, err := t.instance.client.Search(context.Background(),
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

func (t *track) OnFinish(ids [][]spotify.ID) (err error) {
	if t.instance == nil {
		err = errors.New("nil instance")
		return
	}
	// TODO: uncomment
	// for counter := range ids {
	// 	err = t.instance.client.AddTracksToLibrary(context.Background(), ids[counter]...)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	return
}

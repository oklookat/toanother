package spotify

import (
	"context"
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/zmb3/spotify/v2"
)

type artist struct {
	Searcher     searcher
	Follower     artistFollower
	OnProcessing func(current, total int)
	OnNotFound   func(item *base.Artist)
}

// Follow.
func (a *artist) Follow(artists []*base.Artist) (err error) {
	if a.Follower == nil {
		err = errors.New("nil follower")
		return
	}
	if artists == nil {
		err = errors.New("nil artists")
		return
	}

	ids, err := findIds[*base.Artist](artists, a, a.OnProcessing)
	if err != nil {
		return
	}

	for counter := range ids {
		err = a.Follower.FollowArtist(context.Background(), ids[counter]...)
		if err != nil {
			return
		}
	}
	return
}

// Find artist by base Artist.
func (a *artist) Find(artist *base.Artist) (found bool, id spotify.ID, err error) {
	if a.Searcher == nil {
		err = errors.New("nil searcher")
		return
	}
	if artist == nil {
		err = errors.New("nil artist")
		return
	}

	defer func() {
		if err != nil || found || a.OnNotFound == nil {
			return
		}
		a.OnNotFound(artist)
	}()

	result, errd := a.Searcher.Search(context.Background(),
		artist.Name, spotify.SearchTypeArtist)
	if errd != nil {
		return
	}
	if result == nil || result.Artists == nil || result.Artists.Artists == nil || result.Artists.Total < 1 {
		return
	}
	for i := range result.Artists.Artists {
		found = true
		id = result.Artists.Artists[i].ID
		break
	}
	return
}

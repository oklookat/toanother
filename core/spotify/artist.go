package spotify

import (
	"context"
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/zmb3/spotify/v2"
)

type artist struct {
	instance *Instance
}

func (a *artist) Find(artist *base.Artist) (found bool, id spotify.ID, err error) {
	if a.instance == nil {
		err = errors.New("nil instance")
		return
	}
	if artist == nil {
		err = errors.New("nil artist")
		return
	}
	result, errd := a.instance.client.Search(context.Background(),
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
	}
	return
}

func (a *artist) OnFound(ids [][]spotify.ID) (err error) {
	if a.instance == nil {
		err = errors.New("nil instance")
		return
	}
	for counter := range ids {
		err = a.instance.client.FollowArtist(context.Background(), ids[counter]...)
		if err != nil {
			return
		}
	}
	return
}

func (a *artist) OnImport(current int, total int, notFound []any) {
	if a.instance.baseHooks == nil || a.instance.baseHooks.OnImport == nil {
		return
	}
	a.instance.baseHooks.OnImport(current, total, notFound)
}

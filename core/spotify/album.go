package spotify

import (
	"context"
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/zmb3/spotify/v2"
)

type album struct {
	instance *Instance
}

func (a *album) Find(album *base.Album) (found bool, id spotify.ID, err error) {
	if a.instance == nil {
		err = errors.New("nil instance")
		return
	}
	if album == nil {
		err = errors.New("nil album")
		return
	}
	result, err := a.instance.client.Search(context.Background(),
		album.ToSearchable(), spotify.SearchTypeAlbum)
	if err != nil {
		return
	}
	if result == nil || result.Albums == nil || result.Albums.Albums == nil || result.Artists.Total < 1 {
		return
	}
	for _, ft := range result.Albums.Albums {
		found = true
		id = ft.ID
		break
	}
	return
}

func (a *album) OnFinish(ids [][]spotify.ID) (err error) {
	if a.instance == nil {
		err = errors.New("nil instance")
		return
	}
	// TODO: uncomment
	// for counter := range ids {
	// 	err = a.instance.client.AddAlbumsToLibrary(context.Background(), ids[counter]...)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	return
}

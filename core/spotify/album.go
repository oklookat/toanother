package spotify

import (
	"context"
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/utils"
	"github.com/zmb3/spotify/v2"
)

type album struct {
	Searcher     searcher
	Adder        albumAdder
	Currenter    currentUsersAlbums
	Remover      albumsRemover
	OnProcessing func(current, total int)
	OnNotFound   func(item *base.Album)
}

// Follow.
func (a *album) Follow(albums []*base.Album) (err error) {
	if a.Adder == nil {
		err = errors.New("nil adder")
		return
	}
	if albums == nil {
		err = errors.New("nil albums")
		return
	}

	ids, err := findIds[*base.Album](albums, a, a.OnProcessing)
	if err != nil {
		return
	}

	for counter := range ids {
		err = a.Adder.AddAlbumsToLibrary(context.Background(), ids[counter]...)
		if err != nil {
			return
		}
	}
	return
}

// Unfollow all albums.
func (a *album) UnfollowAll() (err error) {
	if a.Currenter == nil {
		err = errors.New("nil currenter")
		return
	}
	if a.Remover == nil {
		err = errors.New("nil remover")
		return
	}

	// get.
	albums, err := a.Currenter.CurrentUsersAlbums(context.Background(), nil)
	if err != nil {
		return
	}
	if albums.Albums == nil {
		return
	}

	// split.
	var ids = make([]spotify.ID, 0)
	for _, sa := range albums.Albums {
		ids = append(ids, sa.ID)
	}
	var idsSplit = utils.SplitSlice(ids, 20)
	if idsSplit == nil {
		return
	}

	// remove.
	for _, slice := range idsSplit {
		err = a.Remover.RemoveAlbumsFromLibrary(context.Background(), slice...)
		if err != nil {
			return
		}
	}

	return
}

// Find album by base.
func (a *album) Find(album *base.Album) (found bool, id spotify.ID, err error) {
	if a.Searcher == nil {
		err = errors.New("nil searcher")
		return
	}
	if album == nil {
		err = errors.New("nil album")
		return
	}

	defer func() {
		if err != nil || found || a.OnNotFound == nil {
			return
		}
		a.OnNotFound(album)
	}()

	result, err := a.Searcher.Search(context.Background(),
		album.ToSearchable(), spotify.SearchTypeAlbum)
	if err != nil {
		return
	}

	if result == nil || result.Albums == nil || result.Albums.Albums == nil || result.Albums.Total < 1 {
		return
	}

	for _, alb := range result.Albums.Albums {
		found = true
		id = alb.ID
		break
	}
	return
}

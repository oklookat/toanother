package spotify

import (
	"context"
	"errors"

	"github.com/oklookat/toanother/core/base"
)

const (
	WORKDIR   = "."
	TOKEN_DIR = WORKDIR + "/spotify_token.json"
)

type Hooks struct {
	OnURL  func(url string)
	Album  *base.Hooks[*base.Album]
	Artist *base.Hooks[*base.Artist]
	Track  *base.Hooks[*base.Track]
}

type Instance struct {
	hooks  *Hooks
	auth   *auth
	album  *album
	artist *artist
	track  *track
}

func New(h *Hooks) (i *Instance, err error) {
	i = &Instance{}
	i.hooks = h
	i.auth = &auth{}

	if err = i.auth.ByToken(context.Background()); err != nil {
		return
	}

	i.album = &album{
		Searcher:     i.auth.Client,
		Adder:        i.auth.Client,
		Currenter:    i.auth.Client,
		Remover:      i.auth.Client,
		OnProcessing: h.Album.OnProcessing,
		OnNotFound:   h.Album.OnNotFound,
	}
	i.artist = &artist{
		Searcher:     i.auth.Client,
		Follower:     i.auth.Client,
		OnProcessing: h.Artist.OnProcessing,
		OnNotFound:   h.Artist.OnNotFound,
	}
	i.track = &track{
		Searcher:     i.auth.Client,
		Adder:        i.auth.Client,
		OnProcessing: h.Track.OnProcessing,
		OnNotFound:   h.Track.OnNotFound,
	}

	return
}

func (i *Instance) ApplySettings(settings *base.SpotifySettings) (err error) {
	if settings == nil {
		err = errors.New("nil settings")
		return
	}
	if err = settings.Apply(); err != nil {
		return
	}
	if err = i.auth.ByToken(context.Background()); err != nil {
		return
	}
	return
}

func (i *Instance) WebAuth() (err error) {
	return i.auth.Web(i.hooks.OnURL)
}

func (i *Instance) Ping() (err error) {
	return i.auth.Ping()
}

func (i *Instance) AddTracks(tracks []*base.Track) (err error) {
	return i.track.AddToLibrary(tracks)
}

func (i *Instance) FollowAlbums(albums []*base.Album) (err error) {
	return i.album.Follow(albums)
}

func (i *Instance) FollowArtists(artists []*base.Artist) (err error) {
	return i.artist.Follow(artists)
}

func (i *Instance) UnfollowAllAlbums() (err error) {
	return i.album.UnfollowAll()
}

package spotify

import (
	"context"

	"github.com/oklookat/toanother/core/base"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

// TODO:
// сделать возможность отмены импорта треков(?).

const (
	// dirs.
	WORKDIR   = "."
	TOKEN_DIR = WORKDIR + "/spotify_token.json"

	// server.
	spotifyURI = "/spotify/callback"
	serverURI  = "http://localhost:8080" + spotifyURI
)

type Hooks struct {
	// open URL in args to auth.
	OnAuthURL func(url string)
}

type Instance struct {
	hooks     *Hooks
	baseHooks *base.Hooks
	// auth.
	isWebAuthCalledBefore bool
	state                 string
	token                 *oauth2.Token
	authenticator         *spotifyauth.Authenticator
	// main.
	client *spotify.Client
	user   *spotify.PrivateUser
}

func New(h *Hooks, bh *base.Hooks) (inst *Instance, err error) {
	inst = &Instance{}
	inst.hooks = h
	inst.baseHooks = bh
	inst.InitAuthenticator()
	inst.state = "abc123"
	if err = inst.readToken(); err != nil {
		return
	}
	if inst.token != nil {
		if err = inst.authByToken(context.Background()); err != nil {
			return
		}
	}
	return
}

func (i *Instance) InitAuthenticator() {
	i.authenticator = spotifyauth.New(
		spotifyauth.WithClientID(base.ConfigFile.Spotify.ID),
		spotifyauth.WithClientSecret(base.ConfigFile.Spotify.Secret),
		spotifyauth.WithRedirectURL(serverURI),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserLibraryRead,
			spotifyauth.ScopeUserLibraryModify,
			spotifyauth.ScopeUserFollowRead,
			spotifyauth.ScopeUserFollowModify,
			spotifyauth.ScopePlaylistReadPrivate,
			spotifyauth.ScopePlaylistModifyPrivate,
			spotifyauth.ScopeUserReadPrivate,
		),
	)
}

func (i *Instance) ImportLikedTracks(tracks []*base.Track) (err error) {
	var tr = &track{
		instance: i,
	}
	var args = &findIdsArgs[*base.Track]{
		instance: i,
		vals:     tracks,
		finder:   tr,
		hooks:    i.baseHooks,
	}
	return findIds(args)
}

func (i *Instance) ImportLikedArtists(artists []*base.Artist) (err error) {
	var ar = &artist{
		instance: i,
	}
	var args = &findIdsArgs[*base.Artist]{
		instance: i,
		vals:     artists,
		finder:   ar,
		hooks:    i.baseHooks,
	}
	return findIds(args)
}

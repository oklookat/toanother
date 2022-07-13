package main

import (
	"context"
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/logger"
	"github.com/oklookat/toanother/core/spotify"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type SpotifyApp struct {
	hooks      *spotify.Hooks
	trackHooks *base.Hooks[*base.Track]
	//
	ctx    context.Context
	Client *spotify.Instance
}

func (s *SpotifyApp) startup(ctx context.Context) (err error) {
	s.ctx = ctx
	s.hooks = &spotify.Hooks{
		OnAuthURL: func(url string) {
			runtime.EventsEmit(s.ctx, EVENT_SPOTIFY_AUTH_URL, url)
		},
	}
	s.trackHooks = &base.Hooks[*base.Track]{
		OnProcessing: func(current, total int) {
			runtime.EventsEmit(s.ctx, EVENT_PROCESSING, current, total)
		},
		OnNotFound: func(item *base.Track) {
			runtime.EventsEmit(s.ctx, EVENT_NOT_FOUND, item)
		},
	}
	s.Client, err = spotify.New(s.hooks)
	return
}

func (s *SpotifyApp) onFinish() {
	runtime.EventsEmit(s.ctx, EVENT_FINISH)
}

func (s *SpotifyApp) GetSettings() *base.SpotifySettings {
	return &base.ConfigFile.Spotify
}

func (s *SpotifyApp) ApplySettings(settings *base.SpotifySettings) (err error) {
	if settings == nil {
		err = errors.New("nil config pointer")
		return
	}
	if err = settings.Apply(); err != nil {
		logger.Log.Error(err.Error())
	} else {
		s.Client.InitAuthenticator()
	}
	return err
}

func (s *SpotifyApp) WebAuth() (err error) {
	return s.Client.WebAuth()
}

func (s *SpotifyApp) Ping() (err error) {
	return s.Client.Ping()
}

func (s *SpotifyApp) ImportLikedTracks(tracks []*base.Track) (err error) {
	if err = s.Client.ImportLikedTracks(tracks, s.trackHooks); err != nil {
		logger.Log.Error(err.Error())
	}
	s.onFinish()
	return
}

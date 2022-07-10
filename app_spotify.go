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
	hooks     *spotify.Hooks
	baseHooks *base.Hooks
	ctx       context.Context
	Client    *spotify.Instance
}

func (s *SpotifyApp) startup(ctx context.Context) (err error) {
	s.ctx = ctx
	s.hooks = &spotify.Hooks{
		OnAuthURL: func(url string) {
			runtime.EventsEmit(s.ctx, "SPOTIFY_AUTH_URL", url)
		},
	}
	s.baseHooks = &base.Hooks{
		OnImport: func(current, total int, notFound []interface{}) {
			runtime.EventsEmit(s.ctx, "OnImport", current, total, notFound)
		},
	}
	s.Client, err = spotify.New(s.hooks, s.baseHooks)
	return
}

func (s *SpotifyApp) onFinish() {
	runtime.EventsEmit(s.ctx, "OnFinish")
}

func (s *SpotifyApp) GetSettings() *base.SpotifySettings {
	return spotify.Settings
}

func (s *SpotifyApp) ApplySettings(settings *base.SpotifySettings) (err error) {
	if settings == nil {
		err = errors.New("nil config pointer")
		return
	}
	if err = settings.Apply(); err != nil {
		logger.Log.Error(err.Error())
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
	if err = s.Client.ImportLikedTracks(tracks); err != nil {
		logger.Log.Error(err.Error())
	}
	s.onFinish()
	return
}

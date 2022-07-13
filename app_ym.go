package main

import (
	"context"
	"errors"

	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/logger"
	"github.com/oklookat/toanother/core/ym"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type YandexMusicApp struct {
	ctx      context.Context
	playlist ym.Playlist
	artist   ym.Artist
	album    ym.Album
}

func (y *YandexMusicApp) startup(ctx context.Context) (err error) {
	if err = ym.Init(); err != nil {
		return
	}
	y.ctx = ctx
	y.playlist.Hooks = &base.Hooks[*base.Playlist]{
		OnProcessing: func(current, total int) {
			runtime.EventsEmit(y.ctx, EVENT_PROCESSING, current, total)
		},
		OnNotFound: func(item *base.Playlist) {
			runtime.EventsEmit(y.ctx, EVENT_NOT_FOUND, item)
		},
	}
	y.artist.Hooks = &base.Hooks[*base.Artist]{
		OnProcessing: func(current, total int) {
			runtime.EventsEmit(y.ctx, EVENT_PROCESSING, current, total)
		},
		OnNotFound: func(item *base.Artist) {
			runtime.EventsEmit(y.ctx, EVENT_NOT_FOUND, item)
		},
	}
	y.album.Hooks = &base.Hooks[*base.Album]{
		OnProcessing: func(current, total int) {
			runtime.EventsEmit(y.ctx, EVENT_PROCESSING, current, total)
		},
		OnNotFound: func(item *base.Album) {
			runtime.EventsEmit(y.ctx, EVENT_NOT_FOUND, item)
		},
	}
	return
}

func (y *YandexMusicApp) onFinish() {
	runtime.EventsEmit(y.ctx, EVENT_FINISH)
}

//from database.
func (y *YandexMusicApp) GetPlaylists() (playlists []*base.Playlist, err error) {
	playlists, err = y.playlist.GetPlaylists()
	if err != nil {
		logger.Log.Error(err.Error())
	}
	y.onFinish()
	return
}

// from database.
func (y *YandexMusicApp) GetTracks(playlistID int64) (tracks []*base.Track, err error) {
	tracks, err = y.playlist.GetTracks(playlistID)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	y.onFinish()
	return
}

// from YM.
func (y *YandexMusicApp) DownloadPlaylists() (playlists []*base.Playlist, err error) {
	playlists, err = y.playlist.DownloadAll()
	if err != nil {
		logger.Log.Error(err.Error())
	}
	y.onFinish()
	return
}

// from database.
func (y *YandexMusicApp) GetArtists() (artists []*base.Artist, err error) {
	artists, err = y.artist.GetAll()
	if err != nil {
		logger.Log.Error(err.Error())
	}
	y.onFinish()
	return
}

// from YM.
func (y *YandexMusicApp) DownloadArtists() (artists []*base.Artist, err error) {
	artists, err = y.artist.DownloadAll()
	if err != nil {
		logger.Log.Error(err.Error())
	}
	y.onFinish()
	return
}

// from database.
func (y *YandexMusicApp) GetAlbums() (albums []*base.Album, err error) {
	albums, err = y.album.GetAll()
	if err != nil {
		logger.Log.Error(err.Error())
	}
	y.onFinish()
	return
}

// from database.
func (y *YandexMusicApp) DownloadAlbums() (albums []*base.Album, err error) {
	albums, err = y.album.DownloadAll()
	if err != nil {
		logger.Log.Error(err.Error())
	}
	y.onFinish()
	return
}

func (y *YandexMusicApp) GetSettings() *base.YandexMusicSettings {
	return &base.ConfigFile.YandexMusic
}

func (y *YandexMusicApp) ApplySettings(settings *base.YandexMusicSettings) (err error) {
	if settings == nil {
		err = errors.New("nil settings pointer")
		return
	}
	if err = settings.Apply(); err != nil {
		logger.Log.Error(err.Error())
	}
	return err
}

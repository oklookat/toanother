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
	hooks    *base.Hooks
	playlist ym.Playlist
	artist   ym.Artist
	album    ym.Album
}

func (y *YandexMusicApp) startup(ctx context.Context) (err error) {
	if err = ym.Init(); err != nil {
		return
	}
	y.ctx = ctx
	y.hooks = &base.Hooks{
		OnFetchFromAPI: func(current, total int) {
			runtime.EventsEmit(y.ctx, "OnFetch", current, total)
		},
		OnFetchFromDatabase: func(current, total int) {
			runtime.EventsEmit(y.ctx, "OnFetchFromDatabase", current, total)
		},
		OnAddingToDatabase: func(current, total int) {
			runtime.EventsEmit(y.ctx, "OnAddingToDatabase", current, total)
		},
	}
	y.playlist.Hooks = y.hooks
	y.artist.Hooks = y.hooks
	y.album.Hooks = y.hooks
	return
}

func (y *YandexMusicApp) onFinish() {
	runtime.EventsEmit(y.ctx, "OnFinish")
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

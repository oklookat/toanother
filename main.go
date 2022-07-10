package main

import (
	"embed"

	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/datadir"
	"github.com/oklookat/toanother/core/logger"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	var err error

	// datadir.
	if err = datadir.Init(); err != nil {
		panic("[datadir] " + err.Error())
	}

	// config.
	var config = &base.Config{}
	if err = config.Init(); err != nil {
		panic("[config] " + err.Error())
	}

	// logger.
	if err = logger.Init(); err != nil {
		panic("[logger] " + err.Error())
	}

	// UI.
	var mainApp = &App{}
	var ymApp = &YandexMusicApp{}
	var spotifyApp = &SpotifyApp{}
	mainApp.strap(ymApp, spotifyApp)
	err = mainApp.run()

	if err != nil {
		logger.Log.Fatal("[UI] " + err.Error())
	}
}

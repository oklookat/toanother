package main

import (
	"context"

	"github.com/oklookat/toanother/core/logger"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	EVENT_SPOTIFY_AUTH_URL = "ON_SPOTIFY_AUTH_URL"
	EVENT_NOT_FOUND        = "ON_NOT_FOUND"
	EVENT_PROCESSING       = "ON_PROCESSING"
	EVENT_FINISH           = "ON_FINISH"
)

type otherApp interface {
	startup(ctx context.Context) (err error)
}

type App struct {
	ctx  context.Context
	apps []otherApp
}

func (a *App) run() (err error) {
	logger.Log.Info("[app] starting UI")

	var binds []interface{} = make([]interface{}, 0)
	binds = append(binds, a)
	for i := range a.apps {
		binds = append(binds, a.apps[i])
	}

	err = wails.Run(&options.App{
		Title:     "toanother",
		Width:     600,
		Height:    600,
		Assets:    assets,
		OnStartup: a.startup,
		Bind:      binds,
		Logger:    logger.Log,
	})

	return
}

// add app.
func (a *App) strap(b ...otherApp) {
	if b == nil {
		return
	}
	if a.apps == nil {
		a.apps = make([]otherApp, 0)
	}
	a.apps = append(a.apps, b...)
}

// execute app(s) startup().
func (a *App) startup(ctx context.Context) {
	logger.Log.Info("[app] starting modules")
	a.ctx = ctx
	if a.apps == nil {
		return
	}
	for i := range a.apps {
		if err := a.apps[i].startup(ctx); err != nil {
			logger.Log.Fatal(err.Error())
		}
	}
	a.apps = nil
}

func (a *App) MessageInfo(message string) error {
	var options = runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Info",
		Message: message,
	}
	_, err := runtime.MessageDialog(a.ctx, options)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	return err
}

func (a *App) MessageError(message string) error {
	var options = runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Error",
		Message: message,
	}
	_, err := runtime.MessageDialog(a.ctx, options)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	return err
}

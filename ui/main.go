package ui

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
)

var app = fyneApp.NewWithID("com.oklookat.toanother")
var mainWindow = app.NewWindow("toanother")

func Start() {
	mainWindow.SetMaster()
	mainWindow.Resize(fyne.NewSize(400, 400))
	//
	drawMainMenu()

	//
	mainWindow.ShowAndRun()
}

func drawMainMenu() {
	// routes.
	var ym = fyne.NewMenuItem("Yandex.Music", func() {})
	var spoty = fyne.NewMenuItem("Spotify", func() {})
	var routesMenu = fyne.NewMenu("Menu", ym, spoty)

	// help.
	var docs = fyne.NewMenuItem("Docs", func() {})
	var repo = fyne.NewMenuItem("Repo", func() {})
	var thParty = fyne.NewMenuItem("Licenses", func() {})
	var helpMenu = fyne.NewMenu("Help", docs, repo, thParty)

	// container.
	var mainMenu = fyne.NewMainMenu(routesMenu, helpMenu)
	mainWindow.SetMainMenu(mainMenu)
}

package ui

import (
	"net/url"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"github.com/oklookat/toanother/ui/internal/utils"
	ymapp "github.com/oklookat/toanother/ui/internal/ymApp"
)

type Instance struct {
	// main.
	app    fyne.App
	window fyne.Window

	//
	ym *ymapp.Instance
}

func (i *Instance) Start() {
	var err error
	defer func() {
		utils.SmoothError(err, i.window)
		i.window.ShowAndRun()
	}()

	//
	i.app = fyneApp.NewWithID("com.oklookat.toanother")
	i.window = i.app.NewWindow("toanother")
	i.window.SetMaster()
	i.window.Resize(fyne.NewSize(400, 400))

	// ym.
	if i.ym, err = ymapp.New(i.window); err != nil {
		return
	}

	i.drawMainMenu()
}

func (i *Instance) drawMainMenu() {
	// routes.
	var ymItem = fyne.NewMenuItem("Yandex.Music", func() {
		i.ym.DrawTabs()
	})
	var spotyItem = fyne.NewMenuItem("Spotify", func() {})
	var routesMenu = fyne.NewMenu("Menu", ymItem, spotyItem)

	// help.
	var docs = fyne.NewMenuItem("Docs", func() {})
	var repo = fyne.NewMenuItem("Repo", func() {
		var ur, _ = url.Parse("https://github.com/oklookat/toanother")
		i.app.OpenURL(ur)
	})
	var licenses = fyne.NewMenuItem("Licenses", func() {
		var ur, _ = url.Parse("https://github.com/oklookat/toanother/blob/main/LICENSE")
		i.app.OpenURL(ur)
	})
	var helpMenu = fyne.NewMenu("Help", docs, repo, licenses)

	// container.
	var mainMenu = fyne.NewMainMenu(routesMenu, helpMenu)
	i.window.SetMainMenu(mainMenu)
}

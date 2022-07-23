package ymapp

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Draw tabs (navigation).
func (i *Instance) DrawTabs() {
	// tabs.
	var playlists = i.playlist.GetTab()
	var artists = container.NewTabItem("Artists", widget.NewLabel("artist"))
	var albums = container.NewTabItem("Albums", widget.NewLabel("albums"))
	var settings = i.settings.GetTab()

	// tab container.
	var tabs = container.NewAppTabs(playlists, artists, albums, settings)
	i.window.Canvas().SetContent(tabs)
}

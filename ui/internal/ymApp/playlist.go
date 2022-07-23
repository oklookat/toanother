package ymapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/core/ym"
)

func newPlaylist(window fyne.Window) (p *playlist) {
	p = &playlist{
		window: window,
	}
	return
}

type playlist struct {
	window fyne.Window
	self   ym.Playlist
}

func (p *playlist) getBase() (playlists []*base.Playlist, err error) {
	playlists, err = p.self.GetPlaylists()
	return
}

// Draw UI (tab route).
func (p *playlist) GetTab() *container.TabItem {
	var rootContainer = container.NewVBox()

	var downloadButt = widget.NewButton("Download", func() {})
	rootContainer.Add(downloadButt)

	var playlists, _ = p.getBase()
	for _, p2 := range playlists {
		var item = widget.NewAccordionItem(p2.Title, widget.NewLabel("Title for title"))
		var acc = widget.NewAccordion(item)
		rootContainer.Add(acc)
	}
	return container.NewTabItem("Playlists", rootContainer)
}

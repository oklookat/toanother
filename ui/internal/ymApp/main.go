package ymapp

import (
	"errors"

	"fyne.io/fyne/v2"
	"github.com/oklookat/toanother/core/ym"
)

type Instance struct {
	window fyne.Window

	//
	settings   *settings
	playlist   *playlist
	coreArtist ym.Artist
	coreAlbum  ym.Album
}

func New(parent fyne.Window) (i *Instance, err error) {
	if parent == nil {
		err = errors.New("nil parent")
		return
	}
	if err = ym.Init(); err != nil {
		return
	}
	i = &Instance{
		window:   parent,
		settings: newSettings(parent),
		playlist: newPlaylist(parent),
	}
	return
}

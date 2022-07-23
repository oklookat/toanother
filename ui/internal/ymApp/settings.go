package ymapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/oklookat/toanother/core/base"
	"github.com/oklookat/toanother/ui/internal/utils"
)

func newSettings(window fyne.Window) (s *settings) {
	s = &settings{
		window: window,
	}
	if base.ConfigFile == nil {
		return
	}
	s.self = base.ConfigFile.YandexMusic
	return
}

type settings struct {
	window fyne.Window
	self   base.YandexMusicSettings
}

// Draw UI (tab route).
func (s *settings) GetTab() *container.TabItem {
	var applyButton *widget.Button
	applyButton = widget.NewButton("Apply", func() {
		s.apply()
		applyButton.Disable()
	})
	applyButton.Disable()
	var secondLine = container.NewHBox(applyButton)

	//
	var loginEntry = widget.NewEntry()
	loginEntry.SetPlaceHolder("Login")
	loginEntry.SetText(s.self.Login)
	loginEntry.OnChanged = func(val string) {
		s.self.Login = val
		applyButton.Enable()
	}
	var firstLine = container.NewMax(loginEntry)

	//
	var rootContainer = container.NewVBox(firstLine, secondLine)
	return container.NewTabItem("Settings", rootContainer)
}

// Write settings to config file.
func (s *settings) apply() {
	var err error
	defer func() {
		utils.SmoothError(err, s.window)
	}()
	err = s.self.Apply()
}

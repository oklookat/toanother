package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/oklookat/toanother/core/logger"
)

// Log error & show error dialog.
//
// If error or parent not exists does nothing.
func SmoothError(err error, parent fyne.Window) {
	if err == nil || parent == nil {
		return
	}
	logger.Log.Error(err.Error())
	dialog.ShowError(err, parent)
}

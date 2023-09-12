package ui

import (
	"PIXL/apptype"
	"PIXL/swatch"

	"fyne.io/fyne/v2"
	//"zerotomastery.io/PIXL/apptype"
	//"zerotomastery.io/PIXL/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}

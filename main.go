package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func loadUI() fyne.CanvasObject {
	list := widget.NewVBox(
		widget.NewLabel("Item 1"),
		widget.NewLabel("Item 2"))

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
		}))

	content := widget.NewMultiLineEntry()
	content.SetText("Note content")

	side := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, list)
	split := widget.NewHSplitContainer(side, content)
	split.Offset = 0.25
	return split
}

func main() {
	a := app.New()
	w := a.NewWindow("Notes")

	w.SetContent(loadUI())
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
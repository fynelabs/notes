package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var (
	content *widget.Entry
	current *note
)

func setNote(n *note) {
	current = n
	if n == nil {
		content.SetText("")
		return
	}

	content.SetText(n.content)
}

func loadUI(notes []*note) fyne.CanvasObject {
	list := widget.NewVBox()
	for _, n := range notes {
		theNote := n
		list.Append(widget.NewButton(n.title(), func() {
			setNote(theNote)
		}))
	}

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
		}))

	content = widget.NewMultiLineEntry()
	if len(notes) > 0 {
		setNote(notes[0])
	}

	side := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, list)
	split := widget.NewHSplitContainer(side, content)
	split.Offset = 0.25
	return split
}

func main() {
	a := app.New()
	w := a.NewWindow("Notes")

	notes := []*note{
		&note{"Note 1\nHas some content"},
		&note{"Note 2\nIs another note"},
	}

	w.SetContent(loadUI(notes))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
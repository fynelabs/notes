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
	list *widget.Box
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

func refreshList(n *notelist) {
	list.Children = nil
	for _, n := range n.notes {
		theNote := n
		list.Append(widget.NewButton(n.title(), func() {
			setNote(theNote)
		}))
	}
}

func loadUI(n *notelist) fyne.CanvasObject {
	list = widget.NewVBox()
	refreshList(n)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			setNote(n.add())
			refreshList(n)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			n.remove(current)
			refreshList(n)
			if len(n.notes) == 0 {
				setNote(nil)
			}
			setNote(n.notes[0])
		}))

	content = widget.NewMultiLineEntry()
	if len(n.notes) > 0 {
		setNote(n.notes[0])
	}

	side := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, list)
	split := widget.NewHSplitContainer(side, content)
	split.Offset = 0.25
	return split
}

func main() {
	a := app.New()
	w := a.NewWindow("Notes")

	notes := &notelist{[]*note{
		&note{"Note 1\nHas some content"},
		&note{"Note 2\nIs another note"},
	}}

	w.SetContent(loadUI(notes))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type ui struct {
	current *note
	notes   *notelist

	content *widget.Entry
	list    *widget.Box
}

func (u *ui) addNote() {
	newNote := u.notes.add()
	u.setNote(newNote)
}

func (u *ui) setNote(n *note) {
	if n == nil {
		u.content.SetText(u.placeholderContent())
		return
	}
	u.current = n
	u.content.SetText(n.content)
	u.refreshList()
}

func (u *ui) refreshList() {
	var list []fyne.CanvasObject
	for _, n := range u.notes.notes {
		thisNote := n
		button := widget.NewButton(n.title(), func() {
			u.setNote(thisNote)
		})
		if n == u.current {
			button.Style = widget.PrimaryButton
		}

		list = append(list, button)
	}

	u.list.Children = list
	u.list.Refresh()
}

func (u *ui) removeCurrentNote() {
	u.notes.remove(u.current)
	if len(u.notes.notes) > 0 {
		u.setNote(u.notes.notes[0])
	} else {
		u.setNote(nil)
	}
	u.refreshList()
}

func (u *ui) loadUI() fyne.CanvasObject {
	u.content = widget.NewMultiLineEntry()
	u.content.SetText(u.placeholderContent())

	u.list = widget.NewVBox()
	u.refreshList()

	if len(u.notes.notes) > 0 {
		u.setNote(u.notes.notes[0])
	}
	u.content.OnChanged = func(content string) {
		if u.current == nil {
			return
		}

		u.current.content = content
		u.notes.save()
		u.refreshList()
	}

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			u.addNote()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			u.removeCurrentNote()
		}),
	)

	side := fyne.NewContainerWithLayout(layout.NewBorderLayout(bar, nil, nil, nil),
		bar, u.list)

	split := widget.NewHSplitContainer(side, u.content)
	split.Offset = 0.25
	return split
}

func (u *ui) placeholderContent() string {
	text := "Welcome!\nTap '+' in the toolbar to add a note."
	return text
}

func main() {
	a := app.NewWithID("xyz.andy.notes")
	w := a.NewWindow("Notes")

	list := &notelist{pref: a.Preferences()}
	list.load()
	notesUI := &ui{notes: list}

	w.SetContent(notesUI.loadUI())
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

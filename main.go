package main

import (
	"fmt"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ui struct {
	current *note
	notes   *notelist

	content *widget.Entry
	list    *fyne.Container
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
			button.Importance = widget.HighImportance
		}

		list = append(list, button)
	}

	u.list.Objects = list
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

	u.list = container.NewVBox()
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
		bar, container.NewVScroll(u.list))

	return newAdaptiveSplit(side, u.content)
}

func (u *ui) registerKeys(w fyne.Window) {
	shortcut := &desktop.CustomShortcut{KeyName: fyne.KeyN, Modifier: desktop.ControlModifier}
	if runtime.GOOS == "darwin" {
		shortcut.Modifier = desktop.SuperModifier
	}

	w.Canvas().AddShortcut(shortcut, func(_ fyne.Shortcut) {
		u.addNote()
	})
}

func (u *ui) placeholderContent() string {
	text := "Welcome!\nTap '+' in the toolbar to add a note."
	if fyne.CurrentDevice().HasKeyboard() {
		modifier := "ctrl"
		if runtime.GOOS == "darwin" {
			modifier = "cmd"
		}
		text += fmt.Sprintf("\n\nOr use they keyboard shortcut %s+N.", modifier)
	}
	return text
}

func main() {
	a := app.NewWithID("xyz.andy.notes")
	w := a.NewWindow("Notes")

	list := &notelist{pref: a.Preferences()}
	list.load()
	notesUI := &ui{notes: list}

	w.SetContent(notesUI.loadUI())
	notesUI.registerKeys(w)

	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

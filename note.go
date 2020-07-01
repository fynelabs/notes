package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne"
)

const (
	countKey = "notecount"
	noteKey  = "note%d"
)

type note struct {
	content string
}

func (n *note) title() string {
	if n.content == "" {
		return "Untitled"
	}

	return strings.SplitN(n.content, "\n", 2)[0]
}

type notelist struct {
	notes []*note
	pref  fyne.Preferences
}

func (l *notelist) add() *note {
	n := &note{}
	l.notes = append([]*note{n}, l.notes...)
	l.save()
	return n
}

func (l *notelist) remove(n *note) {
	defer l.save()
	if len(l.notes) == 0 {
		return
	}

	for i, note := range l.notes {
		if note != n {
			continue
		}

		if i == len(l.notes)-1 {
			l.notes = l.notes[:i]
		} else {
			l.notes = append(l.notes[:i], l.notes[i+1:]...)
		}
		break
	}
}

func (l *notelist) load() {
	l.notes = nil
	count := l.pref.Int(countKey)
	if count == 0 {
		return
	}

	for i := 0; i < count; i++ {
		key := fmt.Sprintf(noteKey, i)
		content := l.pref.String(key)
		l.notes = append(l.notes, &note{content})
	}
}

func (l *notelist) save() {
	for i, note := range l.notes {
		key := fmt.Sprintf(noteKey, i)
		l.pref.SetString(key, note.content)
	}
	l.pref.SetInt(countKey, len(l.notes))
}

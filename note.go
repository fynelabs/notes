package main

import (
	"errors"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

const (
	countKey = "notecount"
	noteKey  = "note%d"
)

type note struct {
	content binding.String
}

func (n *note) title() binding.String {
	return newTitleString(n.content)
}

type notelist struct {
	notes []*note
	pref  fyne.Preferences
}

func (l *notelist) add() *note {
	key := fmt.Sprintf(noteKey, len(l.notes))
	n := &note{binding.BindPreferenceString(key, l.pref)}
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

	for i := count - 1; i >= 0; i-- {
		key := fmt.Sprintf(noteKey, i)
		content := binding.BindPreferenceString(key, l.pref)
		l.notes = append(l.notes, &note{content})
	}
}

func (l *notelist) save() {
	l.pref.SetInt(countKey, len(l.notes))
}

type titleString struct {
	binding.String
}

func (t *titleString) Get() (string, error) {
	content, err := t.String.Get()
	if err != nil {
		return "Error", err
	}

	if content == "" {
		return "Untitled", nil
	}

	return strings.SplitN(content, "\n", 2)[0], nil
}

func (t *titleString) Set(string) error {
	return errors.New("cannot set content from title")
}

func newTitleString(in binding.String) binding.String {
	return &titleString{in}
}

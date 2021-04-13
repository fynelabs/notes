package main

import (
	"errors"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

const (
	countKey       = "notecount"
	noteKey        = "note%d"
	noteDeletedKey = "note%ddeleted"
)

type note struct {
	content binding.String
	deleted binding.Bool
}

func (n *note) title() binding.String {
	return newTitleString(n.content)
}

type notelist struct {
	all  []*note
	pref fyne.Preferences
}

func (l *notelist) add() *note {
	key := fmt.Sprintf(noteKey, len(l.all))
	deleteKey := fmt.Sprintf(noteDeletedKey, len(l.all))
	n := &note{
		binding.BindPreferenceString(key, l.pref),
		binding.BindPreferenceBool(deleteKey, l.pref),
	}
	l.all = append([]*note{n}, l.all...)
	l.save()
	return n
}

func (l *notelist) delete(n *note) {
	n.deleted.Set(true)
}

func (l *notelist) load() {
	l.all = nil
	count := l.pref.Int(countKey)
	if count == 0 {
		return
	}

	for i := count - 1; i >= 0; i-- {
		key := fmt.Sprintf(noteKey, i)
		deleteKey := fmt.Sprintf(noteDeletedKey, i)
		content := binding.BindPreferenceString(key, l.pref)
		deleted := binding.BindPreferenceBool(deleteKey, l.pref)
		l.all = append(l.all, &note{content, deleted})
	}
}

func (l *notelist) notes() []*note {
	var visible []*note
	for _, n := range l.all {
		if del, _ := n.deleted.Get(); del {
			continue
		}
		visible = append(visible, n)
	}
	return visible
}

func (l *notelist) save() {
	l.pref.SetInt(countKey, len(l.all))
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

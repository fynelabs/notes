package main

import (
	"testing"

	"fyne.io/fyne/v2/test"

	"github.com/stretchr/testify/assert"
)

func testlist() *notelist {
	a := test.NewApp()
	n := &notelist{pref: a.Preferences()}

	return n
}

func TestNoteTitle(t *testing.T) {
	n := &note{"Some content"}
	assert.Equal(t, "Some content", n.title())

	n = &note{"line1\nline2"}
	assert.Equal(t, "line1", n.title())

	n = &note{content: ""}
	assert.Equal(t, "Untitled", n.title())
}

func TestNoteListAdd(t *testing.T) {
	notes := testlist()

	notes.add()
	assert.Equal(t, 1, len(notes.notes))
}

func TestNoteListRemove(t *testing.T) {
	first := &note{content: "remove me"}
	second := &note{content: "remove me2"}
	notes := testlist()
	notes.notes = []*note{first, second}

	assert.Equal(t, 2, len(notes.notes))
	notes.remove(first)
	assert.Equal(t, 1, len(notes.notes))
	notes.remove(second)
	assert.Equal(t, 0, len(notes.notes))
}

func TestNoteListLoad(t *testing.T) {
	l := testlist()
	n := l.add()
	defer l.remove(n)
	n.content = "Test"
	l.save()

	// get a new one
	l.load() // load fresh from preferences
	assert.Equal(t, 1, len(l.notes))
	assert.Equal(t, "Test", l.notes[0].content) // same content
}

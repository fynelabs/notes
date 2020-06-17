package main

import (
	"testing"

	"fyne.io/fyne/test"
	"github.com/stretchr/testify/assert"
)

func testlist() *notelist {
	a := test.NewApp()
	n := &notelist{pref: a.Preferences()}
	n.load()
	return n
}

func TestNoteTitle(t *testing.T) {
	n := &note{"Some content"}
	assert.Equal(t, "Some content", n.title())

	n = &note{"line1\nline2"}
	assert.Equal(t, "line1", n.title())
}

func TestNoteListAdd(t *testing.T) {
	l := testlist()
	n := l.add()
	defer l.remove(n)

	assert.Equal(t, 1, len(l.notes))
	assert.Equal(t, n, l.notes[0])
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

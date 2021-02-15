package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
)

func testGUI() *ui {
	str1 := "1"
	str2 := "2"
	l := &notelist{pref: test.NewApp().Preferences(),
		notes: []*note{
			{content: binding.BindString(&str1)},
			{content: binding.BindString(&str2)},
		}}

	gui := &ui{notes: l}
	_ = gui.loadUI()
	return gui
}

func TestUIList(t *testing.T) {
	gui := testGUI()

	assert.Equal(t, 2, gui.list.Length())
}

func TestUIList_TapSetsContent(t *testing.T) {
	gui := testGUI()

	assert.Equal(t, "1", gui.content.Text)

	gui.list.Select(1) // this happens on a goroutine
	time.Sleep(time.Millisecond * 100)
	assert.Equal(t, "2", gui.content.Text)
}

func TestUIAdd(t *testing.T) {
	gui := testGUI()
	gui.addNote()

	assert.Equal(t, 3, gui.list.Length())
}

func TestUIRemove(t *testing.T) {
	gui := testGUI()
	gui.removeCurrentNote()

	assert.Equal(t, 1, gui.list.Length())
}

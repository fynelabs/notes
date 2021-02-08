package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
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

	assert.Equal(t, 2, len(gui.list.Objects))
}

func TestUIList_TapSetsContent(t *testing.T) {
	gui := testGUI()

	assert.Equal(t, "1", gui.content.Text)

	test.Tap(gui.list.Objects[1].(*widget.Button))
	assert.Equal(t, "2", gui.content.Text)
}

func TestUIAdd(t *testing.T) {
	gui := testGUI()
	gui.addNote()

	assert.Equal(t, 3, len(gui.list.Objects))
}

func TestUIRemove(t *testing.T) {
	gui := testGUI()
	gui.removeCurrentNote()

	assert.Equal(t, 1, len(gui.list.Objects))
}

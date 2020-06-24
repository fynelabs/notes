package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"fyne.io/fyne/test"
	"fyne.io/fyne/widget"
)

func testGUI() *ui {
	l := &notelist{pref: test.NewApp().Preferences(),
		notes: []*note{
			&note{content: "1"},
			&note{content: "2"},
		}}

	gui := &ui{notes: l}
	_ = gui.loadUI()
	return gui
}

func TestUIList(t *testing.T) {
	gui := testGUI()

	assert.Equal(t, 2, len(gui.list.Children))
}

func TestUIList_TapSetsContent(t *testing.T) {
	gui := testGUI()

	assert.Equal(t, "1", gui.content.Text)

	test.Tap(gui.list.Children[1].(*widget.Button))
	assert.Equal(t, "2", gui.content.Text)
}

func TestUIAdd(t *testing.T) {
	gui := testGUI()
	gui.addNote()

	assert.Equal(t, 3, len(gui.list.Children))
}

func TestUIRemove(t *testing.T) {
	gui := testGUI()
	gui.removeCurrentNote()

	assert.Equal(t, 1, len(gui.list.Children))
}

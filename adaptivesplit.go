package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func newAdaptiveSplit(left, right fyne.CanvasObject) *fyne.Container {
	split := container.NewHSplit(left, right)
	split.Offset = 0.25
	return container.New(&adaptiveLayout{split: split}, split)
}

type adaptiveLayout struct {
	split *container.Split
}

func (a *adaptiveLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	dev := fyne.CurrentDevice()

	a.split.Horizontal = !dev.IsMobile() || fyne.IsHorizontal(dev.Orientation())
	objects[0].Resize(size)
}

func (a *adaptiveLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return a.split.MinSize()
}

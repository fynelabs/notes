package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type adaptiveSplit struct {
	widget.SplitContainer
}

func (a *adaptiveSplit) Resize(s fyne.Size) {
	dev := fyne.CurrentDevice()

	a.Horizontal = !dev.IsMobile() || fyne.IsHorizontal(dev.Orientation())
	a.SplitContainer.Resize(s)
}

func newAdaptiveSplit(left, right fyne.CanvasObject) *adaptiveSplit {
	split := &adaptiveSplit{}
	split.Leading = left
	split.Trailing = right
	split.ExtendBaseWidget(split)
	return split
}

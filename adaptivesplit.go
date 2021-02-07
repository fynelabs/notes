package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
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

func newAdaptiveSplit(left, right fyne.CanvasObject) *fyne.Container {
	split := container.NewHSplit(left, right)
	split.Offset = 0.25
	return fyne.NewContainerWithLayout(&adaptiveLayout{split: split}, split)
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

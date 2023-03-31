//go:generate fyne bundle -o bundled.go GochiHand.ttf
//go:generate fyne bundle -append -o bundled.go Icon.png

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct {
}

func (m *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground, theme.ColorNameInputBackground,
		theme.ColorNameOverlayBackground, theme.ColorNameMenuBackground:
		if v == theme.VariantLight {
			return &color.NRGBA{R: 0xF0, G: 0xE9, B: 0x9B, A: 0xFF}
		}
		return &color.NRGBA{R: 0x37, G: 0x2B, B: 0x09, A: 0xFF}
	case theme.ColorNameForeground:
		if v == theme.VariantLight {
			return &color.NRGBA{R: 0x46, G: 0x3A, B: 0x11, A: 0xFF}
		}
		return &color.NRGBA{R: 0xF0, G: 0xE9, B: 0x9B, A: 0xFF}
	case theme.ColorNamePrimary:
		return &color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xAA}
	case theme.ColorNameButton, theme.ColorNameFocus, theme.ColorNameSelection:
		return &color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	}

	return theme.DefaultTheme().Color(n, v)
}

func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource {
	return resourceGochiHandTtf
}

func (m *myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *myTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNameLineSpacing:
		return 2
	case theme.SizeNameText:
		return theme.DefaultTheme().Size(n) + 4
	}

	return theme.DefaultTheme().Size(n)
}

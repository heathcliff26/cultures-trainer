package app

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	borderShadowVariantLight = color.Black
	borderShadowVariantDark  = color.RGBA{120, 120, 120, ^uint8(0)}
)

var _ fyne.Theme = borderTheme{}

type borderTheme struct{}

func (borderTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (borderTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (borderTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (borderTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameShadow {
		switch variant {
		case theme.VariantLight:
			return borderShadowVariantLight
		case theme.VariantDark:
			return borderShadowVariantDark
		}
	}
	return theme.DefaultTheme().Color(name, variant)
}

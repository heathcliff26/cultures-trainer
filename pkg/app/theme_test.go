package app

import (
	"testing"

	"fyne.io/fyne/v2/theme"
	"github.com/stretchr/testify/assert"
)

func TestBorderTheme(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(borderShadowVariantLight, borderTheme{}.Color(theme.ColorNameShadow, theme.VariantLight))
	assert.Equal(borderShadowVariantDark, borderTheme{}.Color(theme.ColorNameShadow, theme.VariantDark))

	assert.Equal(theme.DefaultTheme().Color(theme.ColorNameBackground, theme.VariantLight), borderTheme{}.Color(theme.ColorNameBackground, theme.VariantLight))
	assert.Equal(theme.DefaultTheme().Color(theme.ColorNameBackground, theme.VariantDark), borderTheme{}.Color(theme.ColorNameBackground, theme.VariantDark))
	assert.Equal(theme.DefaultTheme().Icon(theme.IconNameAccount), borderTheme{}.Icon(theme.IconNameAccount))
	assert.Equal(theme.DefaultTheme().Size(theme.SizeNameSeparatorThickness), borderTheme{}.Size(theme.SizeNameSeparatorThickness))
}

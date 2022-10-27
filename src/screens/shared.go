package shared

import (
	"example/raylib-game/src/settings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Font rl.Font
var SecondaryFont rl.Font
var FxClick rl.Sound

const (
	FontSmallTextSize  float32 = 16
	FontMediumTextSize         = 24
	FontBigTextSize            = 32
	FontHugeTextSize           = 42
)

var AppSettings settings.Settings

// Logo variables
var LogoIcon rl.Texture2D
var IconRect rl.Rectangle
var TextRect rl.Rectangle

const (
	Logo int = iota
	Title
	Options
	Gameplay
	Ending
)

const Unchanged int = -1

// Load the shared assets
func LoadSharedAssets() {
	// Set up the font
	Font = rl.LoadFont("resources/fonts/montserrat_semibold.ttf")
	rl.GenTextureMipmaps(&Font.Texture)
	rl.SetTextureFilter(Font.Texture, rl.FilterBilinear)

	SecondaryFont = rl.LoadFont("resources/fonts/exo2_medium_italic.ttf")
	rl.GenTextureMipmaps(&SecondaryFont.Texture)
	rl.SetTextureFilter(SecondaryFont.Texture, rl.FilterBilinear)
	AppSettings.LoadFromFile()

	// Logo textures
	LogoIcon = rl.LoadTexture("resources/icons/logo_old.png")
	IconRect = rl.NewRectangle(30, 25, 45, 45)
	TextRect = rl.NewRectangle(82, 27, 250, 50)
}

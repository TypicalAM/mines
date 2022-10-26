package shared

import (
	"example/raylib-game/src/settings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Font rl.Font
var FxClick rl.Sound

const (
	FontSmallTextSize  float32 = 16
	FontMediumTextSize         = 24
	FontBigTextSize            = 32
)

var AppSettings settings.Settings

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
	rl.SetTextureFilter(Font.Texture, rl.FilterAnisotropic4x)
	AppSettings.LoadFromFile()
}

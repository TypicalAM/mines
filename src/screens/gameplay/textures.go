package gameplay

import (
	shared "example/raylib-game/src/screens"
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Adapt the numbers textures to the current colorscheme
func loadThemeTextureNums(filepath string) rl.Texture2D {
	data, err := shared.ResourcesFS.ReadFile(filepath)
	if err != nil {
		rl.TraceLog(rl.LogFatal, fmt.Sprint("Failed to load the texture: ", err))
	}

	img := rl.LoadImageFromMemory(".png", data, int32(len(data)))
	rl.ImageColorReplace(img, rl.NewColor(0, 0, 255, 255), rg.BackgroundColor())
	rl.ImageColorReplace(img, rl.NewColor(0, 255, 0, 255), rg.TextColor())
	rl.ImageColorReplace(img, rl.NewColor(255, 0, 0, 255), rg.GetStyleColor(rg.TogglePressedTextColor))
	rl.ImageColorReplace(img, rl.NewColor(255, 128, 0, 255), rg.GetStyleColor(rg.ButtonDefaultBorderColor))
	return rl.LoadTextureFromImage(img)
}

// Adapt the clock/bomb textures to the current colorscheme
func loadThemeTextureIcons(filepath string) rl.Texture2D {
	data, err := shared.ResourcesFS.ReadFile(filepath)
	if err != nil {
		rl.TraceLog(rl.LogFatal, fmt.Sprint("Failed to load the texture: ", err))
	}

	img := rl.LoadImageFromMemory(".png", data, int32(len(data)))
	rl.ImageColorReplace(img, rl.NewColor(0, 0, 255, 255), rg.BackgroundColor())
	rl.ImageColorReplace(img, rl.Black, rg.TextColor())
	return rl.LoadTextureFromImage(img)
}

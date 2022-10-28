package gameplay

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Adapt the numbers textures to the current colorscheme
func loadThemeTextureNums(filepath string) rl.Texture2D {
	// First, load the image from the texture
	img := rl.LoadImage(filepath)

	// Replace blue with the background color
	rl.ImageColorReplace(img, rl.NewColor(0, 0, 255, 255), rg.BackgroundColor())

	// Replace green with text color
	rl.ImageColorReplace(img, rl.NewColor(0, 255, 0, 255), rg.TextColor())

	// Replace red with button pressed background
	rl.ImageColorReplace(img, rl.NewColor(255, 0, 0, 255), rg.GetStyleColor(rg.TogglePressedTextColor))

	// Replace organge with uncovered background
	rl.ImageColorReplace(img, rl.NewColor(255, 128, 0, 255), rg.GetStyleColor(rg.ButtonDefaultBorderColor))

	return rl.LoadTextureFromImage(img)
}

// Adapt the clock/bomb textures to the current colorscheme
func loadThemeTextureIcons(filepath string) rl.Texture2D {
	// Load img from file
	img := rl.LoadImage(filepath)

	// Replace the blue color with the background
	rl.ImageColorReplace(img, rl.NewColor(0, 0, 255, 255), rg.BackgroundColor())

	// Replace the black color with the text color
	rl.ImageColorReplace(img, rl.Black, rg.TextColor())
	fmt.Println("test: ", filepath, rl.GetImageColor(*img, 24, 24))

	return rl.LoadTextureFromImage(img)
}

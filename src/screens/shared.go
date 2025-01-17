package shared

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TypicalAM/mines/src/settings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ResourcesFS embed.FS

var Font rl.Font
var SecondaryFont rl.Font

const (
	FontSmallTextSize  float32 = 16
	FontMediumTextSize float32 = 24
	FontBigTextSize    float32 = 32
	FontHugeTextSize   float32 = 42
)

var AppSettings settings.Settings
var Scores settings.Scores
var Themes []string

var LogoIcon rl.Texture2D
var IconRect rl.Rectangle
var TextRect rl.Rectangle

var gamepadButtonCooldown float32

const (
	Logo int = iota
	Title
	Options
	Gameplay
	Ending
	Leaderboard
)

const (
	ButtonUnchanged int = iota
	ButtonConfirm
	ButtonGoBack
	ButtonRestart
	ButtonFlag
	ButtonUp
	ButtonDown
	ButtonLeft
	ButtonRight
)

const Unchanged int = -1

// Load the shared assets
func LoadSharedAssets() error {
	if err := writeThemesToConfig(); err != nil {
		return err
	}

	if err := loadFonts(); err != nil {
		return err
	}

	// Iterate over the user themes
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	// Try to load the themes
	files, err := os.ReadDir(filepath.Join(cfgDir, "gomines", "themes"))
	if err != nil {
		return err
	}

	for _, file := range files {
		splitName := strings.Split(file.Name(), ".style")
		if !file.IsDir() && len(splitName) == 2 {
			Themes = append(Themes, splitName[0])
		}
	}

	// Check if we actually have any themes
	if len(Themes) == 0 {
		return errors.New("there are no available themes")
	}

	rl.TraceLog(rl.LogInfo, fmt.Sprintf("Loaded themes: %v", Themes))

	// Load the necessary settings and scores
	if err := AppSettings.LoadFromFile(Themes[0]); err != nil {
		return err
	}

	if err := Scores.LoadFromFile(); err != nil {
		return err
	}

	// Logo textures
	data, err := ResourcesFS.ReadFile("resources/icons/logo.png")
	if err != nil {
		return err
	}

	LogoIcon = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", data, int32(len(data))))
	IconRect = rl.NewRectangle(30, 25, 45, 45)
	TextRect = rl.NewRectangle(82, 27, 250, 50)

	return nil
}

// loadFonts tries to load the fonts from the embed filesystem
func loadFonts() error {
	data, err := ResourcesFS.ReadFile("resources/fonts/montserrat_semibold.ttf")
	if err != nil {
		return err
	}

	// IMPORTANT: Since we are using embed we need to load the font from memory and frankly, I have no clue how to
	// figure out the font size/char info, let's just assume the users don't use any weird characters
	Font = rl.LoadFontFromMemory(".ttf", data, int32(len(data)), 100, nil, 100)
	rl.GenTextureMipmaps(&Font.Texture)
	rl.SetTextureFilter(Font.Texture, rl.FilterBilinear)

	data, err = ResourcesFS.ReadFile("resources/fonts/cartograph_cf_italic.ttf")
	if err != nil {
		return err
	}

	SecondaryFont = rl.LoadFontFromMemory(".ttf", data, int32(len(data)), 100, nil, 100)
	rl.GenTextureMipmaps(&SecondaryFont.Texture)
	rl.SetTextureFilter(SecondaryFont.Texture, rl.FilterBilinear)
	return nil
}

// A function used to navigate the UI using keyboard buttons
func UpdateMovement(current int, availableButtons int) (int, int) {
	if gamepadButtonCooldown <= 0.0 {
		if rl.GetGamepadButtonPressed() != -1 {
			gamepadButtonCooldown = 0.1
		}

		switch rl.GetGamepadButtonPressed() {
		case 3: // PS3 gamepad down
			current++
			if current == availableButtons {
				current = 0
			}
			return current, ButtonDown
		case 1: // PS3 gamepad up
			current--
			if current == -1 {
				current = availableButtons - 1
			}
			return current, ButtonUp
		case 8: // PS3 gamepad flag
			return current, ButtonFlag
		case 15: // PS3 gamepad restart
			return current, ButtonRestart
		case 4: // PS3 gamepad left
			return current, ButtonLeft
		case 2: // PS3 gamepad right
			return current, ButtonRight
		case 7: // PS3 gamepad confirm
			return current, ButtonConfirm
		case 6: // PS3 gamepad go back
			return current, ButtonGoBack
		}
	} else {
		gamepadButtonCooldown -= 0.01
	}

	switch rl.GetKeyPressed() {
	case rl.KeyDown, rl.KeyTab:
		current++
		if current == availableButtons {
			current = 0
		}
		return current, ButtonDown
	case rl.KeyUp:
		current--
		if current == -1 {
			current = availableButtons - 1
		}
		return current, ButtonUp
	case rl.KeyF:
		return current, ButtonFlag
	case rl.KeyR:
		return current, ButtonRestart
	case rl.KeyLeft:
		return current, ButtonLeft
	case rl.KeyRight:
		return current, ButtonRight
	case rl.KeyEnter:
		return current, ButtonConfirm
	case rl.KeyEscape:
		return current, ButtonGoBack
	}

	return current, ButtonUnchanged
}

// write all the themes to the cfg dir
func writeThemesToConfig() error {
	// Write all the themes to the cfg dir
	builtin, err := ResourcesFS.ReadDir("resources/themes")
	if err != nil {
		return err
	}

	cfgDir, err := os.UserConfigDir()
	rl.TraceLog(rl.LogInfo, fmt.Sprintf("Writing themes to %s", cfgDir))
	if err != nil {
		return err
	}

	themeDir := filepath.Join(cfgDir, "gomines", "themes")
	if err := os.MkdirAll(themeDir, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	for _, file := range builtin {
		data, err := ResourcesFS.ReadFile("resources/themes/" + file.Name())
		if err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(themeDir, file.Name()), data, 0644); err != nil {
			return err
		}
	}

	return nil
}

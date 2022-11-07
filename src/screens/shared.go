package shared

import (
	"errors"
	"example/raylib-game/src/settings"
	"fmt"
	"io/ioutil"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Font rl.Font
var SecondaryFont rl.Font
var FxClick rl.Sound

const (
	FontSmallTextSize  float32 = 16
	FontMediumTextSize float32 = 24
	FontBigTextSize    float32 = 32
	FontHugeTextSize   float32 = 42
)

var AppSettings settings.Settings
var Scores settings.Scores
var Themes []string

// Logo variables
var LogoIcon rl.Texture2D
var IconRect rl.Rectangle
var TextRect rl.Rectangle

// Gamepad variables
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
	// Set up the font
	Font = rl.LoadFont("resources/fonts/montserrat_semibold.ttf")
	rl.GenTextureMipmaps(&Font.Texture)
	rl.SetTextureFilter(Font.Texture, rl.FilterBilinear)

	SecondaryFont = rl.LoadFont("resources/fonts/cartograph_cf_italic.ttf")
	rl.GenTextureMipmaps(&SecondaryFont.Texture)
	rl.SetTextureFilter(SecondaryFont.Texture, rl.FilterBilinear)

	// Iterate over the files and add them to the themes variable
	files, err := ioutil.ReadDir("resources/themes")
	if err != nil {
		return err
	}

	// Iterate over the files
	for _, file := range files {
		splitName := strings.Split(file.Name(), ".style")
		if !file.IsDir() && len(splitName) == 2 {
			Themes = append(Themes, splitName[0])
		}
	}

	// Check if we actually have any themes
	if len(Themes) == 0 {
		return errors.New("there are no available themes")
	} else {
		rl.TraceLog(rl.LogInfo, fmt.Sprintf("Loaded themes: %v", Themes))
	}

	// Load the necessary settings and scores
	if err := AppSettings.LoadFromFile(Themes[0]); err != nil {
		return err
	}

	if err := Scores.LoadFromFile(); err != nil {
		return err
	}

	// Logo textures
	LogoIcon = rl.LoadTexture("resources/icons/logo_old.png")
	IconRect = rl.NewRectangle(30, 25, 45, 45)
	TextRect = rl.NewRectangle(82, 27, 250, 50)

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

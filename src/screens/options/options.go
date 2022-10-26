package options

import (
	"example/raylib-game/src/gui"
	"example/raylib-game/src/screens"
	"example/raylib-game/src/settings"
	"fmt"
	"strconv"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables
var framesCounter int32
var finishScreen int

const (
	TextBox int = iota
	Slider
	Combo
)

// Option structure
type option struct {
	optionRectangle rl.Rectangle
	textRectangle   rl.Rectangle
	value           string
	valType         int
}

// Variable for all our options
var options map[string]option

// Can we save?
var saveAndExit bool
var saveErrorMessage string
var saveWrongData bool
var saveOptionsRect rl.Rectangle

// Options screen initialization logic
func InitOptionsScreen() {
	// Basic variables
	framesCounter = 0
	finishScreen = shared.Unchanged

	// Make the buttons take up 1/3rd of the screen
	rectangleWidths := float32(rl.GetScreenWidth()) / 3
	rectangleXPos := (float32(rl.GetScreenWidth()) - rectangleWidths) / 2

	options = map[string]option{
		"width": {
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2-160), rectangleWidths, 60),
			textRectangle: rl.NewRectangle(
				rectangleXPos+rectangleWidths/2-40,
				float32(rl.GetScreenHeight()/2-185),
				rectangleWidths, 60),
			value:   fmt.Sprint(shared.AppSettings.Width),
			valType: Slider,
		},
		"height": {
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2-160+95), rectangleWidths, 60),
			textRectangle:   rl.NewRectangle(rectangleXPos+rectangleWidths/2-40, float32(rl.GetScreenHeight()/2-185+95), rectangleWidths, 60),
			value:           fmt.Sprint(shared.AppSettings.Height),
			valType:         Slider,
		},
		"bombs": {
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2-160+190), rectangleWidths, 60),
			textRectangle:   rl.NewRectangle(rectangleXPos+rectangleWidths/2-40, float32(rl.GetScreenHeight()/2-185+190), rectangleWidths, 60),
			value:           fmt.Sprint(shared.AppSettings.Bombs),
			valType:         Slider,
		},
		"path": {
			value:           shared.AppSettings.SettingsPath,
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2-160+285), rectangleWidths, 60),
			textRectangle:   rl.NewRectangle(rectangleXPos+rectangleWidths/2-40, float32(rl.GetScreenHeight()/2-185+285), rectangleWidths, 60),
			valType:         TextBox,
		},
	}

	// Init save rectangle
	saveOptionsRect = rl.NewRectangle(
		rectangleXPos+rectangleWidths/4,
		float32(rl.GetScreenHeight()/2+235),
		rectangleWidths/2, 60,
	)
}

// Update Options screen
func UpdateOptionsScreen() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		finishScreen = shared.Title
	}

	if saveAndExit {
		saveResults()
	}
}

// Options screen draw logic
func DrawOptionsScreen() {
	//value, _ := strconv.ParseInt("0277bd", 16, 64)
	//	color := rl.GetColor(uint(value))

	// Draw the background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rg.BackgroundColor())

	for key, opt := range options {
		var textToDisplay string

		switch opt.valType {
		case Slider:
			currentValue, _ := strconv.Atoi(opt.value)
			textToDisplay = fmt.Sprintf("%s : %d", key, currentValue)
			opt.value = fmt.Sprintf("%d", int(rg.Slider(opt.optionRectangle, float32(currentValue), 1, 100)))
		case TextBox:
			textToDisplay = fmt.Sprint(key)
			opt.value = gui.TextBoxEx(shared.Font, opt.optionRectangle, opt.value, 25)
		}

		options[key] = opt

		rl.DrawTextEx(shared.Font, textToDisplay, rl.Vector2{X: opt.textRectangle.X, Y: opt.textRectangle.Y}, shared.FontSmallTextSize, 0, rg.GetStyleColor(rg.TextboxTextColor))
	}

	//entry := options["bombs"]
	//rl.DrawRectangleRec(entry.optionRectangle, rl.White)

	if saveWrongData {
		gui.ButtonEx(shared.Font, saveOptionsRect, saveErrorMessage, shared.FontBigTextSize)
	} else {
		saveAndExit = gui.ButtonEx(shared.Font, saveOptionsRect, "SAVE", shared.FontBigTextSize)
	}
}

// Options screen unload logic
func UnloadOptionsScreen() {}

// Options screens should finish?
func FinishOptionsScreen() int {
	return finishScreen
}

// Get all the fields and save it
func saveResults() {
	width, _ := strconv.Atoi(options["width"].value)
	height, _ := strconv.Atoi(options["height"].value)
	bombs, _ := strconv.Atoi(options["bombs"].value)
	path := options["path"].value

	newSettings := settings.Settings{
		Width:        width,
		Height:       height,
		Bombs:        bombs,
		SettingsPath: path,
	}

	err := shared.AppSettings.WriteToFile(newSettings)
	if err == nil {
		finishScreen = shared.Title
	} else {
		saveWrongData = true
		saveAndExit = false
		saveErrorMessage = err.Error()
	}
}

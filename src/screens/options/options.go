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
	FilePicker int = iota
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
var saveWrongData bool
var saveOptionsRect rl.Rectangle

// What theme are we using?
var activeThemeIndex int
var themes = []string{"default_dark", "default_light", "candy","hello_kitty", "monokai", "obsidian", "solarized", "solarized_light", "zahnrad"}

// Options screen initialization logic
func InitOptionsScreen() {
	// Basic variables
	framesCounter = 0
	finishScreen = shared.Unchanged

	// Make the buttons take up 1/3rd of the screen
	rectangleWidths := float32(rl.GetScreenWidth()) / 3
	rectangleXPos := (float32(rl.GetScreenWidth()) - rectangleWidths) / 2

	baseRectY := -250
	baseTextY := -275
	baseOffsetY := 100

	options = map[string]option{
		"Width of the map": {
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY), rectangleWidths, 60),
			textRectangle:   rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Width of the map - 100", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY), rectangleWidths, 60),
			value:           fmt.Sprint(shared.AppSettings.Width),
			valType:         Slider,
		},
		"Height of the map": {
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+baseOffsetY), rectangleWidths, 60),
			textRectangle:   rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Height of the map - 100", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY+baseOffsetY), rectangleWidths, 60),
			value:           fmt.Sprint(shared.AppSettings.Height),
			valType:         Slider,
		},
		"Bombs count percentage": {
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+2*baseOffsetY), rectangleWidths, 60),
			textRectangle:   rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Bombs count percentage - 100", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY+2*baseOffsetY), rectangleWidths, 60),
			value:           fmt.Sprint(shared.AppSettings.Bombs),
			valType:         Slider,
		},
		"Path to the settings file": {
			value:           shared.AppSettings.SettingsPath,
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+3*baseOffsetY), rectangleWidths, 60),
			textRectangle:   rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Path to the settings file", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY+3*baseOffsetY), rectangleWidths, 60),
			valType:         FilePicker,
		},
		"Color scheme": {
			value:           shared.AppSettings.Theme,
			optionRectangle: rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+4*baseOffsetY), rectangleWidths, 60),
			textRectangle:   rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Color scheme", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY+4*baseOffsetY), rectangleWidths, 60),
			valType:         Combo,
		},
	}

	// Init save rectangle
	saveOptionsRect = rl.NewRectangle(
		rectangleXPos+rectangleWidths/4,
		float32(rl.GetScreenHeight()/2+250),
		rectangleWidths/2, 60,
	)

	// Active theme indexing
	for ind, theme := range themes {
		if theme == shared.AppSettings.Theme {
			activeThemeIndex = ind
		}
	}
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
	// Draw the background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rg.BackgroundColor())
	// Draw the logo
	gui.DrawLogoTopLeft(shared.LogoIcon, shared.SecondaryFont, shared.IconRect, shared.TextRect, shared.FontHugeTextSize)

	for key, opt := range options {
		var textToDisplay string

		switch opt.valType {
		case Slider:
			currentValue, _ := strconv.Atoi(opt.value)
			textToDisplay = fmt.Sprintf("%s : %d", key, currentValue)
			opt.value = fmt.Sprintf("%d", int(rg.Slider(opt.optionRectangle, float32(currentValue), 1, 100)))
		case FilePicker:
			textToDisplay = fmt.Sprint(key)
			opt.value = gui.FilePicker(shared.LogoIcon, shared.Font, opt.optionRectangle, opt.value, shared.FontBigTextSize)
		case Combo:
			textToDisplay = fmt.Sprint(key)
			activeThemeIndex = gui.ComboBoxEx(shared.Font, opt.optionRectangle, themes, activeThemeIndex, shared.FontBigTextSize)
			opt.value = fmt.Sprint(themes[activeThemeIndex])
			rg.LoadGuiStyle(fmt.Sprintf("resources/styles/%s.style", opt.value))
		}

		options[key] = opt

		rl.DrawTextEx(shared.Font, textToDisplay, rl.Vector2{X: opt.textRectangle.X, Y: opt.textRectangle.Y}, shared.FontSmallTextSize, 0, rg.GetStyleColor(rg.TextboxTextColor))
	}

	saveAndExit = gui.ButtonEx(shared.Font, saveOptionsRect, "SAVE", shared.FontBigTextSize)
	if saveWrongData {
		rl.DrawRectangleRec(saveOptionsRect, rl.Fade(rl.Red, 0.50))
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
	width, _ := strconv.Atoi(options["Width of the map"].value)
	height, _ := strconv.Atoi(options["Height of the map"].value)
	bombs, _ := strconv.Atoi(options["Bombs count percentage"].value)
	path := options["Path to the settings file"].value
	theme := options["Color scheme"].value

	newSettings := settings.Settings{
		Width:        width,
		Height:       height,
		Bombs:        bombs,
		SettingsPath: path,
		Theme:        theme,
	}

	err := shared.AppSettings.WriteToFile(newSettings)
	if err == nil {
		finishScreen = shared.Title
	} else {
		saveWrongData = true
		saveAndExit = false
	}
}

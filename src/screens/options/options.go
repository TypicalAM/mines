package options

import (
	"example/raylib-game/src/gui"
	shared "example/raylib-game/src/screens"
	"example/raylib-game/src/settings"
	"fmt"
	"strconv"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables
var ScreenState int

const (
	FilePicker int = iota
	Slider
	Combo
)

// Option structure
type option struct {
	bounds     rl.Rectangle
	textBounds rl.Rectangle
	value      string
	valType    int
}

// A map for all our options with described key names
const keyWidth string = "Width of the map"
const keyHeight string = "Height of the map"
const keyBombs string = "Bombs count percentage"
const keySettingsPath string = "Path to the settings file"
const keyColorscheme string = "Color scheme"

var options map[string]option

// Make the preset variables
type preset struct {
	bounds    rl.Rectangle
	isPressed bool
	name      string
	width     int
	height    int
	bombs     int
}

var presets [3]preset

// Can we save?
var saveAndExit bool
var saveWrongData bool
var saveOptionsRect rl.Rectangle

// What theme are we using?
var activeThemeIndex int
var themes = []string{"default_dark", "default_light", "candy", "hello_kitty", "monokai", "obsidian", "solarized", "solarized_light", "zahnrad"}

// Options screen initialization logic
func Init() {
	// Basic variables
	ScreenState = shared.Unchanged

	// Make the buttons take up 1/3rd of the screen
	rectangleWidths := float32(rl.GetScreenWidth()) / 3
	rectangleXPos := (float32(rl.GetScreenWidth()) - rectangleWidths) / 2

	baseRectY := -250
	baseTextY := -275
	baseOffsetY := 100

	// Create the preset rectangles
	presets[0] = preset{
		bounds:    rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY-baseOffsetY), rectangleWidths/3, 60),
		isPressed: false,
		name:      "Beginner",
		width:     8,
		height:    8,
		bombs:     15,
	}
	presets[1] = preset{
		bounds:    rl.NewRectangle(rectangleXPos+float32(1*int(rectangleWidths)/3), float32(rl.GetScreenHeight()/2+baseRectY-baseOffsetY), rectangleWidths/3, 60),
		isPressed: false,
		name:      "Intermediate",
		width:     16,
		height:    16,
		bombs:     15,
	}
	presets[2] = preset{
		bounds:    rl.NewRectangle(rectangleXPos+float32(2*int(rectangleWidths)/3), float32(rl.GetScreenHeight()/2+baseRectY-baseOffsetY), rectangleWidths/3, 60),
		isPressed: false,
		name:      "Advanced",
		width:     30,
		height:    16,
		bombs:     21,
	}

	// Create the options
	options = map[string]option{
		keyWidth: {
			bounds:     rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY), rectangleWidths, 60),
			textBounds: rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Width of the map - 100", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY), rectangleWidths, 60),
			value:      fmt.Sprint(shared.AppSettings.Width),
			valType:    Slider,
		},
		keyHeight: {
			bounds:     rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+baseOffsetY), rectangleWidths, 60),
			textBounds: rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Height of the map - 100", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY+baseOffsetY), rectangleWidths, 60),
			value:      fmt.Sprint(shared.AppSettings.Height),
			valType:    Slider,
		},
		keyBombs: {
			bounds:     rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+2*baseOffsetY), rectangleWidths, 60),
			textBounds: rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Bombs count percentage - 100", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY+2*baseOffsetY), rectangleWidths, 60),
			value:      fmt.Sprint(shared.AppSettings.Bombs),
			valType:    Slider,
		},
		"Path to the settings file": {
			value:      shared.AppSettings.SettingsPath,
			bounds:     rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+3*baseOffsetY), rectangleWidths, 60),
			textBounds: rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Path to the settings file", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY+3*baseOffsetY), rectangleWidths, 60),
			valType:    FilePicker,
		},
		"Color scheme": {
			value:      shared.AppSettings.Theme,
			bounds:     rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+4*baseOffsetY), rectangleWidths, 60),
			textBounds: rl.NewRectangle(rectangleXPos+rectangleWidths/2-rl.MeasureTextEx(shared.Font, "Color scheme", shared.FontSmallTextSize, 0).X/2, float32(rl.GetScreenHeight()/2+baseTextY+4*baseOffsetY), rectangleWidths, 60),
			valType:    Combo,
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
func Update() {
	// Go back to title
	if rl.IsKeyPressed(rl.KeyEscape) {
		ScreenState = shared.Title
	}

	// Update the sliders if the preset was pressed
	for _, preset := range presets {
		if preset.isPressed {
			widthEntry := options[keyWidth]
			widthEntry.value = fmt.Sprint(preset.width)
			options[keyWidth] = widthEntry
			heightEntry := options[keyHeight]
			heightEntry.value = fmt.Sprint(preset.height)
			options[keyHeight] = heightEntry
			bombsEntry := options[keyBombs]
			bombsEntry.value = fmt.Sprint(preset.bombs)
			options[keyBombs] = bombsEntry
		}
	}

	// Save results if we pressed the save results button
	if saveAndExit {
		saveResults()
	}
}

// Options screen draw logic
func Draw() {
	// Draw the background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rg.BackgroundColor())
	// Draw the logo
	gui.DrawLogoTopLeft(shared.LogoIcon, shared.SecondaryFont, shared.IconRect, shared.TextRect, shared.FontHugeTextSize)

	// Draw the preset rectangles
	for pos, preset := range presets {
		presets[pos].isPressed = gui.ButtonEx(shared.Font, preset.bounds, preset.name, shared.FontMediumTextSize)
	}

	for key, opt := range options {
		var textToDisplay string

		switch opt.valType {
		case Slider:
			currentValue, _ := strconv.Atoi(opt.value)
			textToDisplay = fmt.Sprintf("%s : %d", key, currentValue)
			opt.value = fmt.Sprintf("%d", int(rg.Slider(opt.bounds, float32(currentValue), 1, 100)))
		case FilePicker:
			textToDisplay = fmt.Sprint(key)
			opt.value = gui.FilePicker(shared.LogoIcon, shared.Font, opt.bounds, opt.value, shared.FontBigTextSize)
		case Combo:
			textToDisplay = fmt.Sprint(key)
			activeThemeIndex = gui.ComboBoxEx(shared.Font, opt.bounds, themes, activeThemeIndex, shared.FontBigTextSize)
			opt.value = fmt.Sprint(themes[activeThemeIndex])
			rg.LoadGuiStyle(fmt.Sprintf("resources/styles/%s.style", opt.value))
		}

		options[key] = opt

		rl.DrawTextEx(shared.Font, textToDisplay, rl.Vector2{X: opt.textBounds.X, Y: opt.textBounds.Y}, shared.FontSmallTextSize, 0, rg.GetStyleColor(rg.TextboxTextColor))
	}

	saveAndExit = gui.ButtonEx(shared.Font, saveOptionsRect, "SAVE", shared.FontBigTextSize)
	if saveWrongData {
		rl.DrawRectangleRec(saveOptionsRect, rl.Fade(rl.Red, 0.50))
	}
}

// Get all the fields and save it
func saveResults() {
	width, _ := strconv.Atoi(options[keyWidth].value)
	height, _ := strconv.Atoi(options[keyHeight].value)
	bombs, _ := strconv.Atoi(options[keyBombs].value)
	path := options[keySettingsPath].value
	theme := options[keyColorscheme].value

	newSettings := settings.Settings{
		Width:        width,
		Height:       height,
		Bombs:        bombs,
		SettingsPath: path,
		Theme:        theme,
	}

	err := shared.AppSettings.WriteToFile(newSettings)
	if err == nil {
		ScreenState = shared.Title
	} else {
		saveWrongData = true
		saveAndExit = false
	}
}

// Options screen unload logic
func Unload() {}

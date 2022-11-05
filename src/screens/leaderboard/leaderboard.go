package leaderboard

import (
	"example/raylib-game/src/gui"
	shared "example/raylib-game/src/screens"
	"example/raylib-game/src/settings"
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var ScreenState = shared.Unchanged
var col1 []rl.Rectangle
var col2 []rl.Rectangle
var col3 []rl.Rectangle

var categories = []string{"Beginner", "Intemediate", "Expert", "Custom"}
var categoriesBounds rl.Rectangle
var activeCategory int
var availableScores []settings.Entry
var activeCategoryChanged bool

var doneRectButton rl.Rectangle
var textVector rl.Vector2

// Init leaderboard screen
func Init() {
	ScreenState = shared.Unchanged

	// Leaderboard logo
	titleSize := rl.MeasureTextEx(shared.Font, "Leaderboard", shared.FontHugeTextSize*1.5, 0)
	textVector = rl.Vector2{
		X: float32(rl.GetScreenWidth()/2) - titleSize.X/2,
		Y: float32(rl.GetScreenHeight()) / 35,
	}

	// Scoreboard categories
	activeCategory = 0
	availableScores = shared.Scores.Entries
	activeCategoryChanged = false

	// Make the categories bounds
	categoryWidth := float32(rl.GetScreenWidth() / 8)
	categoryXPos := float32(rl.GetScreenWidth())/2 - categoryWidth*2
	categoriesBounds = rl.NewRectangle(categoryXPos, 100, categoryWidth, 50)

	col1 = make([]rl.Rectangle, len(availableScores)+1)
	col2 = make([]rl.Rectangle, len(availableScores)+1)
	col3 = make([]rl.Rectangle, len(availableScores)+1)

	width := float32(rl.GetScreenWidth() / 6)

	// Make the header elements
	col1[0] = rl.NewRectangle(float32(rl.GetScreenWidth())/2-width*3/2, 100+70, width, 60)
	col2[0] = rl.NewRectangle(float32(rl.GetScreenWidth())/2-width/2, 100+70, width, 60)
	col3[0] = rl.NewRectangle(float32(rl.GetScreenWidth())/2+width/2, 100+70, width, 60)

	// Done rectangle
	width = float32(rl.GetScreenWidth() / 3)
	xPos := float32(rl.GetScreenWidth())/2 - width/2
	doneRectButton = rl.NewRectangle(
		xPos+width/4, float32(rl.GetScreenHeight()-100), width/2, 60,
	)
}

// Update the screen
func Update() {
	// Make the width of the elements
	width := float32(rl.GetScreenWidth() / 6)

	// Make the entry elements
	for pos := range availableScores {
		col1[pos+1] = rl.NewRectangle(col1[pos].X, col1[pos].Y+col1[pos].Height, width, 60)
		col2[pos+1] = rl.NewRectangle(col2[pos].X, col2[pos].Y+col2[pos].Height, width, 60)
		col3[pos+1] = rl.NewRectangle(col3[pos].X, col3[pos].Y+col3[pos].Height, width, 60)
	}

	// Check the current categorie
	if activeCategoryChanged {
		availableScores = shared.Scores.FilterScores(activeCategory + 1)
		activeCategoryChanged = false
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		ScreenState = shared.Title
	}
}

// Draw the screen
func Draw() {
	// Draw the background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rg.BackgroundColor())
	// Draw the logo
	gui.DrawLogoTopLeft(shared.LogoIcon, shared.SecondaryFont, shared.IconRect, shared.TextRect, shared.FontHugeTextSize)
	// Draw the title
	rl.DrawTextEx(shared.Font, "Leaderboard", textVector, shared.FontHugeTextSize*1.5, 0, rg.TextColor())

	// Get the active category and check if it has changed
	newActiveCategory := gui.ToggleGroupEx(shared.Font, categoriesBounds, categories, activeCategory, shared.FontBigTextSize)
	if newActiveCategory != activeCategory {
		activeCategory = newActiveCategory
		activeCategoryChanged = true
	}

	rl.DrawRectangleRec(col1[0], rl.Maroon)
	rl.DrawRectangleRec(col2[0], rl.Gray)
	rl.DrawRectangleRec(col3[0], rl.Beige)

	for pos, entry := range col1 {
		var displayedText string
		if pos == 0 {
			rl.DrawRectangleRec(entry, rl.Maroon)
			displayedText = "Name"
			rg.DrawBorderedRectangle(entry.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.GetStyleColor(rg.ButtonDefaultInsideColor))
		} else {
			displayedText = availableScores[pos-1].Name
			rg.DrawBorderedRectangle(entry.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.BackgroundColor())
		}

		entryTextSize := rl.MeasureTextEx(shared.Font, displayedText, shared.FontBigTextSize, 0)
		rl.DrawTextEx(
			shared.Font, displayedText, rl.Vector2{
				X: entry.X + entry.Width/2 - entryTextSize.X/2,
				Y: entry.Y + entry.Height/2 - entryTextSize.Y/2,
			}, shared.FontBigTextSize, 0, rg.TextColor(),
		)
	}
	for pos, entry := range col2 {
		var displayedText string
		if pos == 0 {
			displayedText = "Board"
			rg.DrawBorderedRectangle(entry.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.GetStyleColor(rg.ButtonDefaultInsideColor))
		} else {
			displayedText = fmt.Sprintf("%dx%d (%d%% mines)",
				availableScores[pos-1].BoardWidth,
				availableScores[pos-1].BoardHeight,
				availableScores[pos-1].BoardMines,
			)
			rg.DrawBorderedRectangle(entry.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.BackgroundColor())
		}

		entryTextSize := rl.MeasureTextEx(shared.Font, displayedText, shared.FontBigTextSize, 0)
		rl.DrawTextEx(
			shared.Font, displayedText, rl.Vector2{
				X: entry.X + entry.Width/2 - entryTextSize.X/2,
				Y: entry.Y + entry.Height/2 - entryTextSize.Y/2,
			}, shared.FontBigTextSize, 0, rg.TextColor(),
		)
	}
	for pos, entry := range col3 {
		var displayedText string
		if pos == 0 {
			displayedText = "Time"
			rg.DrawBorderedRectangle(entry.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.GetStyleColor(rg.ButtonDefaultInsideColor))
		} else {
			displayedText = fmt.Sprintf("%d:%d", availableScores[pos-1].Time/60, availableScores[pos-1].Time%60)
			rg.DrawBorderedRectangle(entry.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.BackgroundColor())
		}

		entryTextSize := rl.MeasureTextEx(shared.Font, displayedText, shared.FontBigTextSize, 0)
		rl.DrawTextEx(
			shared.Font, displayedText, rl.Vector2{
				X: entry.X + entry.Width/2 - entryTextSize.X/2,
				Y: entry.Y + entry.Height/2 - entryTextSize.Y/2,
			}, shared.FontBigTextSize, 0, rg.TextColor(),
		)
	}
	// Draw the "we are done" button
	if gui.ButtonEx(shared.Font, doneRectButton, "Done", shared.FontBigTextSize) {
		ScreenState = shared.Title
	}
}

// Unload textures
func Unload() {}

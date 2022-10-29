package leaderboard

import (
	"example/raylib-game/src/gui"
	shared "example/raylib-game/src/screens"
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var finishScreen = shared.Unchanged
var col1 []rl.Rectangle
var col2 []rl.Rectangle
var col3 []rl.Rectangle

var col1Container rl.Rectangle
var col2Container rl.Rectangle
var col3Container rl.Rectangle
//var entries = []Entry{	{"Adam", 10, 213},	{"Maciek", 10, 220},	{"Marek", 23, 220},	{"test", 234, 240},	{"bobo", 20, 250},	{"testet2", 20, 260},}
var doneRectButton rl.Rectangle
var textVector rl.Vector2

// Init leaderboard screen
func InitLeaderboardScreen() {
	finishScreen = shared.Unchanged

	// Leaderboard logo
	titleSize := rl.MeasureTextEx(shared.Font, "Leaderboard", shared.FontHugeTextSize*1.5, 0)
	textVector = rl.Vector2{
		X: float32(rl.GetScreenWidth()/2) - titleSize.X/2,
		Y: float32(rl.GetScreenHeight()) / 35,
	}

	col1 = make([]rl.Rectangle, len(shared.Scores.Entries)+1)
	col2 = make([]rl.Rectangle, len(shared.Scores.Entries)+1)
	col3 = make([]rl.Rectangle, len(shared.Scores.Entries)+1)

	width := float32(rl.GetScreenWidth() / 6)

	// Make the header elements
	col1[0] = rl.NewRectangle(float32(rl.GetScreenWidth())/2-width*3/2, 100, width, 60)
	col2[0] = rl.NewRectangle(float32(rl.GetScreenWidth())/2-width/2, 100, width, 60)
	col3[0] = rl.NewRectangle(float32(rl.GetScreenWidth())/2+width/2, 100, width, 60)

	// Make the entry elements
	for pos := range shared.Scores.Entries {
		col1[pos+1] = rl.NewRectangle(col1[pos].X, col1[pos].Y+col1[pos].Height, width, 60)
		col2[pos+1] = rl.NewRectangle(col2[pos].X, col2[pos].Y+col2[pos].Height, width, 60)
		col3[pos+1] = rl.NewRectangle(col3[pos].X, col3[pos].Y+col3[pos].Height, width, 60)
	}

	// Make the container elements
	col1Container = rl.NewRectangle(col1[0].X, col1[0].Y, col1[0].Width, col1[0].Height*float32(len(shared.Scores.Entries)+1))
	col2Container = rl.NewRectangle(col2[0].X, col2[0].Y, col2[0].Width, col2[0].Height*float32(len(shared.Scores.Entries)+1))
	col3Container = rl.NewRectangle(col3[0].X, col3[0].Y, col3[0].Width, col3[0].Height*float32(len(shared.Scores.Entries)+1))

	// Done rectangle
	width = float32(rl.GetScreenWidth() / 3)
	xPos := float32(rl.GetScreenWidth())/2 - width/2
	doneRectButton = rl.NewRectangle(
		xPos+width/4, float32(rl.GetScreenHeight()-100), width/2, 60,
	)
}

// Update the screen
func UpdateLeaderboardScreen() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		finishScreen = shared.Title
	}
}

// Draw the screen
func DrawLeaderboardScreen() {
	// Draw the background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rg.BackgroundColor())
	// Draw the logo
	gui.DrawLogoTopLeft(shared.LogoIcon, shared.SecondaryFont, shared.IconRect, shared.TextRect, shared.FontHugeTextSize)

	//rl.DrawRectangle(int32(rl.GetScreenWidth()/2), 0, 1, int32(rl.GetScreenHeight()), rl.White)

	// Draw the title
	rl.DrawTextEx(shared.Font, "Leaderboard", textVector, shared.FontHugeTextSize*1.5, 0, rg.TextColor())

	// Draw the "we are done" button
	if gui.ButtonEx(shared.Font, doneRectButton, "Done", shared.FontBigTextSize) {
		finishScreen = shared.Title
	}

	rg.DrawBorderedRectangle(col1Container.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.GetStyleColor(rg.ButtonDefaultInsideColor))
	rg.DrawBorderedRectangle(col2Container.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.GetStyleColor(rg.ButtonDefaultInsideColor))
	rg.DrawBorderedRectangle(col3Container.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.GetStyleColor(rg.ButtonDefaultInsideColor))
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
			displayedText = shared.Scores.Entries[pos-1].Name
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
			displayedText = "Score"
			rg.DrawBorderedRectangle(entry.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(rg.ButtonDefaultBorderColor), rg.GetStyleColor(rg.ButtonDefaultInsideColor))
		} else {
			displayedText = fmt.Sprint(shared.Scores.Entries[pos-1].Score)
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
			displayedText = fmt.Sprintf("%d:%d", shared.Scores.Entries[pos-1].Time/60, shared.Scores.Entries[pos-1].Time%60)
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
}

// Unload textures
func UnloadLeaderboardScreen() {}

// Do we finish?
func FinishLeaderboardScreen() int {
	return finishScreen
}

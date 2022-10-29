package leaderboard

import (
	"example/raylib-game/src/gui"
	shared "example/raylib-game/src/screens"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var finishScreen = shared.Unchanged

type Entry struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Time  int    `json:"time"`
}

var entries = []Entry{
	{"Adam", 10, 213},
	{"Maciek", 10, 220},
	{"Marek", 23, 220},
	{"test", 234, 240},
	{"bobo", 20, 250},
	{"testet2", 20, 260},
}

var doneRectButton rl.Rectangle
var textVector rl.Vector2

// Init leaderboard screen
func InitLeaderboardScreen() {
	finishScreen = shared.Unchanged

	width := rl.GetScreenWidth() / 3
	xPos := rl.GetScreenWidth()/2 - width/2
	doneRectButton = rl.NewRectangle(
		float32(xPos+width/4), float32(rl.GetScreenHeight()-100), float32(width/2), 60,
	)

	titleSize := rl.MeasureTextEx(shared.Font, "Leaderboard", shared.FontHugeTextSize*1.5, 0)
	textVector = rl.Vector2{
		X: float32(rl.GetScreenWidth()/2) - titleSize.X/2,
		Y: float32(rl.GetScreenHeight()) / 14,
	}
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

	rl.DrawRectangle(int32(rl.GetScreenWidth()/2), 0, 1, int32(rl.GetScreenHeight()), rl.White)

	// Draw the title
	rl.DrawTextEx(shared.Font, "Leaderboard", textVector, shared.FontHugeTextSize*1.5, 0, rg.TextColor())

	// Draw the "we are done" button
	if gui.ButtonEx(shared.Font, doneRectButton, "Done", shared.FontBigTextSize) {
		finishScreen = shared.Title
	}

	width := rl.GetScreenWidth() / 6
	col1 := rl.NewRectangle(
		float32(rl.GetScreenWidth())/2-float32(width/2)-float32(width),
		100,
		float32(width),
		60,
	)
	col2 := rl.NewRectangle(
		float32(rl.GetScreenWidth())/2-float32(width)/2,
		100, float32(width), 60,
	)
	col3 := rl.NewRectangle(
		float32(rl.GetScreenWidth())/2+float32(width)/2,
		100, float32(width), 60,
	)
	big := rl.NewRectangle(
		col1.X, col2.Y, 3*float32(width), col1.Height*10,
	)

	rg.DrawBorderedRectangle(
	big.ToInt32(),
		rg.GetStyle32(rg.ButtonBorderWidth),
		rl.GetColor(uint(rg.ButtonDefaultBorderColor)),
		rl.GetColor(uint(rg.ButtonDefaultInsideColor)),
	)
	rl.DrawRectangleRec(col1, rl.White)
	rl.DrawRectangleRec(col2, rl.Gray)
	rl.DrawRectangleRec(col3, rl.Maroon)
}

// Unload textures
func UnloadLeaderboardScreen() {}

// Do we finish?
func FinishLeaderboardScreen() int {
	return finishScreen
}

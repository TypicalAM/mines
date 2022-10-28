package title

import (
	"example/raylib-game/src/gui"
	"example/raylib-game/src/screens"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables
var framesCounter int32
var finishScreen int = shared.Unchanged

// Rectangles for buttons
var startGameRect rl.Rectangle
var leaderboardRect rl.Rectangle
var optionsRect rl.Rectangle

// Logo rectangle
var logoRectangle rl.Rectangle

// Title screen initialization logic
func InitTitleScreen() {
	// Basic variables
	framesCounter = 0
	finishScreen = shared.Unchanged

	// Make the buttons take up 1/3rd of the screen
	rectangleWidths := float32(rl.GetScreenWidth()) / 3
	rectangleXPos := (float32(rl.GetScreenWidth()) - rectangleWidths) / 2

	baseRectY := -250
	baseOffsetY := 100

	// Make the rectangles
	startGameRect = rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+2*baseOffsetY), rectangleWidths, 60)
	leaderboardRect = rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+3*baseOffsetY), rectangleWidths, 60)
	optionsRect = rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+4*baseOffsetY), rectangleWidths, 60)

	// Make the logo
	logoRectangle = rl.NewRectangle(float32(rl.GetScreenWidth())/4,50,float32(rl.GetScreenWidth()/2),175)
}

// Update title screen
func UpdateTitleScreen() {}

// Title screen draw logic
func DrawTitleScreen() {
	// Draw the background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rg.BackgroundColor())

	// Draw the logo
	rl.DrawTexturePro(shared.LogoIcon, rl.NewRectangle(0, 0, 512, 512), rl.NewRectangle(logoRectangle.X, logoRectangle.Y, logoRectangle.Height, logoRectangle.Height), rl.Vector2{}, 0, rl.White)
	rl.DrawTextEx(shared.SecondaryFont, "Mines-go", rl.Vector2{X: float32(rl.GetScreenWidth())/4 + logoRectangle.Height + 10, Y: 90}, 120, 3, rg.GetStyleColor(rg.LabelTextColor))

	// Draw the buttons
	if gui.ButtonEx(shared.Font, startGameRect, "Start the game", shared.FontBigTextSize) {
		finishScreen = shared.Gameplay
	}
	if gui.ButtonEx(shared.Font, leaderboardRect, "Leaderboards", shared.FontBigTextSize) {
		finishScreen = shared.Leaderboard
	}
	if gui.ButtonEx(shared.Font, optionsRect, "Options", shared.FontBigTextSize) {
		finishScreen = shared.Options
	}
}

// Title screen unload logic
func UnloadTitleScreen() {}

// Title shared hould finish
func FinishTitleScreen() int {
	return finishScreen
}

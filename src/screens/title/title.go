package title

import (
	"example/raylib-game/src/screens"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables

var framesCounter int32
var finishScreen int

// Title screen initialization logic
func InitTitleScreen() {
	framesCounter = 0
	finishScreen = shared.Unchanged
}

// Update title screen
func UpdateTitleScreen() {

	if rl.IsKeyPressed(rl.KeyEnter) || rl.IsGestureDetected(rl.GestureTap) {
		finishScreen = shared.Gameplay
	}
}

// Title screen draw logic
func DrawTitleScreen() {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Green)
	rl.DrawText("TITLE SCREEN", 20, 20, 15, rl.DarkGreen)
	rl.DrawText("Press enter or tap to jump to gameplay", 120, 220, 20, rl.DarkGreen)
}

// Title screen unload logic
func UnloadTitleScreen() {}

// Title shared hould finish
func FinishTitleScreen() int {
	return finishScreen
}

package ending

import (
	"example/raylib-game/src/screens"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables

var framesCounter int32
var finishScreen int

// Ending screen initialization logic
func InitEndingScreen() {
	framesCounter = 0
	finishScreen = shared.Unchanged
}

// Update Ending screen
func UpdateEndingScreen() {

	if rl.IsKeyPressed(rl.KeyEnter) || rl.IsGestureDetected(rl.GestureTap) {
		finishScreen = shared.Title
//		rl.PlaySound(game.FxClick)
	}
}

// Ending screen draw logic
func DrawEndingScreen() {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Blue)
	rl.DrawText("Ending SCREEN", 20, 20, 15, rl.DarkBlue)
	rl.DrawText("Press enter or tap to jump to gameplay", 120, 220, 20, rl.DarkBlue)
}

// Ending screen unload logic
func UnloadEndingScreen() {
	// TODO: Unload Ending screen variables here!
}

// Ending shared hould finish
func FinishEndingScreen() int {
	return finishScreen
}

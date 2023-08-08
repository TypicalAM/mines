package ending

import (
	shared "example/raylib-game/src/screens"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables

var ScreenState int

// Ending screen initialization logic
func Init() {
	ScreenState = shared.Unchanged
}

// Update Ending screen
func Update() {
	if rl.IsKeyPressed(rl.KeyEnter) || rl.IsGestureDetected(rl.GestureTap) {
		ScreenState = shared.Title
		//		rl.PlaySound(game.FxClick)
	}
}

// Ending screen draw logic
func Draw() {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Blue)
	rl.DrawText("Ending SCREEN", 20, 20, 15, rl.DarkBlue)
	rl.DrawText("Press enter or tap to jump to gameplay", 120, 220, 20, rl.DarkBlue)
}

// Ending screen unload logic
func Unload() {
	// TODO: Unload Ending screen variables here!
}

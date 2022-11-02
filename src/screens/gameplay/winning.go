package gameplay

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables
var alpha = 0.0
var animation = false

// Initialize the game losing screen
func initWinningScreen() {
	fmt.Println("We have won the game")
	animation = true
}

// Update the game winning screen
func updateWinningScreen() {
	if animation {
		alpha += 0.01
		if alpha >= 0.5 {
			animation = false
		}
	}
}

// Draw the game winning screen
func drawWinningScreen() {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()),
		rl.Fade(rg.BackgroundColor(), float32(alpha)),
	)
}

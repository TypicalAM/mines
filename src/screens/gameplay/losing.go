package gameplay

import (
	"example/raylib-game/src/mines"
	"example/raylib-game/src/screens"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Define local variables
var gameLost bool // Is the game lost?
var score float32 // What is the final score of the player

var bombTile rl.Rectangle // The tile which lost the user his game
var bombExplosion [11]rl.Texture2D // The bomb explosion animation frames
var explosionFrame int32 // Which explosion frame are we on?

// Initialize the game finish screen variables
func initFinishScreen() {
	// Finish the game and calculate the score
	isPlaying = false
	gameLost = true
	score = calculateScore()

	// Load the bomb explosion texutres
	for i := range bombExplosion {
		bombExplosion[i] = rl.LoadTexture(fmt.Sprintf("resources/icons/explosion/frame%d.png", i+1))
	}

	// Initialize the bomb explosion variables
	explosionFrame = 0

	// Uncover every bomb
	for row := range mineBoard.Board {
		for col := range mineBoard.Board[row] {
			if mineBoard.Board[row][col] == mines.Bomb {
				mineBoard.TileState[row][col] = mines.Uncovered
			}
		}
	}
}

// Unload the used files
func UnloadLose() {
	for i := range bombExplosion {
		rl.UnloadTexture(bombExplosion[i])
	}
}

// Calculate the score that the user had
func calculateScore() float32 {
	var bombsCaught int

	// Count the bombs caught (flagged or covered)
	for row := range mineBoard.TileState {
		for col := range mineBoard.TileState[row] {
			if mineBoard.TileState[row][col] == mines.Flagged && mineBoard.Board[row][col] == mines.Bomb {
				bombsCaught += 1
			}
		}
	}

	// Return the percentage of bombs caught
	return float32(bombsCaught) / float32(mineBoard.Mines)
}

// Game lost update logic
func updateGameLostScreen() {
	framesCounter++

	if framesCounter > 5 {
		explosionFrame++
		framesCounter = 0
	}

	if explosionFrame > 10 {
		explosionFrame = 0
	}
}

// Draw the game over screen
func drawGameLostScreen() {
	text := fmt.Sprintf("Your score is: %d", int(score*100))
	textSize := rl.MeasureTextEx(
		shared.Font,
		text,
		shared.FontBigTextSize,
		0,
	)

	rl.DrawTextEx(
		shared.Font,
		text,
		rl.Vector2{
			X: (float32(rl.GetScreenWidth()) - textSize.X) / 2,
			Y: float32(rl.GetScreenHeight()/2) - textSize.Y,
		},
		shared.FontBigTextSize,
		0, rl.Maroon,
	)

	rl.DrawTexture(
		bombExplosion[explosionFrame],
		bombTile.ToInt32().X-23, bombTile.ToInt32().Y-25,
		rl.White,
	)
}

package gameplay

import (
	"example/raylib-game/src/mines"
	shared "example/raylib-game/src/screens"
	"fmt"
	"time"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Define local variables
var framesCounter int32 // Frames counter
var finishScreen int    // Determines if the screen should finish

const numsTextureSize float32 = 56 // The texture width and height (56x56)
var numberTextures rl.Texture2D    // The textures of the numbres (1-8, flag, empty and uncovered)
var bombIconTexture rl.Texture2D   // The texture of the bomb icon above the board
var clockIconTexture rl.Texture2D  // The texture of the clock icon above the board

var width int
var height int
var boardRectangles [][]rl.Rectangle // Rectangles making the playing board on screen
var mineBoard mines.MineBoard        // The mines playing board

type hover struct {
	isHovered bool
	row       int
	col       int
}

var tileHoverState hover // If any tile was hovered, and which was hoverd
var timePlaying time.Time // Time of the first meaningful mouse press
var isPlaying bool        // If the player is in game
var gameWon bool 					// If the game is won

// Flags and bombs text placements
var flagsText string
var flagsTextXPos float32
var flagsIconXPos int32
var clockText string
var clockTextXPos float32
var clockIconXPos int32

// Gameplay screen initialization logic
func InitGameplayScreen() {
	// Init basic variables
	framesCounter = 0
	finishScreen = shared.Unchanged
	gameLost = false

	// Init timing variables
	timePlaying = time.Time{}
	isPlaying = false

	// Generate the mineboard
	width = shared.AppSettings.Width
	height = shared.AppSettings.Height
	mineBoard, _ = mines.GenerateBoard(width, height, shared.AppSettings.Bombs)

	// Determine the placement of the game grid
	cellWidth := ((rl.GetScreenWidth() - 120) / width)
	cellHeight := (rl.GetScreenHeight() - 120) / height

	backGroundStartX := float32(rl.GetScreenWidth()-cellWidth*width) / 2
	backGroundStartY := float32(rl.GetScreenHeight()-cellHeight*height) / 2

	// Fill the game grid with rectangles
	boardRectangles = make([][]rl.Rectangle, height)
	for i := range boardRectangles {
		boardRectangles[i] = make([]rl.Rectangle, width)
		for j := range boardRectangles[i] {
			boardRectangles[i][j] = rl.NewRectangle(
				backGroundStartX+float32(j*cellWidth),
				backGroundStartY+float32(i*cellHeight),
				float32(cellWidth),
				float32(cellHeight),
			)
		}
	}

	// Load the appropriate textures
	numberTextures = loadThemeTextureNums("resources/icons/color_coded_flags.png")
	bombIconTexture = loadThemeTextureIcons("resources/icons/bomb_small_icon.png")
	clockIconTexture = loadThemeTextureIcons("resources/icons/clock_small_icon.png")

	// Set the tile hover state
	tileHoverState = hover{
		isHovered: false,
		row:       0,
		col:       0,
	}

	// Flags and clock text
	flagsText = "0"
	flagsTextXPos = float32(rl.GetScreenWidth()/2 + 85)
	flagsIconXPos = int32(rl.GetScreenWidth()/2 + 40)
	clockText = "00:00"
	clockTextXPos = float32(rl.GetScreenWidth()/2 - 85)
	clockIconXPos = int32(rl.GetScreenWidth()/2 - 125)
}

// Gameplay screen update logic
func UpdateGameplayScreen() {
	if gameWon {
		updateWinningScreen()
	}

	if mineBoard.Flags == mineBoard.Mines && mineBoard.CheckIfWon() {
		gameWon = true
		isPlaying = false
		initWinningScreen()
	}

	if gameLost {
		updateGameLostScreen()
		return
	}

	// Increase the timer if we are playing
	if isPlaying {
		// Change the time and the flags text
		myTime := time.Time{}.Add(time.Since(timePlaying))
		clockText = fmt.Sprintf("%02d:%02d", myTime.Minute(), myTime.Second())
		flagsText = fmt.Sprint(mineBoard.Flags)
	}

	// Ensure that nothing is hovered right now
	tileHoverState.isHovered = false

	for row := range boardRectangles {
		for col, tile := range boardRectangles[row] {
			// Check if the mouse is at the tile
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), tile) {

				// Set the currently hovered tile as hovered
				tileHoverState = hover{
					isHovered: true,
					row:       row,
					col:       col,
				}

				// If we detect the left click then uncover the tile if its not flagged
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && mineBoard.TileState[row][col] != mines.Flagged {
					mineBoard.TileState[row][col] = mines.Uncovered
					if mineBoard.UncoverValues(true, row, col) {
						bombTile = tile
						initFinishScreen()
					}

					// Start the game timer
					if !isPlaying {
						isPlaying = true
						timePlaying = time.Now()
					}
				}

				// If we detect the right click then flag/unflag the tile
				if rl.IsMouseButtonPressed(rl.MouseRightButton) {
					// If the tile is flagged, unflag it and if it is covered, flag it
					if mineBoard.TileState[row][col] == mines.Flagged {
						mineBoard.TileState[row][col] = mines.Covered
						mineBoard.Flags--
					} else if mineBoard.TileState[row][col] == mines.Covered {
						mineBoard.TileState[row][col] = mines.Flagged
						mineBoard.Flags++
					}

					// Start the game timer
					if !isPlaying {
						isPlaying = true
						timePlaying = time.Now()
					}
				}
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		finishScreen = shared.Title
	}

}

// Gameplay screen draw logic
func DrawGameplayScreen() {
	// Draw the background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rg.BackgroundColor())

	// Draw the game board
	for row := range boardRectangles {
		for col, tile := range boardRectangles[row] {

			if tileHoverState.isHovered && tileHoverState.row == row && tileHoverState.col == col && mineBoard.TileState[row][col] == mines.Covered {
				// Draw the hover effect
				rl.DrawRectangleRec(tile, rl.Fade(rl.LightGray, 0.65))
				continue
			}
			// Set the rectangle position of the texture
			var rectPosition float32

			switch mineBoard.TileState[row][col] {
			case mines.Uncovered:
				if mineBoard.Board[row][col] != mines.Bomb {
					rectPosition = numsTextureSize * float32(mineBoard.Board[row][col])
				} else {
					rectPosition = numsTextureSize * 10
				}
			case mines.Covered:
				rectPosition = numsTextureSize * 9
			case mines.Flagged:
				rectPosition = numsTextureSize * 11
			}

			// Draw the tile texture
			rl.DrawTexturePro(
				numberTextures,
				rl.NewRectangle(rectPosition, 0, numsTextureSize, numsTextureSize),
				boardRectangles[row][col],
				rl.Vector2{X: 0, Y: 0},
				0, rl.White,
			)
		}
	}

	// Draw the current flags and the icon
	rl.DrawTextEx(shared.Font, flagsText, rl.Vector2{X: flagsTextXPos, Y: 25}, shared.FontMediumTextSize, 0, rg.TextColor())
	rl.DrawTexture(bombIconTexture, flagsIconXPos, 25, rl.White)

	// Draw the current time playing amount and the clock icon
	rl.DrawTextEx(shared.Font, clockText, rl.Vector2{X: clockTextXPos, Y: 25}, shared.FontMediumTextSize, 0, rg.TextColor())
	rl.DrawTexture(clockIconTexture, clockIconXPos, 25, rl.White)

	if gameLost {
		drawGameLostScreen()
	}
}

// Gameplay screen unload logic
func UnloadGameplayScreen() {
	// Unload the appropriate textures
	rl.UnloadTexture(numberTextures)
	rl.UnloadTexture(bombIconTexture)
	rl.UnloadTexture(clockIconTexture)
}

// Gameplay screen should finish
func FinishGameplayScreen() int {
	return finishScreen
}

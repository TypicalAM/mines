package gameplay

import (
	"fmt"
	"time"

	"github.com/TypicalAM/mines/src/mines"
	shared "github.com/TypicalAM/mines/src/screens"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Define the states that the game can be in
const (
	Playing int = iota
	Winning
	Losing
)

// Define exported variables
var ScreenState int // Determines the screen state
var GameState int   // Determines the game state (playing, winning, losing)

// Define local variables
var framesCounter int32 // Frames counter

const numsTextureSize float32 = 56 // The texture width and height (56x56)
var numberTextures rl.Texture2D    // The textures of the numbres (1-8, flag, empty and uncovered)
var bombIconTexture rl.Texture2D   // The texture of the bomb icon above the board
var clockIconTexture rl.Texture2D  // The texture of the clock icon above the board

var width int
var height int
var boardRectangles [][]rl.Rectangle // Rectangles making the playing board on screen
var mineBoard mines.MineBoard        // The mines playing board

var bgAlpha = 0.0
var textAlpha = 0.0
var bgAnimation = false
var textAnimation = false

type hover struct {
	isHovered bool
	row       int
	col       int
}

var tileHoverState hover  // If any tile was hovered, and which was hoverd
var timePlaying time.Time // Time of the first meaningful mouse press
var isPlaying bool        // If the player is in game

// Flags and bombs text placements
var flagsText string
var flagsTextXPos float32
var flagsIconXPos int32
var clockText string
var clockTextXPos float32
var clockIconXPos int32

// Keyboard & gamepad navtiation variables
var keyboardMode bool
var selectedButton int
var buttonPressed int
var cursorPosCol int
var cursorPosRow int

// Gameplay screen initialization logic
func Init() {
	// Init basic variables
	framesCounter = 0
	ScreenState = shared.Unchanged
	GameState = Playing

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

	// Adjust the cell size so that cells don't get stretched
	cellRatio := float32(cellWidth) / float32(cellHeight)
	for cellRatio < 0.9 || cellRatio > 1.1 {
		if cellHeight > cellWidth {
			cellHeight -= 4
		} else {
			cellWidth -= 4
		}
		cellRatio = float32(cellWidth) / float32(cellHeight)
	}

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

	isPlaying = false

	// Keyboard & gamepad variables
	keyboardMode = false
	selectedButton = 0
	buttonPressed = 0
	cursorPosCol = 0
	cursorPosRow = 0
}

// Gameplay screen update logic
func Update() {
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
					uncoverTile(row, col)
				}

				// If we detect the right click then flag/unflag the tile
				if rl.IsMouseButtonPressed(rl.MouseRightButton) {
					flagTile(row, col)
				}
			}
		}
	}

	if mineBoard.CheckIfWon() && GameState == Playing {
		InitWinning()
	}

	_, buttonPressed = shared.UpdateMovement(0, 0)
	if buttonPressed != shared.ButtonUnchanged {
		keyboardMode = true
	}

	switch buttonPressed {
	case shared.ButtonUp:
		cursorPosRow--
		if cursorPosRow == -1 {
			cursorPosRow = mineBoard.Height
		}
	case shared.ButtonDown:
		cursorPosRow++
		if cursorPosRow == mineBoard.Height {
			cursorPosRow = 0
		}
	case shared.ButtonLeft:
		cursorPosCol--
		if cursorPosCol == -1 {
			cursorPosCol = mineBoard.Width
		}
	case shared.ButtonRight:
		cursorPosCol++
		if cursorPosCol == mineBoard.Width {
			cursorPosCol = 0
		}
	case shared.ButtonConfirm:
		if keyboardMode && mineBoard.TileState[cursorPosRow][cursorPosCol] != mines.Flagged {
			uncoverTile(cursorPosRow, cursorPosCol)
		}
	case shared.ButtonFlag:
		if keyboardMode {
			flagTile(cursorPosRow, cursorPosCol)
		}
	case shared.ButtonRestart:
		ScreenState = shared.Gameplay
	case shared.ButtonGoBack:
		ScreenState = shared.Title
	}
}

// Gameplay screen draw logic
func Draw() {
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
				tile,
				rl.Vector2{X: 0, Y: 0},
				0, rl.White,
			)

			// Draw the cursor if on this tile
			if keyboardMode && row == cursorPosRow && col == cursorPosCol {
				rl.DrawRectangleLinesEx(tile, 4, rg.TextColor())
			}
		}
	}

	// Draw the current flags and the icon
	rl.DrawTextEx(shared.Font, flagsText, rl.Vector2{X: flagsTextXPos, Y: 25}, shared.FontMediumTextSize, 0, rg.TextColor())
	rl.DrawTexture(bombIconTexture, flagsIconXPos, 25, rl.White)

	// Draw the current time playing amount and the clock icon
	rl.DrawTextEx(shared.Font, clockText, rl.Vector2{X: clockTextXPos, Y: 25}, shared.FontMediumTextSize, 0, rg.TextColor())
	rl.DrawTexture(clockIconTexture, clockIconXPos, 25, rl.White)
}

// Gameplay screen unload logic
func Unload() {
	// Unload the appropriate textures
	rl.UnloadTexture(numberTextures)
	rl.UnloadTexture(bombIconTexture)
	rl.UnloadTexture(clockIconTexture)

	// Unload the winning or losing screens
	UnloadWinning()
	UnloadLosing()
}

// Uncover the hovered tile
func uncoverTile(row int, col int) {
	// Start the game timer
	if !isPlaying {
		isPlaying = true
		timePlaying = time.Now()
		// Check if we are on a bomb, if yes move it
		mineBoard.CheckAndMove(row, col)
	}

	// Uncover the values
	mineBoard.TileState[row][col] = mines.Uncovered
	if mineBoard.UncoverValues(true, row, col) {
		bombTile = boardRectangles[row][col]
		InitLosing()
	}
}

// Flag the tile
func flagTile(row int, col int) {
	// If the tile is flagged, unflag it and if it is covered, flag it
	switch mineBoard.TileState[row][col] {
	case mines.Flagged:
		mineBoard.TileState[row][col] = mines.Covered
		mineBoard.Flags--

	case mines.Covered:
		mineBoard.TileState[row][col] = mines.Flagged
		mineBoard.Flags++
	}

	// Start the game timer
	if !isPlaying {
		isPlaying = true
		timePlaying = time.Now()
	}
}

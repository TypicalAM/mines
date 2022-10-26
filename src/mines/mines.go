package mines

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
)

const Bomb int = -1
const Covered int = 0
const Uncovered int = 1
const Flagged int = 2

// A mine field
type MineBoard struct {
	Board     [][]int
	TileState [][]int
	Flags     int
	Mines     int
	Width     int
	Height    int
}

// Generate the mines board
func GenerateBoard(width int, height int, bombsPercent int) (MineBoard, error) {
	// Check if there are more s than we can put
	bombs := int(float32(bombsPercent*width*height)/100)
	if bombs >= width*height {
		return MineBoard{}, errors.New("there are more bombs than fields on the map")
	}

	// Create the game and tilestate board
	board := make([][]int, height)
	tileState := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
		tileState[i] = make([]int, width)
	}

	// Fill the game board with s
	sLeft := bombs
	for sLeft != 0 {
		posX := rand.Intn(height)
		posY := rand.Intn(width)
		for board[posX][posY] != Bomb {
			board[posX][posY]--
			sLeft--
		}
	}

	// Fill the numbers
	for row := range board {
		for col := range board[row] {
			// Skip the cell if it is a bomb
			if board[row][col] == Bomb {
				continue
			}

			// Get the nubmer of mines in the neighbouring 8 tiles
			var nearMines int
			positions := []struct {
				X int
				Y int
			}{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}}
			for _, pos := range positions {
				if col+pos.X < len(board[row]) && col+pos.X >= 0 && row+pos.Y < len(board) && row+pos.Y >= 0 && board[row+pos.Y][col+pos.X] == -1 {
					nearMines++
				}
			}

			// Set the tile as the number of neighbouring bombs
			board[row][col] = nearMines
		}
	}

	// Create the MineBoard object and return it
	return MineBoard{
		Board:     board,
		TileState: tileState,
		Mines:     bombs,
		Width:     width,
		Height:    height,
	}, nil
}

// Write the mineboard to a file
func WriteMineBoard(filepath string, board *MineBoard) error {

	// Try to marshall the board data
	jsonData, err := json.MarshalIndent(*board, " ", "\t")
	if err != nil {
		return errors.New("couldn't convert the board into json")
	}

	// Try to open the file for writing
	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("couldn't open the file for writing")
	}

	// Try to write the data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		return errors.New("couldn't write the json data to the file")
	}

	// All good
	return nil
}

// Check the clicked value, if the player has lost, bool is true
func (board *MineBoard) UncoverValues(firstRun bool, x int, y int) (isLost bool) {
	boardValue := *board
	value := boardValue.Board[x][y]

	// If we have a bomb, we lost the game
	if value == Bomb {
		return true
	}

	if value != 0 && !firstRun {
		return false
	}

	// Run recursively for every possible non-bomb neighbour
	neighbours := []struct {
		X int
		Y int
	}{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}}
	for _, pos := range neighbours {
		// Check if the position is valid
		if x+pos.X < len(board.Board) && x+pos.X >= 0 && y+pos.Y < len(board.Board[0]) && y+pos.Y >= 0 {

			// Check if the tile is covered, and its position is not 0
			if boardValue.TileState[x+pos.X][y+pos.Y] == Covered && board.Board[x+pos.X][y+pos.Y] != -1 {

				// Set the tile state to uncovered
				boardValue.TileState[x+pos.X][y+pos.Y] = Uncovered

				// Run recursively for every neighbour
				boardValue.UncoverValues(false, x+pos.X, y+pos.Y)
			}
		}
	}

	// We don't have a bomb, return the value
	return false
}

// Check if the game is won
func (board *MineBoard) CheckIfWon() bool {
	var totalValid int
	var coveredTiles int
	boardValue := *board

	// Check if every mine tile is covered or flagged
	for row := range boardValue.Board {
		for col := range boardValue.Board[row] {
			if boardValue.Board[row][col] == Bomb && boardValue.TileState[row][col] == Flagged {
				totalValid++
			} else if boardValue.TileState[row][col] == Covered {
				coveredTiles++
			}
		}
	}

	// If every mine is flagged and we have no uncovered tiles, we win
	if totalValid == board.Mines && coveredTiles == 0 {
		return true
	} else {
		return false
	}
}

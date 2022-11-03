package gameplay

import (
	"example/raylib-game/src/gui"
	shared "example/raylib-game/src/screens"
	"fmt"
	"strconv"
	"strings"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables
var bgAlpha = 0.0
var textAlpha = 0.0
var bgAnimation = false
var textAnimation = false
var newRecord bool = false
var scoreboardPlace int = 0

// Initialize the game losing screen
func InitWinning() {
	fmt.Println("We have won the game")
	GameState = Winning
	bgAnimation = true

	textAnimation = false
	isPlaying = false

	textAlpha = 0.0
	bgAlpha = 0.0

	timeSplit := strings.Split(clockText, ":")
	minutes, _ := strconv.Atoi(timeSplit[0])
	seconds, _ := strconv.Atoi(timeSplit[1])

	newRecord, scoreboardPlace = shared.Scores.CanItBeInTheScoreboard(minutes*60 + seconds)
}

// Update the game winning screen
func UpdateWinning() {
	// Fade in the background
	if bgAnimation {
		bgAlpha += 0.01
		if bgAlpha >= 0.5 {
			bgAnimation = false
			textAnimation = true
		}
	}

	// Fade in the text
	if textAnimation {
		textAlpha += 0.03
		if textAlpha >= 1.0 {
			textAnimation = false
		}
	}
}

// Draw the game winning screen
func DrawWinning() {
	// The fade in background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Fade(
		rg.BackgroundColor(),
		float32(bgAlpha),
	))

	measure := rl.MeasureTextEx(shared.Font, "You have won the game!", shared.FontHugeTextSize*2, 0)

	// The fade in text
	rl.DrawTextEx(shared.Font, "You have won the game", rl.Vector2{
		X: float32(rl.GetScreenWidth())/2 - measure.X/2,
		Y: float32(rl.GetScreenHeight())/2 - measure.Y/2 - 330,
	}, shared.FontHugeTextSize*2, 0, rl.Fade(rl.Red, float32(textAlpha)))

	if textAnimation || bgAnimation || !newRecord {
		return
	}

	rectangleWidths := float32(rl.GetScreenWidth()) / 3
	rectangleXPos := (float32(rl.GetScreenWidth()) - rectangleWidths) / 2

	textboxRect := rl.NewRectangle(
		rectangleXPos+rectangleWidths/4,
		float32(rl.GetScreenHeight()/2+250),
		rectangleWidths/2, 60,
	)

	baseRectY := -250
	baseOffsetY := 100

	textboxesBounds := make([]rl.Rectangle, 5)
	for i := range textboxesBounds {
		textboxesBounds[i] = rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+i*baseOffsetY), rectangleWidths, 80)
	}

	switch scoreboardPlace {
	case 0:
	// Our score on top, and then 4 below
	case 1:
	// First score, then ours, then 3 below
	case len(shared.Scores.Entries) - 1:
	// Three scores on top, then ours, then 1 below
	case len(shared.Scores.Entries):
	// Our score is the last score
	}
	gui.TextBoxEx(shared.Font, textboxRect, "text", shared.FontBigTextSize, 10)
}

// Unload the winning files
func UnloadWinning() {}

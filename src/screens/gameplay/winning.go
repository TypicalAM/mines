package gameplay

import (
	"example/raylib-game/src/gui"
	shared "example/raylib-game/src/screens"
	"example/raylib-game/src/settings"
	"fmt"
	"strconv"
	"strings"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables
var newRecord bool = false
var scoreboardPlace int = 0
var scoresRect [5]rl.Rectangle
var saveRect rl.Rectangle
var scoreSaved bool
var newScoreName string
var displayedScores []string
var gameTime int
var scoreboardEntries []settings.Entry

// Initialize the game losing screen
func InitWinning() {
	GameState = Winning
	isPlaying = false
	bgAnimation = true
	textAnimation = false
	textAlpha = 0.0
	bgAlpha = 0.0
	scoreSaved = false

	timeSplit := strings.Split(clockText, ":")
	minutes, _ := strconv.Atoi(timeSplit[0])
	seconds, _ := strconv.Atoi(timeSplit[1])
	gameTime = minutes*60 + seconds

	newRecord, scoreboardPlace = shared.Scores.CanItBeInTheScoreboard(shared.AppSettings, gameTime)
	newScoreName = ""

	// Filter the scores by the settings
	var filter int

	if shared.AppSettings.Width == 8 && shared.AppSettings.Height == 8 && shared.AppSettings.Bombs == 15 {
		filter = settings.Beginner
	} else if shared.AppSettings.Width == 16 && shared.AppSettings.Height == 16 && shared.AppSettings.Bombs == 15 {
		filter = settings.Intermediate
	} else if shared.AppSettings.Width == 30 && shared.AppSettings.Height == 16 && shared.AppSettings.Bombs == 21 {
		filter = settings.Expert
	} else {
		filter = settings.Custom
	}

	scoreboardEntries = shared.Scores.FilterScores(filter)

	rectangleWidths := float32(rl.GetScreenWidth()) / 3
	rectangleXPos := (float32(rl.GetScreenWidth()) - rectangleWidths) / 2

	saveRect = rl.NewRectangle(
		rectangleXPos+rectangleWidths/4,
		float32(rl.GetScreenHeight()/2+250),
		rectangleWidths/2, 60,
	)

	baseRectY := -250
	baseOffsetY := 100

	for i := range scoresRect {
		scoresRect[i] = rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+i*baseOffsetY), rectangleWidths, 80)
	}

	switch scoreboardPlace {
	case 0:
		displayedScores = []string{
			"mine",
			fmt.Sprint(scoreboardEntries[0].Time),
			fmt.Sprint(scoreboardEntries[1].Time),
			fmt.Sprint(scoreboardEntries[2].Time),
			fmt.Sprint(scoreboardEntries[3].Time),
		}
	case 1:
		displayedScores = []string{
			fmt.Sprint(scoreboardEntries[0].Time),
			"mine",
			fmt.Sprint(scoreboardEntries[1].Time),
			fmt.Sprint(scoreboardEntries[2].Time),
			fmt.Sprint(scoreboardEntries[3].Time),
		}
	case len(scoreboardEntries) - 2:
		displayedScores = []string{
			fmt.Sprint(scoreboardEntries[len(scoreboardEntries)-4].Time),
			fmt.Sprint(scoreboardEntries[len(scoreboardEntries)-3].Time),
			fmt.Sprint(scoreboardEntries[len(scoreboardEntries)-2].Time),
			"mine",
			fmt.Sprint(scoreboardEntries[len(scoreboardEntries)-1].Time),
		}
	case len(scoreboardEntries) - 1:
		displayedScores = []string{
			fmt.Sprint(scoreboardEntries[len(scoreboardEntries)-4]),
			fmt.Sprint(scoreboardEntries[len(scoreboardEntries)-3]),
			fmt.Sprint(scoreboardEntries[len(scoreboardEntries)-2]),
			fmt.Sprint(scoreboardEntries[len(scoreboardEntries)-1]),
			"mine",
		}
	default:
		displayedScores = []string{
			fmt.Sprint(scoreboardEntries[scoreboardPlace-2].Time),
			fmt.Sprint(scoreboardEntries[scoreboardPlace-1].Time),
			"mine",
			fmt.Sprint(scoreboardEntries[scoreboardPlace].Time),
			fmt.Sprint(scoreboardEntries[scoreboardPlace+1].Time),
		}
	}
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

	if scoreSaved {
		newScoreName = strings.Join(strings.Fields(newScoreName), " ")
		if err := shared.Scores.InsertNewScore(shared.AppSettings, newScoreName, gameTime, scoreboardPlace); err != nil {
			rl.TraceLog(rl.LogFatal, "Couldn't save the new score")
		} else {
			ScreenState = shared.Title
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
	}, shared.FontHugeTextSize*2, 0, rl.Fade(rg.TextColor(), float32(textAlpha)))

	if textAnimation || bgAnimation || !newRecord {
		return
	}

	for pos, rect := range scoresRect {
		if displayedScores[pos] == "mine" {
			newScoreName = gui.TextBoxEx(shared.Font, rect, newScoreName, shared.FontBigTextSize, 20)
		} else {
			gui.ButtonEx(shared.Font, rect, displayedScores[pos], shared.FontBigTextSize)
		}
	}

	scoreSaved = gui.ButtonEx(shared.Font, saveRect, "SAVE", shared.FontBigTextSize)
}

// Unload the winning files
func UnloadWinning() {}

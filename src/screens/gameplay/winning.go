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
var scoresRect []rl.Rectangle
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

	newScoreName = "Anonymous"
	newRecord, scoreboardPlace = shared.Scores.CanItBeInTheScoreboard(filter, gameTime)
	shared.Scores.InsertNewScore(shared.AppSettings, newScoreName, gameTime)

	scoreboardEntries = shared.Scores.FilterScores(filter)
	if len(scoreboardEntries) >= 5 {
		displayedScores = make([]string, 5)

		slice := scoreboardEntries
		switch scoreboardPlace {
		case 0, 1:
			slice = scoreboardEntries[:5]
		case len(scoreboardEntries) - 2, len(scoreboardEntries) - 1:
			slice = scoreboardEntries[5:]
		default:
			slice = scoreboardEntries[scoreboardPlace-2 : scoreboardPlace+3]
		}

		// Iterate over the slice and add the results to the displayed scores
		for pos, entry := range slice {
			displayedScores[pos] = fmt.Sprintf("%s - %d", entry.Name, entry.Time)
		}
	} else {
		// Crewate the displayedscores array and add the only results that we have
		displayedScores = make([]string, len(scoreboardEntries))

		for pos, entry := range scoreboardEntries {
			displayedScores[pos] = fmt.Sprintf("%s - %d", entry.Name, entry.Time)
		}
	}
	fmt.Println(scoreboardEntries)

	newScoreName = ""

	rectangleWidths := float32(rl.GetScreenWidth()) / 3
	rectangleXPos := (float32(rl.GetScreenWidth()) - rectangleWidths) / 2

	saveRect = rl.NewRectangle(
		rectangleXPos+rectangleWidths/4,
		float32(rl.GetScreenHeight()/2+250),
		rectangleWidths/2, 60,
	)

	baseRectY := -250
	baseOffsetY := 100

	// Make the scores rectangles
	scoresRect = make([]rl.Rectangle, len(displayedScores))
	for i := range scoresRect {
		scoresRect[i] = rl.NewRectangle(rectangleXPos, float32(rl.GetScreenHeight()/2+baseRectY+i*baseOffsetY), rectangleWidths, 80)
	}

	rl.SetExitKey(rl.KeyEscape)
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
		// If there is an anonymous score, change its name to the new one
		for pos, entry := range shared.Scores.Entries {
			if entry.Name == "Anonymous" {
				shared.Scores.Entries[pos].Name = strings.Join(strings.Fields(newScoreName), " ")
			}
		}

		// Change to the title screen (the unload function will save the results) 
		ScreenState = shared.Title
	}
}

// Draw the game winning screen
func DrawWinning() {
	// The fade in background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Fade(
		rg.BackgroundColor(),
		float32(bgAlpha),
	))

	// The fade in text
	measure := rl.MeasureTextEx(shared.Font, "You have won the game!", shared.FontHugeTextSize, 0)
	rl.DrawTextEx(shared.Font, "You have won the game", rl.Vector2{
		X: float32(rl.GetScreenWidth())/2 - measure.X/2,
		Y: float32(rl.GetScreenHeight())/2 - measure.Y/2 - 330,
	}, shared.FontHugeTextSize, 0, rl.Fade(rg.TextColor(), float32(textAlpha)))

	if textAnimation || bgAnimation || !newRecord {
		return
	}

	for pos, rect := range scoresRect {
		if strings.Split(displayedScores[pos], " - ")[0] == "Anonymous" {
			newScoreName = gui.TextBoxEx(shared.Font, rect, newScoreName, shared.FontBigTextSize, 20)
		} else {
			gui.ButtonEx(shared.Font, rect, displayedScores[pos], shared.FontBigTextSize)
		}
	}

	scoreSaved = gui.ButtonEx(shared.Font, saveRect, "SAVE", shared.FontBigTextSize)
}

// Unload the winning files
func UnloadWinning() {
	// If there is an anonymous score, delete it from the scores slice
	for pos, entry := range shared.Scores.Entries {
		if entry.Name == "Anonymous" {
			shared.Scores.Entries = append(shared.Scores.Entries[:pos], shared.Scores.Entries[pos+1:]...)
		}
	}

	// Write the scores
	if err := shared.Scores.WriteToFile(); err != nil {
		rl.TraceLog(rl.LogFatal, "Couldn't save the new score")
	}

	rl.SetExitKey(rl.KeyQ)
}

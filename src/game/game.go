package game

import (
	shared "example/raylib-game/src/screens"
	"example/raylib-game/src/screens/ending"
	"example/raylib-game/src/screens/gameplay"
	"example/raylib-game/src/screens/leaderboard"
	"example/raylib-game/src/screens/logo"
	"example/raylib-game/src/screens/options"
	"example/raylib-game/src/screens/title"
	"fmt"
	"math/rand"
	"time"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables
const screenWidth int32 = 1600
const screenHeight int32 = 900

var transAlpha float32
var onTransition bool
var transFadeOut bool

var currentScreen int = shared.Unchanged
var transToScreen int = shared.Unchanged

// Main entry point to the game
func InitalizeGame() {
	// Init the window
	rl.InitWindow(screenWidth, screenHeight, "Go mines")

	// Set the exit key to q
	rl.SetExitKey(rl.KeyQ)

	// The random int sequences are deterministic, if we
	// don't set a seed the mineboard will be the same
	// every time
	rand.Seed(time.Now().UnixNano())

	// Load the shared assets
	if err := shared.LoadSharedAssets(); err != nil {
		rl.TraceLog(rl.LogFatal, "Failed to load the share assets: ", err)
	}

	// Load the app-wide gui style
	rg.LoadGuiStyle(fmt.Sprintf("resources/styles/%s.style", shared.AppSettings.Theme))

	// Setup first screen
	currentScreen = shared.Title
	title.Init()

	rl.SetTargetFPS(60)

	// Update the frames
	for !rl.WindowShouldClose() {
		UpdateDrawFrame()
	}

	// Unload screen
	switch currentScreen {
	case shared.Logo:
		logo.Unload()
	case shared.Title:
		title.Unload()
	case shared.Options:
		options.Unload()
	case shared.Gameplay:
		gameplay.Unload()
	case shared.Ending:
		ending.Unload()
	case shared.Leaderboard:
		leaderboard.Unload()
	}

	// Clean up after ourselves
	rl.CloseWindow()
}

// Change to the screen
func ChangeToScreen(screen int) {
	// Unload current screen
	switch currentScreen {
	case shared.Logo:
		logo.Unload()
	case shared.Title:
		title.Unload()
	case shared.Options:
		options.Unload()
	case shared.Gameplay:
		gameplay.Unload()
	case shared.Ending:
		ending.Unload()
	case shared.Leaderboard:
		leaderboard.Unload()
	}

	// Init next screen
	switch screen {
	case shared.Logo:
		logo.Init()
	case shared.Title:
		title.Init()
	case shared.Options:
		options.Init()
	case shared.Gameplay:
		gameplay.Init()
	case shared.Ending:
		ending.Init()
	case shared.Leaderboard:
		leaderboard.Init()
	}
	currentScreen = screen
}

// Request transition to the next screen
func Transition(screen int) {
	onTransition = true
	transFadeOut = false
	transToScreen = screen
	transAlpha = 0.0
}

// Udpate the transition effect
func UpdateTransition() {
	if !transFadeOut {
		transAlpha += 0.05

		if transAlpha > 1.01 {
			transAlpha = 1.0

			ChangeToScreen(transToScreen)
			transFadeOut = true
		}
	} else {
		transAlpha -= 0.02

		if transAlpha < -0.01 {
			transAlpha = 0.0
			transFadeOut = false
			onTransition = false
			transToScreen = shared.Unchanged
		}
	}
}

// Draw the transition effect
func DrawTransition() {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Fade(rl.Black, transAlpha))
}

// Update and draw one frame
func UpdateDrawFrame() {

	if !onTransition {
		switch currentScreen {
		case shared.Logo:
			logo.Update()

			if logo.ScreenState == shared.Title {
				Transition(shared.Title)
			}
		case shared.Title:
			title.Update()

			switch title.ScreenState {
			case shared.Gameplay:
				Transition(shared.Gameplay)
			case shared.Leaderboard:
				Transition(shared.Leaderboard)
			case shared.Options:
				Transition(shared.Options)
			}
		case shared.Options:
			options.Update()

			if options.ScreenState == shared.Title {
				Transition(shared.Title)
			}
		case shared.Gameplay:
			if gameplay.GameState == gameplay.Playing {
				gameplay.Update()
			}

			switch gameplay.GameState {
			case gameplay.Winning:
				gameplay.UpdateWinning()
			case gameplay.Losing:
				gameplay.UpdateLosing()
			}

			switch gameplay.ScreenState {
			case shared.Ending:
				Transition(shared.Ending)
			case shared.Gameplay:
				Transition(shared.Gameplay)
			case shared.Title:
				Transition(shared.Title)
			}
		case shared.Ending:
			ending.Update()

			if ending.ScreenState == shared.Title {
				Transition(shared.Title)
			}
		case shared.Leaderboard:
			leaderboard.Update()

			if leaderboard.ScreenState == shared.Title {
				Transition(shared.Title)
			}
		}
	} else {
		UpdateTransition() // Update transition (fade-in, fade-out)
	}

	rl.BeginDrawing()

	switch currentScreen {
	case shared.Logo:
		logo.Draw()
	case shared.Title:
		title.Draw()
	case shared.Options:
		options.Draw()
	case shared.Gameplay:
		gameplay.Draw()

		switch gameplay.GameState {
		case gameplay.Winning:
			gameplay.DrawWinning()
		case gameplay.Losing:
			gameplay.DrawLosing()
		}
	case shared.Ending:
		ending.Draw()
	case shared.Leaderboard:
		leaderboard.Draw()
	}

	if onTransition {
		DrawTransition()
	}

	rl.EndDrawing()
}

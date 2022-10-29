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

var currentScreen int = shared.Unchanged
var transAlpha float32
var onTransition bool
var transFadeOut bool
var transFromScreen int = shared.Unchanged
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
	shared.LoadSharedAssets()

	// Load the app-wide gui style
	rg.LoadGuiStyle(fmt.Sprintf("resources/styles/%s.style", shared.AppSettings.Theme))

	// Setup first screen
	currentScreen = shared.Title
	title.InitTitleScreen()

	rl.SetTargetFPS(60)

	// Update the frames
	for !rl.WindowShouldClose() {
		UpdateDrawFrame()
	}

	// Unload screen
	switch currentScreen {
	case shared.Logo:
		logo.UnloadLogoScreen()
	case shared.Title:
		title.UnloadTitleScreen()
	case shared.Options:
		options.UnloadOptionsScreen()
	case shared.Gameplay:
		gameplay.UnloadGameplayScreen()
	case shared.Ending:
		ending.UnloadEndingScreen()
	case shared.Leaderboard:
		leaderboard.UnloadLeaderboardScreen()
	}

	// Clean up after ourselves
	// rl.UnloadSound(FxClick)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

// Change to the screen
func ChangeToScreen(screen int) {
	// Unload current screen
	switch currentScreen {
	case shared.Logo:
		logo.UnloadLogoScreen()
	case shared.Title:
		title.UnloadTitleScreen()
	case shared.Options:
		options.UnloadOptionsScreen()
	case shared.Gameplay:
		gameplay.UnloadGameplayScreen()
	case shared.Ending:
		ending.UnloadEndingScreen()
	case shared.Leaderboard:
		leaderboard.UnloadLeaderboardScreen()
	}

	// Init next screen
	switch screen {
	case shared.Logo:
		logo.InitLogoScreen()
	case shared.Title:
		title.InitTitleScreen()
	case shared.Options:
		options.InitOptionsScreen()
	case shared.Gameplay:
		gameplay.InitGameplayScreen()
	case shared.Ending:
		ending.InitEndingScreen()
	case shared.Leaderboard:
		leaderboard.InitLeaderboardScreen()
	}

	currentScreen = screen
}

// Request transition to the next screen
func TransitionToScreen(screen int) {
	onTransition = true
	transFadeOut = false
	transFromScreen = currentScreen
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
			transFromScreen = shared.Unchanged
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
			logo.UpdateLogoScreen()

			if logo.FinishLogoScreen() == shared.Title {
				TransitionToScreen(shared.Title)
			}
		case shared.Title:
			title.UpdateTitleScreen()

			switch title.FinishTitleScreen() {
			case shared.Gameplay:
				TransitionToScreen(shared.Gameplay)
			case shared.Leaderboard:
				TransitionToScreen(shared.Leaderboard)
			case shared.Options:
				TransitionToScreen(shared.Options)
			}
		case shared.Options:
			options.UpdateOptionsScreen()

			if options.FinishOptionsScreen() == shared.Title {
				TransitionToScreen(shared.Title)
			}
		case shared.Gameplay:
			gameplay.UpdateGameplayScreen()

			switch gameplay.FinishGameplayScreen() {
			case shared.Ending:
				TransitionToScreen(shared.Ending)
			case shared.Gameplay:
				TransitionToScreen(shared.Gameplay)
			case shared.Title:
				TransitionToScreen(shared.Title)
			}
		case shared.Ending:
			ending.UpdateEndingScreen()
			if ending.FinishEndingScreen() == shared.Title {
				TransitionToScreen(shared.Title)
			}
		case shared.Leaderboard:
			leaderboard.UpdateLeaderboardScreen()
			if leaderboard.FinishLeaderboardScreen() == shared.Title {
				TransitionToScreen(shared.Title)
			}
		}
	} else {
		UpdateTransition() // Update transition (fade-in, fade-out)
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	switch currentScreen {
	case shared.Logo:
		logo.DrawLogoScreen()
	case shared.Title:
		title.DrawTitleScreen()
	case shared.Options:
		options.DrawOptionsScreen()
	case shared.Gameplay:
		gameplay.DrawGameplayScreen()
	case shared.Ending:
		ending.DrawEndingScreen()
	case shared.Leaderboard:
		leaderboard.DrawLeaderboardScreen()
	}

	if onTransition {
		DrawTransition()
	}

	rl.EndDrawing()
}

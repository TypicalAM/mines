package game

import (
	"example/raylib-game/src/screens"
	"example/raylib-game/src/screens/ending"
	"example/raylib-game/src/screens/gameplay"
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
const screenWidth int32 = 1280
const screenHeight int32 = 720

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

	// The random int sequences are deterministic, if we
	// don't set a seed the mineboard will be the same
	// every time
	rand.Seed(time.Now().UnixNano())

	// Load the shared assets
	shared.LoadSharedAssets()

	// Load the app-wide gui style
	rg.LoadGuiStyle(fmt.Sprintf("resources/styles/%s.style", shared.AppSettings.Theme))

	// Setup first screen
	currentScreen = shared.Options
	options.InitOptionsScreen()

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

			if title.FinishTitleScreen() == shared.Gameplay {
				fmt.Println("Transitioning to the gameplay")
				TransitionToScreen(shared.Gameplay)
			}
		case shared.Options:
			options.UpdateOptionsScreen()

			if options.FinishOptionsScreen() == shared.Title {
				TransitionToScreen(shared.Title)
			}
		case shared.Gameplay:
			gameplay.UpdateGameplayScreen()

			if gameplay.FinishGameplayScreen() == shared.Ending {
				TransitionToScreen(shared.Ending)
			} else if gameplay.FinishGameplayScreen() == shared.Gameplay {
				TransitionToScreen(shared.Gameplay)
			}
		case shared.Ending:
			ending.UpdateEndingScreen()
			if ending.FinishEndingScreen() == shared.Title {
				TransitionToScreen(shared.Title)
			}
		}
	} else {
		UpdateTransition() // Update transition (fade-in, fade-out)
	}
	//----------------------------------------------------------------------------------

	// Draw
	//----------------------------------------------------------------------------------
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
	}

	// Draw full screen rectangle in front of everything
	if onTransition {
		DrawTransition()
	}

	//DrawFPS(10, 10);

	rl.EndDrawing()
	//----------------------------------------------------------------------------------
}

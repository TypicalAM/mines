package logo

import (
	shared "example/raylib-game/src/screens"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Local variables

var framesCounter int32
var ScreenState int
var logoPositionX int32
var logoPositionY int32
var lettersCount int32
var topRecWidth int32
var leftRecHeight int32
var bottomRecWidth int32
var rightRecHeight int32
var state int32
var alpha float32 = 1.0

// Logo screen functions
func Init() {
	ScreenState = shared.Unchanged
	framesCounter = 0
	lettersCount = 0

	logoPositionX = int32(rl.GetScreenWidth())/2 - 128
	logoPositionY = int32(rl.GetScreenHeight())/2 - 128

	topRecWidth = 16
	leftRecHeight = 16
	bottomRecWidth = 16
	rightRecHeight = 16

	state = 0
	alpha = 1.0
}

// Logo screen update
func Update() {
	switch state {
	case 0:
		framesCounter++
		if framesCounter == 80 {
			state = 1
			framesCounter = 0
		}

	case 1:
		topRecWidth += 8
		leftRecHeight += 8
		if topRecWidth == 256 {
			state = 2
		}

	case 2:
		bottomRecWidth += 8
		rightRecHeight += 8
		if bottomRecWidth == 256 {
			state = 3
		}

	case 3:
		framesCounter++
		if lettersCount < 6 {
			if framesCounter/12 == 1 {
				lettersCount++
				framesCounter = 0
			}

			break
		}

		if framesCounter <= 200 {
			break
		}

		alpha -= 0.02
		if alpha <= 0.0 {
			alpha = 0.0
			ScreenState = shared.Title
		}
	}
}

// Logo Screen Draw Logic
func Draw() {
	switch state {
	case 0:
		if (framesCounter/10)%2 == 1 {
			rl.DrawRectangle(logoPositionX, logoPositionY, 16, 16, rl.Black)
		}

	case 1: // Draw bars animation: top and left
		rl.DrawRectangle(logoPositionX, logoPositionY, topRecWidth, 16, rl.Black)
		rl.DrawRectangle(logoPositionX, logoPositionY, 16, leftRecHeight, rl.Black)

	case 2: // Draw bars animation: bottom and right
		rl.DrawRectangle(logoPositionX, logoPositionY, topRecWidth, 16, rl.Black)
		rl.DrawRectangle(logoPositionX, logoPositionY, 16, leftRecHeight, rl.Black)

		rl.DrawRectangle(logoPositionX+240, logoPositionY, 16, rightRecHeight, rl.Black)
		rl.DrawRectangle(logoPositionX, logoPositionY+240, bottomRecWidth, 16, rl.Black)

	case 3: // Draw "raylib" text-write animation + "powered by"
		rl.DrawRectangle(logoPositionX, logoPositionY, topRecWidth, 16, rl.Fade(rl.Black, alpha))
		rl.DrawRectangle(logoPositionX, logoPositionY+16, 16, leftRecHeight-32, rl.Fade(rl.Black, alpha))

		rl.DrawRectangle(logoPositionX+240, logoPositionY+16, 16, rightRecHeight-32, rl.Fade(rl.Black, alpha))
		rl.DrawRectangle(logoPositionX, logoPositionY+240, bottomRecWidth, 16, rl.Fade(rl.Black, alpha))

		rl.DrawRectangle(int32(rl.GetScreenWidth())/2-112, int32(rl.GetScreenHeight())/2-112, 224, 224, rl.Fade(rl.RayWhite, alpha))

		rl.DrawText("raylib"[:lettersCount], int32(rl.GetScreenWidth())/2-44, int32(rl.GetScreenHeight())/2+48, 50, rl.Fade(rl.Black, alpha))

		if framesCounter > 20 {
			rl.DrawText("powered by", logoPositionX, logoPositionY-27, 20, rl.Fade(rl.DarkGray, alpha))
		}
	}
}

// Logo Screen Unload logic
func Unload() {}

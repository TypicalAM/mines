package gui

import (
	"time"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var backspaceHeld = false
var nextBackspace = time.Now()
var framesCounter = 0

// TextBoxEx - TextBox element, but with custom fonting and case sensitivity
func TextBoxEx(font rl.Font, bounds rl.Rectangle, text string, maxChars int) string {
	bound := bounds.ToInt32()
	letter := int32(-1)

	state := rg.GetInteractionState(bounds)
	borderColor := rg.TextboxBorderColor
	if state == rg.Pressed || state == rg.Focused {
		borderColor = rg.ToggleActiveBorderColor

		framesCounter++
		letter = rl.GetCharPressed()
		if letter != -1 {
			if letter >= 32 && letter < 127 {
				text += string(letter)
			}
		}

		backspacing := rl.IsKeyPressed(rl.KeyBackspace)
		if backspacing {
			nextBackspace = time.Now().Add(rg.BackspaceRepeatDelay)
		} else if rl.IsKeyDown(rl.KeyBackspace) {
			backspacing = time.Since(nextBackspace) >= 0
			if backspacing {
				nextBackspace = time.Now().Add(rg.BackspaceRepeatInterval)
			}
		}

		if backspacing && len(text) > 0 || len(text) > maxChars {
			text = text[:len(text)-1]
		}
	}

	rg.DrawBorderedRectangle(
		bound,
		rg.GetStyle32(rg.TextboxBorderWidth),
		rg.GetStyleColor(borderColor),
		rg.GetStyleColor(rg.TextboxInsideColor),
	)

	rl.DrawTextEx(
		font,
		text,
		rl.Vector2{
			X: float32(bound.X + 2),
			Y: float32(bound.Y + rg.GetStyle32(rg.TextboxBorderWidth) + bound.Height/2 - rg.GetStyle32(rg.TextboxTextFontsize)/2),
		},
		float32(rg.GetStyle32(rg.TextboxTextFontsize)),
		0,
		rg.GetStyleColor(rg.TextboxTextColor),
	)

	if state == rg.Focused || state == rg.Pressed {
		if (framesCounter/20)%2 == 0 {
			rl.DrawRectangle(
				bound.X+4+int32(rl.MeasureTextEx(font, text, float32(rg.GetStyle32(rg.TextboxTextFontsize)), 0).X),
				bound.Y+2,
				2,
				bound.Height-10,
				rg.GetStyleColor(rg.TextboxLineColor),
			)
		}
	}

	return text
}

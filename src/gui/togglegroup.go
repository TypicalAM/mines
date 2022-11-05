package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ToggleGroupEx - Toggle Group element, returns toggled button index
func ToggleGroupEx(font rl.Font, bounds rl.Rectangle, toggleText []string, active int, size float32) int {
	for i := range toggleText {
		if i == active {
			ToggleButtonEx(
				font,
				rl.NewRectangle(bounds.X+float32(i)*(bounds.Width), bounds.Y, bounds.Width, bounds.Height),
				toggleText[i],
				true,
				size,
			)
		} else {
			if ToggleButtonEx(
				font,
				rl.NewRectangle(bounds.X+float32(i)*(bounds.Width), bounds.Y, bounds.Width, bounds.Height),
				toggleText[i],
				false,
				size,
			) {
				active = i
			}
		}
	}
	return active
}

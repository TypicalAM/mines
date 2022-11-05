package gui

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type togglebuttonColoring struct {
	Border, Inside, Text rg.Property
}

var togglebuttonColors = map[rg.ControlState]togglebuttonColoring{
	rg.Normal: {rg.ToggleDefaultBorderColor, rg.ToggleDefaultInsideColor, rg.ToggleDefaultTextColor},
	// Hijacking 'Clicked' for the 'active' state.
	rg.Clicked: {rg.ToggleActiveBorderColor, rg.ToggleActiveInsideColor, rg.ToggleDefaultTextColor},
	rg.Pressed: {rg.TogglePressedBorderColor, rg.TogglePressedInsideColor, rg.TogglePressedTextColor},
	rg.Focused: {rg.ToggleHoverBorderColor, rg.ToggleHoverInsideColor, rg.ToggleHoverTextColor},
}

// ToggleButton - Toggle Button element, returns true when active
func ToggleButtonEx(font rl.Font, bounds rl.Rectangle, text string, active bool, size float32) bool {
	textSize := rl.MeasureTextEx(font, text, size, 0)

	state := rg.GetInteractionState(bounds)
	if state == rg.Clicked {
		active = !active
		state = rg.Normal
	}

	if state == rg.Normal && active {
		state = rg.Clicked
	}

	colors, exists := togglebuttonColors[state]
	if !exists {
		return active
	}

	b := bounds.ToInt32()
	rg.DrawBorderedRectangle(b, rg.GetStyle32(rg.ToggleBorderWidth), rg.GetStyleColor(colors.Border), rg.GetStyleColor(colors.Inside))

	rl.DrawTextEx(font, text,
		rl.Vector2{
			X: bounds.X + bounds.Width/2 - textSize.X/2,
			Y: bounds.Y + bounds.Height/2 - textSize.Y/2,
		},
		size, 0, rl.White)

	return active
}

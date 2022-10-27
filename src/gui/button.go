package gui

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// buttonColoring describes the per-state properties for a Button control.
type buttonColoring struct {
	Border, Inside, Text rg.Property
}

// buttonColors lists the styling for each supported state.
var buttonColors = map[rg.ControlState]buttonColoring{
	rg.Normal:  {rg.ButtonDefaultBorderColor, rg.ButtonDefaultInsideColor, rg.ButtonDefaultTextColor},
	rg.Clicked: {rg.ButtonDefaultBorderColor, rg.ButtonDefaultInsideColor, rg.ButtonDefaultTextColor},
	rg.Focused: {rg.ButtonHoverBorderColor, rg.ButtonHoverInsideColor, rg.ButtonHoverTextColor},
	rg.Pressed: {rg.ButtonPressedBorderColor, rg.ButtonPressedInsideColor, rg.ButtonPressedTextColor},
}

// ButtonEx - Button element, but with support for custom fonts
func ButtonEx(font rl.Font, bound rl.Rectangle, text string, size float32) bool {
	state := rg.GetInteractionState(bound)
	colors, exist := buttonColors[state]
	if !exist {
		return false
	}

	textSize := rl.MeasureTextEx(font, text, size, 0)
	rg.DrawBorderedRectangle(bound.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(colors.Border), rg.GetStyleColor(colors.Inside))
	rl.DrawTextEx(font, text, rl.Vector2{X: bound.X + bound.Width/2 - textSize.X/2, Y: bound.Y + bound.Height/2 - textSize.Y/2}, size, 0, rg.GetStyleColor(colors.Text))

	return state == rg.Clicked
}

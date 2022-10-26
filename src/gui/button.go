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
func ButtonEx(font rl.Font, bounds rl.Rectangle, text string, size float32) bool {
	textSize := rl.MeasureTextEx(font, text, size, 0)
	textHeight := int32(textSize.X)
	textWidth := int32(textSize.Y)

	rg.ConstrainRectangle(
		&bounds,
		textWidth,
		textWidth+rg.GetStyle32(rg.ButtonTextPadding),
		textHeight,
		textHeight+rg.GetStyle32(rg.ButtonTextPadding)/2,
	)

	state := rg.GetInteractionState(bounds)
	colors, exist := buttonColors[state]
	if !exist {
		return false
	}

	bound := bounds.ToInt32()
	rg.DrawBorderedRectangle(bound,
		rg.GetStyle32(rg.ButtonBorderWidth),
		rg.GetStyleColor(colors.Border),
		rg.GetStyleColor(colors.Inside),
	)

	rl.DrawTextEx(
		font,
		text,
		rl.Vector2{
			X: float32(bound.X+int32(bounds.Width)/2-textWidth),
			Y: float32(bound.Y)+bounds.Height/2-float32(textHeight/4),
		},
		size,
		0,
		rg.GetStyleColor(colors.Text),
	)

	return state == rg.Clicked
}

package gui

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// comboboxColoring describes the per-state colors for a Combobox control.
type comboboxColoring struct {
	Border rg.Property
	Inside rg.Property
	List   rg.Property
	Text   rg.Property
}

// comboboxColors lists the styling for each supported state.
var comboboxColors = map[rg.ControlState]comboboxColoring{
	rg.Normal:  {rg.ButtonDefaultBorderColor, rg.ComboboxDefaultInsideColor, rg.ComboboxDefaultListTextColor, rg.ComboboxDefaultTextColor},
	rg.Clicked: {rg.ComboboxDefaultBorderColor, rg.ComboboxDefaultInsideColor, rg.ComboboxDefaultListTextColor, rg.ComboboxDefaultTextColor},
	rg.Focused: {rg.ComboboxHoverBorderColor, rg.ComboboxHoverInsideColor, rg.ComboboxHoverListTextColor, rg.ComboboxHoverTextColor},
	rg.Pressed: {rg.ComboboxPressedBorderColor, rg.ComboboxPressedInsideColor, rg.ComboboxPressedListTextColor, rg.ComboboxPressedTextColor},
}

// ComboBox draws a simplified version of a ComboBox with custom fonting
func ComboBoxEx(font rl.Font, bounds rl.Rectangle, textChoices []string, activeChoice int, size float32) int {
	// Reject invalid selections and disable rendering.
	comboCount := len(textChoices)
	if activeChoice < 0 || activeChoice >= comboCount {
		rl.TraceLog(rl.LogWarning, "ComboBox active expects 0 <= active <= %d", comboCount)
		return -1
	}

	activeText := textChoices[activeChoice]

	// Calculate text dimensions.
	textSize := rl.MeasureTextEx(font, activeText, size, 0)
	textWidth := textSize.X
	textHeight := textSize.Y

	counter := rl.NewRectangle(bounds.X+bounds.Width-bounds.Height, bounds.Y, bounds.Height, bounds.Height)
	state := rg.GetInteractionState(counter)
	colors, exists := comboboxColors[state]

	if !exists {
		return activeChoice
	}

	if state == rg.Clicked {
		activeChoice = (activeChoice + 1) % comboCount
	}

	newBounds := rl.NewRectangle(bounds.X, bounds.Y, bounds.Width-bounds.Height, bounds.Height)
	rg.DrawBorderedRectangle(newBounds.ToInt32(), rg.GetStyle32(rg.ComboboxBorderWidth), rg.GetStyleColor(rg.ComboboxDefaultBorderColor), rg.GetStyleColor(rg.ComboboxDefaultInsideColor))
	rl.DrawTextEx(font, activeText, rl.Vector2{X: bounds.X + newBounds.Width/2 - textWidth/2, Y: bounds.Y + bounds.Height/2 - textHeight/2}, size, 0, rg.GetStyleColor(rg.ComboboxDefaultTextColor))

	// Render the accompanying "clicks" box showing the element counter.
	rg.DrawBorderedRectangle(counter.ToInt32(), rg.GetStyle32(rg.ButtonBorderWidth), rg.GetStyleColor(colors.Border), rg.GetStyleColor(colors.Inside))
	counterText := fmt.Sprint(activeChoice + 1)
	measureCounterText := rl.MeasureTextEx(font, counterText, size, 0)
	rl.DrawTextEx(
		font,
		counterText,
		rl.Vector2{
			X: counter.X + counter.Width/2 - measureCounterText.X/2,
			Y: counter.Y + counter.Height/2 - measureCounterText.Y/2,
		},
		size, 0, rg.GetStyleColor(colors.List),
	)

	return activeChoice
}

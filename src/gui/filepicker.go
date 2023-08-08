package gui

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ncruces/zenity"
)

// Pick a file and show it in the text box
func FilePicker(font rl.Font, bound rl.Rectangle, text string, size float32) string {
	textSize := rl.MeasureTextEx(font, text, size, 0)
	textBound := rl.NewRectangle(bound.X, bound.Y, bound.Width-bound.Height, bound.Height)

	rg.DrawBorderedRectangle(textBound.ToInt32(), rg.GetStyle32(rg.ComboboxBorderWidth), rg.GetStyleColor(rg.ComboboxDefaultBorderColor), rg.GetStyleColor(rg.ComboboxDefaultInsideColor))
	if textSize.X > bound.Width {
		var chars int
		for i := 0; rl.MeasureTextEx(font, text[i:], size, 0).X > bound.Width-bound.Height-30; i++ {
			chars++
		}
		textSize = rl.MeasureTextEx(font, text[chars:], size, 0)
		rl.DrawTextEx(font, fmt.Sprint("...", text[chars:][3:]), rl.Vector2{X: textBound.X + textBound.Width/2 - textSize.X/2, Y: textBound.Y + textBound.Height/2 - textSize.Y/2}, size, 0, rg.GetStyleColor(rg.ComboboxDefaultTextColor))
	} else {
		rl.DrawTextEx(font, text, rl.Vector2{X: textBound.X + textBound.Width/2 - textSize.X/2, Y: textBound.Y + textBound.Height/2 - textSize.Y/2}, size, 0, rg.GetStyleColor(rg.ComboboxDefaultTextColor))
	}

	if ButtonEx(font, rl.NewRectangle(bound.X+bound.Width-bound.Height, bound.Y, bound.Height, bound.Height), "+", size) {
		if choice, err := zenity.SelectFile(zenity.Filename(""), zenity.Directory()); err == nil {
			return fmt.Sprintf("%s/%s", choice, "settings.json")
		}
	}

	return text
}

package gui

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO
func ButtonIcon(icon rl.Texture2D, font rl.Font,
	boundIcon rl.Rectangle, boundText rl.Rectangle,
	text string, size float32, spacing float32) bool {

//	textSize := rl.MeasureTextEx(font, text, size, spacing)
	stateIcon := rg.GetInteractionState(boundIcon)
	stateTextRect := rg.GetInteractionState(boundText)

	colors, exist := buttonColors[stateTextRect]
	if !exist {
		return false
	}

	boundTextInt := boundText.ToInt32()
	boundIconInt := boundIcon.ToInt32()
	rg.DrawBorderedRectangle(boundTextInt, rg.GetStyle32(rg.ButtonBorderWidth),rg.GetStyleColor(colors.Border),rg.GetStyleColor(colors.Inside))
	rl.DrawTexture(icon, boundIconInt.X, boundIconInt.Y, rg.GetStyleColor(colors.Inside))

//	rl.DrawTextEx(
//		font,
//		text,
//		rl.Vector2{
//			X: float32(bound.X + int32(bounds.Width)/2 - textWidth),
//			Y: float32(bound.Y) + bounds.Height/2 - float32(textHeight/4),
//		},
//		size,
//		0,
//		rg.GetStyleColor(colors.Text),
//	)

	return stateIcon == rg.Clicked
}

package gui

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Draw the logo in the top left corner
func DrawLogoTopLeft(icon rl.Texture2D, font rl.Font, iconBounds rl.Rectangle, textBounds rl.Rectangle, size float32) {
	measurement := rl.MeasureTextEx(font, "MINES-GO", size, 0)
	rl.DrawRectangleRec(iconBounds, rg.BackgroundColor())
	rl.DrawRectangleRec(textBounds, rg.BackgroundColor())
	rl.DrawTexturePro(icon, rl.NewRectangle(0, 0, 32, 32), iconBounds, rl.Vector2{}, 0, rl.White)
	rl.DrawTextEx(font, "MINES-GO", rl.Vector2{
		X: textBounds.X,
		Y: textBounds.Y+textBounds.Height/2-measurement.Y/2,
	}, size, 0, rg.GetStyleColor(rg.LabelTextColor))
}

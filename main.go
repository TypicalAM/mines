package main

import (
	"embed"
	"example/raylib-game/src/game"
	shared "example/raylib-game/src/screens"
)

//go:embed resources/fonts/*
var fonts embed.FS

//go:embed resources/icons/*
var icons embed.FS

//go:embed resources/themes/*
var themes embed.FS

// Initialize game
func main() {
	shared.FontsFS = fonts
	shared.IconsFS = icons
	shared.ThemesFS = themes
	game.InitalizeGame()
}

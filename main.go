package main

import (
	"embed"
	"example/raylib-game/src/game"
	shared "example/raylib-game/src/screens"
)

//go:embed resources/*
var resources embed.FS

// Initialize game
func main() {
	shared.ResourcesFS = resources
	game.InitalizeGame()
}

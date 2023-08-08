package main

import (
	"embed"
	"github.com/TypicalAM/mines/src/game"
	shared "github.com/TypicalAM/mines/src/screens"
)

//go:embed resources/*
var resources embed.FS

// Initialize game
func main() {
	shared.ResourcesFS = resources
	game.InitalizeGame()
}

package models

import (
	"go_game_1/config"
	"math/rand"
)

type Prey struct {
	Coords
}

func NewPrey() *Prey {
	return &Prey{
		Coords{
			X: rand.Intn(config.WindowWidth / config.TileSize),
			Y: rand.Intn(config.WindowHeight / config.TileSize),
		},
	}
}

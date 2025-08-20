package utils

import (
	"go_game_1/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func Draw(screen *ebiten.Image, X, Y int, color color.RGBA) {
	vector.DrawFilledRect(
		screen,
		float32(X*config.TileSize),
		float32(Y*config.TileSize),
		config.TileSize,
		config.TileSize,
		color,
		false,
	)
}

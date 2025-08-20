package config

import "image/color"

const (
	WindowWidth  = 320
	WindowHeight = 240
	TileSize     = 5
)

var (
	SnakeColor      = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	PreyColor       = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	BackgroundColor = color.RGBA{R: 0, G: 0, B: 0, A: 255}
)

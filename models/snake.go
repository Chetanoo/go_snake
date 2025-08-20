package models

import "go_game_1/config"

type Coords struct {
	X, Y int
}

type Snake struct {
	Body        []Coords
	Direction   Coords
	GrowCounter int
}

func NewSnake() *Snake {
	return &Snake{
		Body: []Coords{
			{X: config.WindowWidth / config.TileSize / 2, Y: config.WindowHeight / config.TileSize / 2},
		},
		Direction: Coords{X: 1, Y: 0},
	}
}

func (s *Snake) Move() {
	NewHead := Coords{
		X: s.Body[0].X + s.Direction.X,
		Y: s.Body[0].Y + s.Direction.Y,
	}
	s.Body = append([]Coords{NewHead}, s.Body...)

	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}

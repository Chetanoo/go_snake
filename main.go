package main

import (
	"fmt"
	"go_game_1/config"
	"go_game_1/models"
	"go_game_1/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	snake         *models.Snake
	prey          *models.Prey
	speed         int
	score         int
	growthCounter int
	gameOver      bool
	updateCounter int
	bestScore     int
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.snake = models.NewSnake()
			g.prey = models.NewPrey()
			g.score = 0
			g.gameOver = false
		}
		return nil
	}

	g.updateCounter++
	if g.updateCounter < g.speed {
		return nil
	}
	g.updateCounter = 0

	g.snake.Move()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.snake.Direction.X == 0 {
		g.snake.Direction = models.Coords{X: -1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.snake.Direction.X == 0 {
		g.snake.Direction = models.Coords{X: 1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) && g.snake.Direction.Y == 0 {
		g.snake.Direction = models.Coords{X: 0, Y: -1}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.snake.Direction.Y == 0 {
		g.snake.Direction = models.Coords{X: 0, Y: 1}
	}

	head := g.snake.Body[0]

	if head.X < 0 || head.Y < 0 || head.X >= config.WindowWidth/config.TileSize || head.Y >= config.WindowHeight/config.TileSize {
		g.SetGameOver()
	}

	for _, coords := range g.snake.Body[1:] {
		if head.X == coords.X && head.Y == coords.Y {
			g.SetGameOver()
		}
	}

	if head.X == g.prey.Coords.X && head.Y == g.prey.Coords.Y {
		g.snake.GrowCounter++
		g.score++
		g.prey = models.NewPrey()

		if g.speed > 2 {
			g.speed--
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(config.BackgroundColor)
	for _, coords := range g.snake.Body {
		utils.Draw(screen, coords.X, coords.Y, config.SnakeColor)
	}
	utils.Draw(screen, g.prey.Coords.X, g.prey.Coords.Y, config.PreyColor)

	font := basicfont.Face7x13
	if g.gameOver {
		text.Draw(screen, "Game Over", font, config.WindowWidth/2-40, config.WindowHeight/2, config.SnakeColor)
		text.Draw(screen, fmt.Sprintf("Best Score: %d", g.bestScore), font, config.WindowWidth/2-40, config.WindowHeight/2+40, config.SnakeColor)
		text.Draw(screen, "Press R to restart", font, config.WindowWidth/2-40, config.WindowHeight/2+20, config.SnakeColor)
	}
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, font, 5, config.WindowHeight-5, config.SnakeColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) SetGameOver() {
	g.gameOver = true
	if g.bestScore < g.score {
		g.bestScore = g.score
	}
	g.speed = 10
}

func main() {
	game := &Game{
		snake:         models.NewSnake(),
		prey:          models.NewPrey(),
		speed:         10,
		score:         0,
		growthCounter: 0,
		gameOver:      false,
	}
	ebiten.SetWindowSize(config.WindowWidth*2, config.WindowHeight*2)
	ebiten.SetWindowTitle("Sneko!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

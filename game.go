package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	paddles []Paddle
	ball    Ball
	scores  []int
}

// Control the size of the window
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Draw objects to the screen
func (g *Game) Draw(screen *ebiten.Image) {
	for _, paddle := range g.paddles {
		vector.DrawFilledRect(
			screen,
			float32(paddle.X), float32(paddle.Y),
			float32(paddle.W), float32(paddle.H),
			paddle.Color, false,
		)
	}

	vector.DrawFilledRect(
		screen,
		float32(g.ball.X), float32(g.ball.Y),
		float32(g.ball.W), float32(g.ball.H),
		color.White, false,
	)

	scoreTextP1 := "Score P1: " + fmt.Sprint(g.scores[0])
	text.Draw(screen, scoreTextP1, basicfont.Face7x13, screenWidth / 2 - 200, 20, color.White)

	scoreTextP2 := "Score P2: " + fmt.Sprint(g.scores[1])
	text.Draw(screen, scoreTextP2, basicfont.Face7x13, screenWidth / 2 + 100, 20, color.White)
}

// Update the game state
func (g *Game) Update() error {
	for idx, _ := range g.paddles {
		g.paddles[idx].MoveOnKeyPress()
	}
	g.ball.Move()
	g.CollideWithWall()
	g.CollideWithPaddle()
	g.PaddleWallCollision()
	return nil
}

func (g *Game) Reset() {
	g.ball.X = screenWidth / 2
	g.ball.Y = screenHeight / 2
 
	g.ball.RandomDirection()
}

func (g *Game) CollideWithWall() {
	if g.ball.X >= screenWidth {
		g.scores[0]++
		g.Reset()
	} else if g.ball.X <= 0 {
		g.scores[1]++
		g.Reset()
	}

	if g.ball.Y >= screenHeight {
		g.ball.YT = -ballSpeed
	} else if g.ball.Y <= 0 {
		g.ball.YT = ballSpeed
	}
}

func (g *Game) CollideWithPaddle() {
	for idx, _ := range g.paddles {
		if 
			g.ball.X+g.ball.W >= g.paddles[idx].X && 
			g.ball.X <= g.paddles[idx].X+g.paddles[idx].W && 
			g.ball.Y+g.ball.H >= g.paddles[idx].Y &&
			g.ball.Y <= g.paddles[idx].Y+g.paddles[idx].H {
				g.ball.XT = -g.ball.XT 
		}
	 
	}
}

func (g *Game) PaddleWallCollision() {
	for idx, _ := range g.paddles {
		if g.paddles[idx].Y+g.paddles[idx].H >= screenHeight {
			g.paddles[idx].Y = screenHeight - g.paddles[idx].H
		} else if g.paddles[idx].Y <= 0 {
			g.paddles[idx].Y = 0
		}
	}
}
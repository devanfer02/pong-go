package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
Generic Object
- This is a generic object that can be used to represent any object in the game
*/
type Object struct {
	X, Y float32 // Position of the object
	W, H float32 // Width and height of the object
}

type Paddle struct {
	Object
	PlayerType PaddlePlayerType 
	Color color.Color
}

type Ball struct {
	Object
	XT, YT float32 // Velocity per tick
}

func (p *Paddle) MoveOnKeyPress() {
	if p.PlayerType == Player1 {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			p.Y -= paddleSpeed
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			p.Y += paddleSpeed
		}
	} else {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.Y -= paddleSpeed
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.Y += paddleSpeed
		}
	}
		
}

func (b *Ball) Move() {
	b.X += b.XT
	b.Y += b.YT
}

func (b *Ball) RandomDirection() {
	
	if rand.Intn(10) % 2 == 0 {
		b.XT = -b.XT
	} 
	if rand.Intn(10) % 2 == 0{
		b.YT = -b.YT
	}
}
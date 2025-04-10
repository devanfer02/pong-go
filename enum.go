package main

import "image/color"

type PaddlePlayerType = string 

const (
	Player1 PaddlePlayerType = "Player 1"
	Player2 PaddlePlayerType = "Player 2"
)

var (
	Red    = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	Green  = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	Blue   = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	Yellow = color.RGBA{R: 255, G: 255, B: 0, A: 255}
	Cyan   = color.RGBA{R: 0, G: 255, B: 255, A: 255}
	Magenta= color.RGBA{R: 255, G: 0, B: 255, A: 255}
	Black  = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	White  = color.RGBA{R: 255, G: 255, B: 255, A: 255}
)

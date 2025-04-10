package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	ballSpeed    = 4.5
	paddleSpeed  = 6
)

func main() {
	ebiten.SetWindowTitle("PongGo")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	g := Game{
		paddles: []Paddle{
			{
				Object: Object{
					X: 50,
					Y: 200,
					W: 15,
					H: 100,
				},
				PlayerType: Player1,
				Color: Red,
			},
			{
				Object: Object{
					X: 600,
					Y: 200,
					W: 15,
					H: 100,
				},
				PlayerType: Player2,
				Color: Blue,
			},

		},
		ball: Ball{
			Object: Object{
				X: screenWidth / 2,
				Y: screenHeight / 2,
				W: 15,
				H: 15,
			},	
			XT: ballSpeed,
			YT: ballSpeed,
		},
		scores: []int{0, 0},
	}
	
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal("Error running game:", err)
	}
}

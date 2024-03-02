package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	rectY = 0.0
	velocity = 0.0
	gravity = 0.1
	e = 0.9
)

type Game struct{}

func (g *Game) Update() error {
	// Update logic here
	rectY += velocity
	velocity += gravity
	if(rectY>420){
		velocity = -velocity*e
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the rectangle
	ebitenutil.DrawRect(screen, 220, rectY, 100, 50, color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Rectangle Animation")

	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

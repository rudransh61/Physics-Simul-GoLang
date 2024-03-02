package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480

	// Pendulum parameters
	ceilingX     = screenWidth / 2
	ceilingY     = 100
	pendulumSize = 20
	stringLength = 150
	initialAngle = math.Pi / 4 // 45 degrees in radians
)

type Position struct {
	X, Y float64
}

type Pendulum struct {
	Angle           float64
	AngularVelocity float64
	BobPosition     Position
}

func (p *Pendulum) Update() {
	// Simulate physics - simple pendulum motion with quadratic damping
	gravity := 1.0
	k := 0.001 // Damping coefficient
	damping := 1.0 / (1.0 + k*p.AngularVelocity*p.AngularVelocity)

	p.AngularVelocity += (-gravity/stringLength)*math.Sin(p.Angle)*damping
	p.Angle += p.AngularVelocity

	// Update bob position
	p.BobPosition.X = ceilingX + stringLength*math.Sin(p.Angle)
	p.BobPosition.Y = ceilingY + stringLength*math.Cos(p.Angle)
}

type Game struct {
	Pendulum *Pendulum
}

func (g *Game) Update() error {
	// Update pendulum logic here
	g.Pendulum.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the pendulum
	// Draw the string
	ebitenutil.DrawLine(screen, ceilingX, ceilingY, g.Pendulum.BobPosition.X, g.Pendulum.BobPosition.Y, color.White)

	// Draw the pendulum bob (circle)
	ebitenutil.DrawRect(screen, g.Pendulum.BobPosition.X-pendulumSize/2, g.Pendulum.BobPosition.Y-pendulumSize/2, pendulumSize, pendulumSize, color.RGBA{R: 0, G: 0, B: 255, A: 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pendulum Simulation")

	game := &Game{
		Pendulum: &Pendulum{
			Angle: initialAngle,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

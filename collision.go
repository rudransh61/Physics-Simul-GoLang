package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
	particleSize = 10
	numParticles = 10
	speed        = 2
)

type Particle struct {
	X, Y  float64
	VelX, VelY float64
	Color color.Color
}

type CollisionSimulator struct {
	Particles []Particle
}

func (c *CollisionSimulator) Update() {
	for i := range c.Particles {
		// Update particle position based on velocity
		c.Particles[i].X += c.Particles[i].VelX
		c.Particles[i].Y += c.Particles[i].VelY

		// Check for collisions with walls
		if c.Particles[i].X < 0 || c.Particles[i].X > screenWidth {
			c.Particles[i].VelX *= -1 // Reverse velocity on collision with horizontal walls
		}
		if c.Particles[i].Y < 0 || c.Particles[i].Y > screenHeight {
			c.Particles[i].VelY *= -1 // Reverse velocity on collision with vertical walls
		}
	}
}

func (c *CollisionSimulator) Draw(screen *ebiten.Image) {
	for _, particle := range c.Particles {
		x := particle.X
		y := particle.Y

		// Draw the particle
		ebitenutil.DrawRect(screen, x-particleSize/2, y-particleSize/2, particleSize, particleSize, particle.Color)
	}
}

func (c *CollisionSimulator) InitParticles() {
	for i := 0; i < numParticles; i++ {
		c.Particles = append(c.Particles, Particle{
			X:     rand.Float64() * screenWidth,
			Y:     rand.Float64() * screenHeight,
			VelX:  (rand.Float64() - 0.5) * 2 * speed,
			VelY:  (rand.Float64() - 0.5) * 2 * speed,
			Color: color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255},
		})
	}
}

type Game struct {
	CollisionSimulator *CollisionSimulator
}

func (g *Game) Update() error {
	g.CollisionSimulator.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.CollisionSimulator.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Collision Simulator")

	collisionSimulator := &CollisionSimulator{}
	collisionSimulator.InitParticles()

	game := &Game{
		CollisionSimulator: collisionSimulator,
	}

	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

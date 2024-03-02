package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth    = 800
	screenHeight   = 600
	particleSize   = 10
	numParticles   = 5
	particleOffset = 50
)

type Particle struct {
	X     float64
	Y     float64
	Speed float64
	Phase float64
	Color color.Color
}

type SineWaveSimulation struct {
	Particles []Particle
}

func (s *SineWaveSimulation) Update() {
	// Update logic for the traveling sine wave here
	for i := range s.Particles {
		s.Particles[i].X += s.Particles[i].Speed
		s.Particles[i].Y = math.Sin((s.Particles[i].X+s.Particles[i].Phase)/50) * 100 + screenHeight/2
	}
}

func (s *SineWaveSimulation) Draw(screen *ebiten.Image) {
	// Draw logic for the traveling sine wave here
	for _, particle := range s.Particles {
		x := particle.X
		y := particle.Y

		// Draw the particle
		ebitenutil.DrawRect(screen, x-particleSize/2, y-particleSize/2, particleSize, particleSize, particle.Color)
	}
}

type Game struct {
	SineWaveSimulation *SineWaveSimulation
}

func (g *Game) Update() error {
	g.SineWaveSimulation.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SineWaveSimulation.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Multiple Particles Sine Wave Simulation")

	sineWaveSimulation := &SineWaveSimulation{
		Particles: make([]Particle, numParticles),
	}

	// Initialize particles with different phases
	for i := range sineWaveSimulation.Particles {
		sineWaveSimulation.Particles[i] = Particle{
			X:     float64(i * particleOffset),
			Y:     screenHeight / 2,
			Speed: 2,
			Phase: float64(i) * math.Pi * 2 / float64(numParticles),
			Color: color.White,
		}
	}

	game := &Game{
		SineWaveSimulation: sineWaveSimulation,
	}

	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

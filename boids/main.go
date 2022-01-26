package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ebiten is a Go graphics library

// area of our game simulation
const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015 // after trial and error
	// viewRadius is no of pixel each boid can see
	// adjRate, to make transition smooth and not very speed
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	// this will be memory shared b/w threads, also store boid id,
	// used to communicate b/w boid relative postion

)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

// ebiten will call Draw a number of times per second to update the graphics on screen
func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.postion.x+1), int(boid.postion.y), green)
		screen.Set(int(boid.postion.x-1), int(boid.postion.y), green)
		screen.Set(int(boid.postion.x), int(boid.postion.y-1), green)
		screen.Set(int(boid.postion.x), int(boid.postion.y+1), green)
		// these 4 lines make a diamond Boid

	}
}

// when someone calls for screen Layout
func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in my room")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

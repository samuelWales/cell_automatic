package life_nodeath

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const cell_side = 5
const CAPACITY_X = 210
const CAPACITY_Y = 150

const WIDTH = 5 * 210
const HEIGHT = 5 * 150

var plane [CAPACITY_X][CAPACITY_Y]int = [CAPACITY_X][CAPACITY_Y]int{}
var buffer [CAPACITY_X][CAPACITY_Y]int = [CAPACITY_X][CAPACITY_Y]int{}
var counter int = 0

type Game struct{}

func Rule(x, y int) {
	n := plane[x-1][y-1] + plane[x-1][y+0] + plane[x-1][y+1] + plane[x+0][y-1] + plane[x+0][y+1] + plane[x+1][y-1] + plane[x+1][y+0] + plane[x+1][y+1]
	if plane[x][y] == 0 && n == 3 {
		buffer[x][y] = 1
	} else if n == 0 || n == 1 || n == 2 || n == 3 || n == 5 || n == 6 || n == 7 || n == 8 {
		buffer[x][y] = plane[x][y]
	} else {
		buffer[x][y] = 0
	}
}

// Logic
func (g *Game) Update() error {
	for x := 1; x < CAPACITY_X-1; x++ {
		for y := 1; y < CAPACITY_Y-1; y++ {
			go Rule(x, y)
		}
	}
	plane = buffer

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, WIDTH, HEIGHT, color.Black, false)

	cell_color := color.RGBA{52, 201, 36, 100}
	for x := 1; x < CAPACITY_X-1; x++ {
		for y := 1; y < CAPACITY_Y-1; y++ {
			if plane[x][y] > 0 {
				rect_x := (x - 1) * cell_side
				rect_y := (y - 1) * cell_side
				width := cell_side
				height := cell_side
				vector.DrawFilledRect(screen, float32(rect_x), float32(rect_y), float32(width), float32(height), cell_color, false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}

func LifeNoDeath() {
	for x := 1; x < CAPACITY_X-1; x++ {
		for y := 1; y < CAPACITY_Y-1; y++ {
			if rand.Float32() < 0.05 {
				plane[x][y] = 1
			}
		}
	}

	ebiten.SetTPS(7)
	game := &Game{}
	ebiten.SetWindowTitle("Conway's Game of Life")
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

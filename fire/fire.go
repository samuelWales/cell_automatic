package fire

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
const EMPTY = 0
const TREE = 1
const FIRE = 2
const P = 0.03   // вероятность того, что в пусто клетке вырастет дерево
const F = 0.0001 // вероятность того, что деоево загорится от других факторов

const WIDTH = 5 * 210
const HEIGHT = 5 * 150

var img *ebiten.Image

// const STROKE_WIDTH = 1
var plane [CAPACITY_X][CAPACITY_Y]int = [CAPACITY_X][CAPACITY_Y]int{}
var buffer [CAPACITY_X][CAPACITY_Y]int = [CAPACITY_X][CAPACITY_Y]int{}
var neighbours [8][2]int = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
var neighbours_x [8]int = [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
var neighbours_y [8]int = [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
var wind [8]float32 = [8]float32{0.1, 1, 1, 0.1, 1, 0.1, 0.1, 0.1} // северо-восточное направление ветра

type Game struct{}

func Rule(x, y int) {
	if plane[x][y] == TREE {
		var index int = 0
		for _, neighbour := range neighbours {
			dx := neighbour[0]
			dy := neighbour[1]
			if plane[x+dx][y+dy] == FIRE && rand.Float32() <= wind[index] {
				buffer[x][y] = FIRE
				break
			}
			index++
		}
		if rand.Float32() <= F {
			buffer[x][y] = FIRE
		}
	}
	if plane[x][y] == FIRE {
		buffer[x][y] = EMPTY
	}
	if plane[x][y] == EMPTY {
		if rand.Float32() <= P {
			buffer[x][y] = TREE
		} else {
			buffer[x][y] = plane[x][y]
		}
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
	earth_color := color.RGBA{45, 26, 19, 100}
	vector.DrawFilledRect(screen, 0, 0, WIDTH, HEIGHT, earth_color, false)

	tree_color := color.RGBA{33, 124, 29, 100}
	fire_color := color.RGBA{229, 165, 10, 100}
	for x := 1; x < CAPACITY_X-1; x++ {
		for y := 1; y < CAPACITY_Y-1; y++ {
			if plane[x][y] == TREE {
				rect_x := (x - 1) * cell_side
				rect_y := (y - 1) * cell_side
				width := cell_side
				height := cell_side
				vector.DrawFilledRect(screen, float32(rect_x), float32(rect_y), float32(width), float32(height), tree_color, false)
			}
			if plane[x][y] == FIRE {
				rect_x := (x - 1) * cell_side
				rect_y := (y - 1) * cell_side
				width := cell_side
				height := cell_side
				vector.DrawFilledRect(screen, float32(rect_x), float32(rect_y), float32(width), float32(height), fire_color, false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}

func Fire() {
	for x := 1; x < CAPACITY_X-1; x++ {
		for y := 1; y < CAPACITY_Y-1; y++ {
			if rand.Float32() < 0.6 {
				plane[x][y] = TREE
			} else {
				plane[x][y] = EMPTY
			}
		}
	}

	ebiten.SetTPS(7)
	game := &Game{}
	ebiten.SetWindowTitle("Fire model")
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

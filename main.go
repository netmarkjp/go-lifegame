package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

// Screen of lifegame
type Screen struct {
	width             int
	height            int
	currentGeneration [][]bool // x, y
	nextGeneration    [][]bool // x, y
	output            *bufio.Writer
}

// Init screen
func (s *Screen) Init(width, height int) {
	s.width = width
	s.height = height
	s.initCurrentGeneration(width, height)
	s.initNextGeneration(width, height)
	bufSize := (width*height*2+width)/4096 + 1
	s.output = bufio.NewWriterSize(os.Stdout, bufSize*4096)
}

func (s *Screen) initCurrentGeneration(width, height int) {
	s.currentGeneration = make([][]bool, height)
	for i := 0; i < len(s.currentGeneration); i++ {
		s.currentGeneration[i] = make([]bool, width)
	}
}

func (s *Screen) initNextGeneration(width, height int) {
	s.nextGeneration = make([][]bool, height)
	for i := 0; i < len(s.nextGeneration); i++ {
		s.nextGeneration[i] = make([]bool, width)
	}
}

// CurrentAlive returns liveness of current generation
func (s *Screen) CurrentAlive(x, y int) bool {
	return s.currentGeneration[y][x]
}

func (s *Screen) progressCell(x, y int) bool {

	// (x-1, y-1) (x, y-1) (x+1, y-1)
	// (x-1, y)	  (x, y)   (x+1, y)
	// (x-1, y+1) (x, y+1) (x+1, y+1)

	alive := 0
	dead := 0

	cells := [][]int{
		{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
		{x - 1, y}, {x + 1, y},
		{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
	}

	for _, cell := range cells {
		if cell[0] < 0 || cell[1] < 0 {
			dead++
			continue
		} else if cell[0] >= s.width-1 || cell[1] >= s.height-1 {
			dead++
			continue
		}
		if s.CurrentAlive(cell[0], cell[1]) {
			alive++
		} else {
			dead++
		}
	}

	if alive == 3 {
		return true
	} else if s.CurrentAlive(x, y) && (alive == 2 || alive == 3) {
		return true
	} else if alive <= 1 {
		return false
		// } else if alive >= 4 {
	} // else {
	return false
	//}
}

func (s *Screen) stepGeneration() {
	s.initNextGeneration(s.width, s.height)
	for y, linedata := range s.nextGeneration {
		for x := range linedata {
			s.nextGeneration[y][x] = s.progressCell(x, y)
		}
	}
	copy(s.currentGeneration, s.nextGeneration)
}

// SetInitialAlive set alive cell at start
func (s *Screen) SetInitialAlive(x, y int) {
	s.currentGeneration[y][x] = true
}

// Render render screen
func (s *Screen) Render() {
	for {
		s.output.WriteString("\033[H\033[2J")
		for _, linedata := range s.currentGeneration {
			for _, alive := range linedata {
				if alive {
					s.output.WriteString("＊")
				} else {
					s.output.WriteString("  ")
				}
			}
			s.output.WriteString("\n")
		}
		s.output.Flush()

		s.stepGeneration()
		time.Sleep(100 * time.Millisecond)
	}
}

func example1() {
	screen := new(Screen)
	screen.Init(50, 40)

	var n, m int

	// ブロック
	n = 2
	m = 2
	screen.SetInitialAlive(n, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n+1, m+1)

	//蜂の巣
	n = 10
	m = 2
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n+2, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n+3, m+1)
	screen.SetInitialAlive(n+1, m+2)
	screen.SetInitialAlive(n+2, m+2)

	//ボート
	n = 20
	m = 2
	screen.SetInitialAlive(n, m)
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n+2, m+1)
	screen.SetInitialAlive(n+1, m+2)

	//船
	n = 30
	m = 2
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n+2, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n+2, m+1)
	screen.SetInitialAlive(n, m+2)
	screen.SetInitialAlive(n+1, m+2)

	//池
	n = 40
	m = 2
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n+2, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n+3, m+1)
	screen.SetInitialAlive(n, m+2)
	screen.SetInitialAlive(n+3, m+2)
	screen.SetInitialAlive(n+1, m+3)
	screen.SetInitialAlive(n+2, m+3)

	//ブリンカー
	n = 2
	m = 10
	screen.SetInitialAlive(n, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n, m+2)

	//ヒキガエル
	n = 10
	m = 10
	screen.SetInitialAlive(n, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n, m+2)
	screen.SetInitialAlive(n+1, m+1)
	screen.SetInitialAlive(n+1, m+2)
	screen.SetInitialAlive(n+1, m+3)

	//ビーコン
	n = 20
	m = 10
	screen.SetInitialAlive(n, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n+1, m+1)
	n = 22
	m = 12
	screen.SetInitialAlive(n, m)
	screen.SetInitialAlive(n, m+1)
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n+1, m+1)

	screen.Render()
}

func example2() {
	screen := new(Screen)
	x := 80
	y := 80
	screen.Init(x, y)

	rand.Seed(20171114)
	for ix := 0; ix < x; ix++ {
		for iy := 0; iy < y; iy++ {
			if rand.Intn(10000) < 2000 {
				screen.SetInitialAlive(ix, iy)
			}
		}
	}

	screen.Render()
}

func example3() {
	screen := new(Screen)
	x := 40
	y := 40
	screen.Init(x, y)
	var n, m int

	//グライダー
	n = 5
	m = 5
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n+2, m+1)
	screen.SetInitialAlive(n, m+2)
	screen.SetInitialAlive(n+1, m+2)
	screen.SetInitialAlive(n+2, m+2)

	screen.Render()
}

func example4() {
	screen := new(Screen)
	x := 20
	y := 20
	screen.Init(x, y)
	var n, m int

	//長さ10の直線
	n = 5
	m = 5
	screen.SetInitialAlive(n, m)
	screen.SetInitialAlive(n+1, m)
	screen.SetInitialAlive(n+2, m)
	screen.SetInitialAlive(n+3, m)
	screen.SetInitialAlive(n+4, m)
	screen.SetInitialAlive(n+5, m)
	screen.SetInitialAlive(n+6, m)
	screen.SetInitialAlive(n+7, m)
	screen.SetInitialAlive(n+8, m)
	screen.SetInitialAlive(n+9, m)

	screen.Render()
}

func main() {
	example4()
}

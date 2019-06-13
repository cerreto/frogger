package main

import . "g2d"
import "strconv"

var bestSc = 0
var Sc = 0
var h = 512
var w = 448
var screen = Size{w, h}
var x, y = 15, 15
var vx, vy = 5, 5
var frog = LoadImage("https://i.imgur.com/I2cubui.png")
var car1 = LoadImage("https://i.imgur.com/aUYxC98.png")
var car2 = LoadImage("https://i.imgur.com/XW22W2b.png")
var car3 = LoadImage("https://i.imgur.com/8SLloTU.png")
var bkg = LoadImage("https://i.imgur.com/vAZz8KH.png")

//entity class
type ent struct {
	x, y   int
	h, w   int
	ix, iy int
	dx, dy int
}

//creates entities
func newent(pos Point, h, w int, ix, iy int, dx, dy int) *ent {
	b := &ent{pos.X, pos.Y, h, w, ix, iy, dx, dy}
	return b
}

//moves entities
func (b *ent) Move() {
	if b.x > (w + 200) {
		b.x = -100
	}
	if b.x < (-200) {
		b.x = w + 100
	}
	b.x += b.dx
	if (fr.x > b.x && fr.x < b.x+b.w) && fr.y == b.y {
		fr.x = fr.ix
		fr.y = fr.iy
		Sc = 0
	}
}

func environment() {
	b1.Move()
	b2.Move()
	b3.Move()
	b4.Move()
	b5.Move()
    b6.Move()
    DrawImage(bkg, Point{0, 0})
	DrawImage(car1, Point{b1.x, b1.y})
	DrawImage(car2, Point{b2.x, b2.y})
	DrawImage(car3, Point{b3.x, b3.y})
	DrawImage(car3, Point{b4.x, b4.y})
	DrawImage(car1, Point{b5.x, b5.y})
    DrawImage(car2, Point{b6.x, b6.y})
}

var b1 = newent(Point{70, h * 16 / 32}, 32, 64, 70, h*16/32, 10, 0)
var b2 = newent(Point{50, h * 12 / 32}, 32, 32, 50, h*12/32, -10, 0)
var b3 = newent(Point{40, h * 14 / 32}, 32, 32, 40, h*14/32, -5, 0)
var b4 = newent(Point{100, h * 6 / 32}, 32, 32, 40, h*6/32, -6, 0)
var b5 = newent(Point{200, h * 8 / 32}, 32, 64, 70, h*8/32, 10, 0)
var b6 = newent(Point{100, h * 22 / 32}, 32, 32, 50, h*22/32, -10, 0)
var fr = newent(Point{ ( w / 2 ) - 8 , h * 26 / 32}, 32, 32, (w+32)/2, h*26/32, 0, 0)

var jump = 32

func tic() {
	ClearCanvas()
	environment()
	if KeyPressed("w") || KeyPressed("ArrowUp") {
		fr.y -= jump
	}
	if KeyPressed("s") || KeyPressed("ArrowDown") {
		if fr.y+fr.h+jump <= h {
			fr.y += jump
		}
	}
	if KeyPressed("a") || KeyPressed("ArrowLeft") {
		if 0 <= fr.x-jump {
			fr.x -= jump
		}
	}
	if KeyPressed("d") || KeyPressed("ArrowRight") {
		if fr.x+fr.w+jump <= w {
			fr.x += jump
		}
	}
	if fr.y+fr.h < w*15/100 {
		Sc++
		fr.x = fr.ix
		fr.y = fr.iy
		if Sc > bestSc {
			bestSc = Sc
		}
	}
	DrawImage(frog, Point{fr.x, fr.y})
    SetColor(Color{255, 255, 255})
	DrawText("Best Score: " + strconv.Itoa(bestSc), Point{w * 70 / 100, h * 85 / 100}, w*4/100)
	DrawText("Score: " + strconv.Itoa(Sc), Point{w * 70 / 100, h * 90 / 100}, w*4/100)
	DrawLine(Point{0, w * 86 / 100}, Point{h, w * 86 / 100})
}

func main() {
	InitCanvas(screen)
	SetFrameRate(60)
	MainLoop(tic)
}

package main 

import . "g2d"
import "strconv"
var bestSc = 0
var Sc = 0
var h,w=512,448
var screen=Size{h,w}
var x,y=15,15
var vx,vy=5,5
var frog=LoadImage("./assets/frog.png")
var car1=LoadImage("./assets/truck-right.png")
var car2=LoadImage("./assets/car2-left.png")
var car3=LoadImage("./assets/car3-left.png")

//entity class
type ent struct {
    x, y    int
    h, w    int
    ix,iy 	int
    dx, dy  int
}

//creates entities
func newent(pos Point, h,w int, ix,iy int, dx,dy int) *ent {
    b := &ent{pos.X, pos.Y, h, w, ix,iy, dx, dy}
    return b
}

//moves entities
func (b *ent) Move() {
    if b.x>(w+200) {
        b.x=-50
    }
    if b.x<(-200) {
        b.x=w+50
    }
    b.x += b.dx
    if (b.x <= fr.x && fr.x<=b.x+b.w) && b.y==fr.y {
    	fr.x=fr.ix
    	fr.y=fr.iy
    	Sc = 0
    }
}

func environment() {
	b1.Move()
	b2.Move()
	b3.Move()
	DrawImage(car1,Point{b1.x,b1.y})
	DrawImage(car2,Point{b2.x,b2.y})
	DrawImage(car3,Point{b3.x,b3.y})
}

var b1=newent(Point{70,h*10/32},64,32,70,320,10,0)
var b2=newent(Point{50,h*12/32},32,32,50,480,-10,0)
var b3=newent(Point{40,h*14/32},32,32,40,544,-5,0)
var fr=newent(Point{(w+32)/2,h*26/32},32,32,(w+32)/2,h*26/32,0,0)
var jump=32
func tic() {
	ClearCanvas()
	environment()
	if KeyPressed("w") || KeyPressed("ArrowUp") {
		 fr.y-=jump
	}
	if KeyPressed("s") || KeyPressed("ArrowDown") {
		if (fr.y+fr.h+jump <= h-fr.h-jump) {fr.y+=jump}
	}
	if KeyPressed("a") || KeyPressed("ArrowLeft") {
		if (0<fr.x -(fr.w)/2&& fr.x<w-fr.w-fr.w) {fr.x-=jump}
	}
	if KeyPressed("d") || KeyPressed("ArrowRight") {
		if (0<fr.x && fr.x<w-jump-fr.w) {fr.x+=jump}
	}
	if fr.y+fr.h<w*15/100 {
		Sc ++
		fr.x=fr.ix
		fr.y=fr.iy
		if Sc>bestSc{
			bestSc = Sc
		}
	}
	DrawImage(frog,Point{fr.x,fr.y})
	DrawText("Best Sc:"+strconv.Itoa(bestSc), Point{w*80/100, w*88/100}, w*4/100)
	DrawText("Sc:" + strconv.Itoa(Sc), Point{w*80/100, w*92/100}, w*4/100)
	DrawLine(Point{0, w*86/100}, Point{w, w*86/100})

	for i := 0; i < 30; i++ {
		DrawLine(Point{i*w/25, w*15/100}, Point{i*(w/25)+w*2/100,w*15/100})
	}

}

func main() {
	InitCanvas(screen)
	SetFrameRate(60)
	SetColor(Color{0,0,0})
	MainLoop(tic)
}



package main 

import . "g2d"
import "strconv"
var bestSc = 0
var Sc = 0
var h,w=512,448
var screen=Size{w,h}
var x,y=15,15
var vx,vy=5,5
var frog=LoadImage("./assets/frog.png")
var car1=LoadImage("./assets/truck-right.png")
var car2=LoadImage("./assets/car2-left.png")
var car3=LoadImage("./assets/car3-left.png")
var log=LoadImage("./assets/log.png")

//entity class
type ent struct {
    x, y    int
    h, w    int
    ix,iy 	int
    dx, dy  int
    log 	bool
}


//creates entities
func newent(pos Point, h,w int, ix,iy int, dx,dy int, log bool) *ent {
    b := &ent{pos.X, pos.Y, h, w, ix,iy, dx, dy,log}
    return b
}

//moves entities
func (b *ent) Move() {
    if b.x>(w+200) {
        b.x=-100
    }
    if b.x<(-200) {
        b.x=w+100
    }
    b.x += b.dx
	if (fr.x>b.x && fr.x<b.x+b.w) && fr.y==b.y {
	   if b.log==false {
	   	fr.x=fr.ix
	   	fr.y=fr.iy
	    Sc = 0
	   } else {
	   	fr.x+=b.dx
	   	if fr.x<0 || fr.x+fr.w>w {
	   		fr.x=fr.ix
	   		fr.y=fr.iy
	    	Sc = 0
	   	}
	   }
	   }
	}

func environment() {
	b1.Move()
	b2.Move()
	b3.Move()
	l1.Move()
	l2.Move()
	DrawImage(car1,Point{b1.x,b1.y})
	DrawImage(car2,Point{b2.x,b2.y})
	DrawImage(car3,Point{b3.x,b3.y})
	DrawImage(log,Point{l1.x,l1.y})
	DrawImage(log,Point{l2.x,l2.y})
}

var b1=newent(Point{70,h*10/32},32,64,70,h*10/32,10,0,false)
var b2=newent(Point{50,h*12/32},32,32,50,h*12/32,-10,0,false)
var b3=newent(Point{40,h*14/32},32,32,40,h*14/32,-5,0,false)
var fr=newent(Point{(w+32)/2,h*26/32},32,32,(w+32)/2,h*26/32,0,0,false)
var l1=newent(Point{100,h*6/32},32,96,100,h*6/32,-3,0,true)
var l2=newent(Point{200,h*4/32},32,96,200,h*4/32,3,0,true)

var jump=32
func tic() {
	ClearCanvas()
	environment()
	if KeyPressed("w") || KeyPressed("ArrowUp") {
		 fr.y-=jump
	}
	if KeyPressed("s") || KeyPressed("ArrowDown") {
		if (fr.y+fr.h+jump<=h) {fr.y+=jump}
	}
	if KeyPressed("a") || KeyPressed("ArrowLeft") {
		if (0<=fr.x-jump) {fr.x-=jump}
	}
	if KeyPressed("d") || KeyPressed("ArrowRight") {
		if (fr.x+fr.w+jump<=w) {fr.x+=jump}
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
	DrawText("Best Score:"+strconv.Itoa(bestSc), Point{w*70/100, h*85/100}, w*4/100)
	DrawText("Score:" + strconv.Itoa(Sc), Point{w*70/100, h*90/100}, w*4/100)
	DrawLine(Point{0, w*86/100}, Point{h, w*86/100})

	for i := 0; i < 32; i+=2 {
		DrawLine(Point{i*(h/32), w*15/100}, Point{h/32+(i*h/32),w*15/100})
	}

}

func main() {
	InitCanvas(screen)
	Println(h,w)
	SetFrameRate(60)
	SetColor(Color{0,0,0})
	MainLoop(tic)
}



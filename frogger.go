package main

import . "g2d"
import "strconv"

var bestScore = 0
var Score = 0
var h = 512
var w = 448
var Canvas = Size{ w, h }
var x, y = 15, 15
var vx, vy = 5, 5
var frog = LoadImage( "https://i.imgur.com/I2cubui.png" )
var car1 = LoadImage( "https://i.imgur.com/aUYxC98.png" )
var car2 = LoadImage( "https://i.imgur.com/XW22W2b.png" )
var truck = LoadImage( "https://i.imgur.com/8SLloTU.png" )
var background = LoadImage( "https://i.imgur.com/vAZz8KH.png" )

type ent struct {
  x, y   int
  h, w   int
  ix, iy int
  dx, dy int
}

func newent( pos Point, h, w int, ix, iy int, dx, dy int ) *ent {
  b := &ent{ pos.X, pos.Y, h, w, ix, iy, dx, dy }
  return b
}

func ( b *ent ) Move() {
  if b.x > ( w + 200 ) {
    b.x = -100
  }
  if b.x < ( -200 ) {
    b.x = w + 100
  }
  b.x += b.dx
  if (frog.x > b.x && frog.x < b.x + b.w ) && frog.y == b.y {
    frog.x = frog.ix
    frog.y = frog.iy
    Score = 0
  }
}

func Environment() {
  car11.Move()
  car12.Move()
  car21.Move()
  car22.Move()
  truck1.Move()
  truck2.Move()
  DrawImage( background, Point{ 0, 0 } )
  DrawImage( car1, Point{ car11.x, car11.y } )
  DrawImage( car2, Point{ car21.x, car21.y } )
  DrawImage( car1, Point{ car12.x, car12.y } )
  DrawImage( car2, Point{ car22.x, car22.y } )
  DrawImage( truck, Point{ truck1.x, truck1.y } )
  DrawImage( truck, Point{ truck2.x, truck2.y } )
}

var car11 = newent( Point{ 70, h * 16 / 32 }, 32, 64, 70, h * 16 / 32, 10, 0 )
var car12 = newent( Point{ 200, h * 8 / 32 }, 32, 64, 70, h * 8 / 32, 10, 0 )
var car21 = newent( Point{ 50, h * 12 / 32 }, 32, 32, 50, h * 12 / 32, -10, 0 )
var car22 = newent( Point{ 100, h * 22 / 32 }, 32, 32, 50, h * 22 / 32, -10, 0 )
var truck1 = newent( Point{ 40, h * 14 / 32 }, 32, 32, 40, h * 14 / 32, -5, 0 )
var truck2 = newent( Point{ 100, h * 6 / 32 }, 32, 32, 40, h * 6 / 32, -6, 0 )
var frog = newent( Point{ ( w / 2 ) - 8 , h * 26 / 32 }, 32, 32, ( w +32 ) / 2, h * 26 / 32, 0, 0 )

var jump = 32

func Tick() {
  ClearCanvas()
  Environment()
  if KeyPressed( "w" ) || KeyPressed( "ArrowUp" ) {
    frog.y -= jump
  }
  if KeyPressed( "s" ) || KeyPressed( "ArrowDown" ) {
    if frog.y + frog.h + jump <= h {
      frog.y += jump
    }
  }
  if KeyPressed( "a" ) || KeyPressed( "ArrowLeft" ) {
    if 0 <= frog.x-jump {
      frog.x -= jump
    }
  }
  if KeyPressed( "d" ) || KeyPressed( "ArrowRight" ) {
    if frog.x + frog.w + jump <= w {
      frog.x += jump
    }
  }
  if frog.y+frog.h < w * 15 / 100 {
    Score++
    frog.x = frog.ix
    frog.y = frog.iy
    if Score > bestScore {
      bestScore = Score
    }
  }
  DrawImage( frog, Point{ frog.x, frog.y } )
    SetColor( Color{ 255, 255, 255 } )
  DrawText( "Best Score: " + strconv.Itoa( bestScore ), Point{ w * 70 / 100, h * 85 / 100 }, w * 4 / 100 )
  DrawText( "Score: " + strconv.Itoa( Score ), Point{ w * 70 / 100, h * 90 / 100 }, w * 4 / 100 )
  DrawLine( Point{ 0, w * 86 / 100 }, Point{ h, w * 86 / 100 } )
}

func main() {
  InitCanvas( Canvas )
  SetFrameRate( 60 )
  MainLoop( Tick )
}

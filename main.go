package main

import (
	"github.com/fogleman/gg"
)

var (
	x  float64 = 0.1
	y  float64 = 0
	z  float64 = 0
	a  float64 = 14
	b  float64 = 45
	c  float64 = 8 / 3
	dx float64
	dy float64
	dz float64
)

const (
	size      int     = 900
	floatSize float64 = float64(size)
	dt        float64 = 0.001
)

func main() {
	//context 1
	dc1 := gg.NewContext(size, size)
	dc1.SetRGB(0, 0, 0)
	dc1.DrawRectangle(0, 0, floatSize, floatSize)
	dc1.Fill()
	//context 2
	dc2 := gg.NewContext(size, size)
	dc2.SetRGB(0, 0, 0)
	dc2.DrawRectangle(0, 0, floatSize, floatSize)
	dc2.Fill()

	for i := 0; i < 1000000; i++ {
		//=======MATHS=====\\
		dx, dy, dz = (a*(y-x))*dt, (x*(b-z)-y)*dt, (x*y-c*z)*dt
		x, y, z = x+dx, y+dy, z+dz
		//=======VIZUALISATION====\\

		//context 1

		dc1.SetRGB255(int(z*255), int(y*255), 255)
		dc1.DrawCircle((z * floatSize / 70), (y*floatSize/70)+floatSize/2, 0.15)
		dc1.Fill()

		//context 2
		dc2.SetRGB255(int(x*255), int(y*255), 255)
		dc2.DrawCircle((x*(floatSize/20))+floatSize/2, (y*floatSize/20)+floatSize/2, 0.15)
		dc2.Fill()

	}
	dc1.SavePNG("z.png")
	dc2.SavePNG("x.png")

}

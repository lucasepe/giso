package main

import (
	"github.com/lucasepe/giso"
)

func main() {
	red := "#a03c32"
	blue := "#323ca0"
	gray := "#868791"

	iso := giso.NewIsometric(320, 320, 50)

	iso.AddShape(giso.Prism(3, 3, 1), gray)
	iso.AddShape(giso.Pyramid(1, 1, 1).Translate(0, 2, 1), red)
	iso.AddShape(giso.Prism(1, 1, 1).Translate(2, 0, 1), blue)

	iso.Render()
	iso.SavePNG("example_01.png")
}

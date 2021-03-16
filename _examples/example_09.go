package main

import (
	"github.com/lucasepe/giso"
)

func main() {
	red := "#a03c32"
	blue := "#323ca0"

	iso := giso.NewIsometric(320, 320, 20)

	iso.AddShape(giso.Prism(2, 7, 5).Translate(3, 0, 0), red)
	iso.AddShape(giso.Prism(4, 7, 2).Translate(-1, 0, 0), blue)

	iso.Render()
	iso.SavePNG("example_02.png")
}

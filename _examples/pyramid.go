package main

import (
	"github.com/lucasepe/giso"
)

func main() {
	iso := giso.NewIsometric(320, 320, 100)

	iso.AddShape(giso.Pyramid(1, 1, 2), "#535999")

	iso.Render()
	iso.SavePNG("pyramid.png")
}

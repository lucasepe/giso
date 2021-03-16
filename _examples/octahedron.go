package main

import (
	"github.com/lucasepe/giso"
)

func main() {
	iso := giso.NewIsometric(320, 320, 130)

	iso.AddShape(giso.Octahedron(), "#a03c32")

	iso.Render()
	iso.SavePNG("octahedron.png")
}

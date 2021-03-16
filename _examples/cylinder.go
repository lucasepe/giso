package main

import (
	"github.com/lucasepe/giso"
)

func main() {
	iso := giso.NewIsometric(320, 320, 40)

	iso.AddShape(giso.Cylinder(1, 3, 20), "#76d69e")

	iso.Render()
	iso.SavePNG("cylinder.png")
}

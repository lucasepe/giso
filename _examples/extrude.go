package main

import (
	"github.com/lucasepe/giso"
)

func main() {
	iso := giso.NewIsometric(320, 320, 20)

	iso.AddShape(giso.Extrude(giso.Star(2.5, 1.5, 5), 3), "#d4cd0f")

	iso.Render()
	iso.SavePNG("extrude.png")
}

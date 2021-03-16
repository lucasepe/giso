package main

import (
	"github.com/lucasepe/giso"
)

func main() {
	iso := giso.NewIsometric(320, 320, 100)

	iso.AddShape(giso.Prism(1, 1, 1), "#a03c32")

	iso.Render()
	iso.SavePNG("prism.png")
}

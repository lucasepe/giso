package main

import (
	"github.com/lucasepe/giso"
)

func main() {
	iso := giso.NewIsometric(320, 320, 120)

	iso.AddShape(giso.Stairs(5), "#db538e")

	iso.Render()
	iso.SavePNG("stairs.png")
}

package main

import (
	"math"

	. "github.com/lucasepe/giso"
)

func main() {
	iso := NewIsometric(500, 500, 46)

	iso.AddShape(Prism(4, 5, 2).Translate(1, -1, 0), "#2196f3")
	iso.AddShape(Prism(1, 4, 1), "#2196f3")
	iso.AddShape(Prism(1, 3, 1).Translate(-1, 1, 0), "#2196f3")
	iso.AddShape(Stairs(10).Translate(-1, 0, 0), "#2196f3")
	iso.AddShape(Stairs(10).Translate(0, 3, 1).RotateZ(At(0.5, 3.5, 1), -math.Pi/2), "#2196f3")
	iso.AddShape(Prism(2, 4, 1).Translate(3, 0, 2), "#2196f3")
	iso.AddShape(Prism(1, 3, 1).Translate(2, 1, 2), "#2196f3")
	iso.AddShape(Stairs(10).Translate(2, 0, 2).RotateZ(At(2.5, 0.5, 1), -math.Pi/2), "#2196f3")
	iso.AddShape(Pyramid(1, 1, 1).Translate(2, 3, 3).Scale(At(2, 4, 3), 0.5, 0.5, 0.5), "#b4b400")
	iso.AddShape(Pyramid(1, 1, 1).Translate(4, 3, 3).Scale(At(5, 4, 3), 0.5, 0.5, 0.5), "#b400b4")
	iso.AddShape(Pyramid(1, 1, 1).Translate(4, 1, 3).Scale(At(5, 1, 3), 0.5, 0.5, 0.5), "#00b4b4")
	iso.AddShape(Pyramid(1, 1, 1).Translate(2, 1, 3).Scale(At(2, 1, 3), 0.5, 0.5, 0.5), "#28b428")
	iso.AddShape(Prism(1, 1, 0.2).Translate(3, 2, 3), "#323232")
	iso.AddShape(Octahedron().Translate(3, 2, 3.2), "#00b4b4")

	iso.Render()
	iso.SavePNG("castle.png")
}

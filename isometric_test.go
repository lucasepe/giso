package giso

import (
	"math"
	"testing"
)

// https://imgur.com/a/HyKcWhJ
func TestCastle(t *testing.T) {
	iso := NewIsometric(320, 320, 30)

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
	iso.AddShape(Octahedron(At(3, 2, 3.2)), "#00b4b4")

	iso.Render()
	iso.SavePNG("castle.png")
}

func TestTranslate(t *testing.T) {
	iso := NewIsometric(320, 320, 50)

	pr1 := Prism(1, 1, 1)
	pr2 := pr1.Translate(0, 0, 1.1)
	iso.AddShape(pr1, "#2196f3")
	iso.AddShape(pr2, "#f33d21")
	iso.AddShape(pr2.Translate(0, 0, 1.1), "#2196f3")

	iso.Render()
	iso.SavePNG("translate.png")
}

func TestTranslate2(t *testing.T) {
	iso := NewIsometric(320, 320, 40)

	iso.AddShape(Prism(1, 2.4, 1).Translate(0, 0, 1.1), "#2196f3")
	iso.AddShape(Prism(1, 1, 1), "#d160ca")
	iso.AddShape(Prism(1, 1, 1).Translate(0, 1.4, 0), "#f33d21")

	iso.Render()
	iso.SavePNG("ko1.png")

	iso.Clear("#ffffff")
	iso.AddShape(Prism(1, 1, 1), "#d160ca")
	iso.AddShape(Prism(1, 1, 1).Translate(0, 1.4, 0), "#f33d21")
	iso.AddShape(Prism(1, 2.4, 1).Translate(0, 0, 1.1), "#2196f3")

	iso.Render()
	iso.SavePNG("ko2.png")

	iso.Clear("#ffffff")
	iso.AddShape(Prism(1, 1, 1).Translate(0, 1.4, 0), "#f33d21")
	iso.AddShape(Prism(1, 1, 1), "#d160ca")
	iso.AddShape(Prism(1, 2.4, 1).Translate(0, 0, 1.1), "#2196f3")

	iso.Render()
	iso.SavePNG("ok.png")
}

func TestElle(t *testing.T) {
	iso := NewIsometric(320, 320, 20)

	iso.AddShape(Prism(2, 7, 5).Translate(3, 0, 0), "#146b06")
	iso.AddShape(Prism(4, 7, 2).Translate(-1, 0, 0), "#2196f3")

	iso.Render()
	iso.SavePNG("elle.png")
}

func TestExtrude(t *testing.T) {
	iso := NewIsometric(320, 320, 20)

	iso.AddShape(Extrude(Star(3, 1, 5), 4).Translate(2, 2, 0), "#146b06")

	iso.Render()
	iso.SavePNG("extrusion.png")
}

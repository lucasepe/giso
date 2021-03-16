package main

import "github.com/lucasepe/giso"

func main() {

	iso := giso.NewIsometric(320, 320, 50)

	pr1 := giso.Prism(1, 1, 1)
	pr2 := pr1.Translate(0, 0, 1.1)

	iso.AddShape(pr1, "#2196f3")
	iso.AddShape(pr2, "#f33d21")
	iso.AddShape(pr2.Translate(0, 0, 1.1), "#2196f3")

	iso.Render()
	iso.SavePNG("translate.png")
}

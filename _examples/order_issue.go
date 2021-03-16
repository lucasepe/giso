package main

import "github.com/lucasepe/giso"

func main() {
	iso := giso.NewIsometric(320, 320, 40)

	iso.AddShape(giso.Prism(1, 2.4, 1).Translate(0, 0, 1.1), "#2196f3")
	iso.AddShape(giso.Prism(1, 1, 1), "#d160ca")
	iso.AddShape(giso.Prism(1, 1, 1).Translate(0, 1.4, 0), "#f33d21")

	iso.Render()
	iso.SavePNG("order_issue_ko1.png")

	iso.Clear("#ffffff")
	iso.AddShape(giso.Prism(1, 1, 1), "#d160ca")
	iso.AddShape(giso.Prism(1, 1, 1).Translate(0, 1.4, 0), "#f33d21")
	iso.AddShape(giso.Prism(1, 2.4, 1).Translate(0, 0, 1.1), "#2196f3")

	iso.Render()
	iso.SavePNG("order_issue_ko2.png")

	iso.Clear("#ffffff")
	iso.AddShape(giso.Prism(1, 1, 1).Translate(0, 1.4, 0), "#f33d21")
	iso.AddShape(giso.Prism(1, 1, 1), "#d160ca")
	iso.AddShape(giso.Prism(1, 2.4, 1).Translate(0, 0, 1.1), "#2196f3")

	iso.Render()
	iso.SavePNG("order_issue_ok.png")
}

package giso

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/teacat/noire"
)

// https://github.com/FabianTerhorst/Isometric/blob/master/lib/src/main/java/io/fabianterhorst/isometric/Isometric.java

type FilledPath struct {
	*Path
	fillColorHex string
}

type Isometric struct {
	originX, originY float64
	scale            float64
	angle            float64
	lightPosition    *Vector
	lightAngle       *Vector
	colorDifference  float64
	lightColor       color.Color
	transformation   [][]float64
	dc               *gg.Context
	paths            []*FilledPath
}

func NewIsometric(width, height int, scale float64) *Isometric {
	angle := math.Pi / 6
	res := &Isometric{
		scale:           scale,
		angle:           angle,
		colorDifference: 0.20,
		lightColor:      color.White,
		originX:         0.5 * float64(width),
		originY:         0.9 * float64(height),
		dc:              gg.NewContext(width, height),
		paths:           make([]*FilledPath, 0),
		transformation: [][]float64{
			{
				scale * math.Cos(angle),
				scale * math.Sin(angle),
			},
			{
				scale * math.Cos(math.Pi-angle),
				scale * math.Sin(math.Pi-angle),
			},
		},
	}

	res.Clear("#ffffff")
	res.SetLightPosition(&Vector{2, -1, 3})

	return res
}

func (iso *Isometric) Clear(fillColorHex string) {
	iso.dc.SetHexColor(fillColorHex)
	iso.dc.Clear()
}

func (iso *Isometric) SetLightColor(hexColor string) {
	iso.lightColor = ParseHexColor(hexColor)
}

func (iso *Isometric) SetLightPosition(pos *Vector) {
	iso.lightPosition = pos
	iso.lightAngle = iso.lightPosition.Normalize()
}

func (iso *Isometric) AddShape(shape *Shape, hexColor string) {
	paths := shape.orderedPaths()
	for _, el := range paths {
		iso.paths = append(iso.paths, &FilledPath{el, hexColor})
	}
}

func (iso *Isometric) Render() {
	for _, el := range iso.paths {
		fc := iso.computeColor(el.Path, el.fillColorHex)
		tp := iso.translatePath(el.Path)

		iso.drawPath(tp.points, fc)
	}
}

func (iso *Isometric) drawPath(points []*Point, c color.Color) {
	iso.dc.Push()

	iso.dc.MoveTo(points[0].x, points[0].y)
	for i := 1; i < len(points); i++ {
		iso.dc.LineTo(points[i].x, points[i].y)
	}
	iso.dc.ClosePath()

	iso.dc.Push()
	iso.dc.SetColor(c)
	iso.dc.Fill()
	iso.dc.Pop()
}

func (iso *Isometric) translatePath(pat *Path) *Path {
	res := &Path{
		points: make([]*Point, len(pat.points)),
	}

	for i, el := range pat.points {
		res.points[i] = iso.translatePoint(el)
	}
	return res
}

func (iso *Isometric) SavePNG(filename string) error {
	return iso.dc.SavePNG(filename)
}

/**
 * X rides along the angle extended from the origin
 * Y rides perpendicular to this angle (in isometric view: PI - angle)
 * Z affects the y coordinate of the drawn point
 */
func (iso *Isometric) translatePoint(pt *Point) *Point {
	xMap := Point{
		pt.x * iso.transformation[0][0],
		pt.x * iso.transformation[0][1],
		0.0,
	}

	yMap := Point{
		pt.y * iso.transformation[1][0],
		pt.y * iso.transformation[1][1],
		0.0,
	}

	x := iso.originX + xMap.x + yMap.x
	y := iso.originY - xMap.y - yMap.y - (pt.z * iso.scale)

	return &Point{x, y, 0.0}
}

func (iso *Isometric) computeColor(el *Path, source string) color.Color {
	baseColor := noire.NewHex(source)

	v1 := FromTwoPoints(el.points[1], el.points[0])
	v2 := FromTwoPoints(el.points[2], el.points[1])
	normal := CrossProduct(v1, v2)
	normal = normal.Normalize()
	/**
	 * Brightness is between -1 and 1 and is computed based
	 * on the dot product between the light source vector and normal.
	 */
	brightness := DotProduct(normal, iso.lightAngle)
	res := baseColor.Lighten(brightness * iso.colorDifference).Hex()
	return ParseHexColor(res)
}

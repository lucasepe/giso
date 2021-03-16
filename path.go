package giso

import (
	"math"
)

type Path struct {
	points []*Point
}

func NewPath(points []*Point) Path {
	res := Path{
		points: points,
	}

	if res.points == nil {
		res.points = make([]*Point, 0)
	}

	return res
}

// Push append a point to the end of the path.
func (p *Path) Push(pt *Point) *Path {
	p.points = append(p.points, pt)
	return p
}

// Reverse returns a new path with the points
// in reverse order.
func (p *Path) Reverse() *Path {
	res := &Path{points: make([]*Point, len(p.points))}

	for i, j := 0, len(res.points)-1; i < j; i, j = i+1, j-1 {
		res.points[i], res.points[j] = p.points[j], p.points[i]
	}
	return res
}

// Translate translates the path.
func (p *Path) Translate(dx, dy, dz float64) *Path {
	res := &Path{points: make([]*Point, len(p.points))}

	for i, el := range p.points {
		res.points[i] = el.Translate(dx, dy, dz)
	}

	return res
}

// Scale scales the path about the given origin.
func (p *Path) Scale(origin *Point, dx, dy, dz float64) *Path {
	res := &Path{points: make([]*Point, len(p.points))}

	for i, el := range p.points {
		res.points[i] = el.Scale(origin, dx, dy, dz)
	}

	return res
}

// RotateX rotates the point about origin on the X axis.
func (p *Path) RotateX(origin *Point, angle float64) *Path {
	res := &Path{points: make([]*Point, len(p.points))}

	for i, el := range p.points {
		res.points[i] = el.RotateX(origin, angle)
	}

	return res
}

// RotateY rotates the point about origin on the Y axis.
func (p *Path) RotateY(origin *Point, angle float64) *Path {
	res := &Path{points: make([]*Point, len(p.points))}

	for i, el := range p.points {
		res.points[i] = el.RotateY(origin, angle)
	}

	return res
}

// RotateZ rotates the point about origin on the Z axis.
func (p *Path) RotateZ(origin *Point, angle float64) *Path {
	res := &Path{points: make([]*Point, len(p.points))}

	for i, el := range p.points {
		res.points[i] = el.RotateZ(origin, angle)
	}

	return res
}

// Depth returns the estimated depth of
// the path as defined by the average depth
// of its points.
func (p *Path) Depth() float64 {
	if len(p.points) == 0 {
		return 1
	}

	res := 0.0
	for _, el := range p.points {
		res = res + el.Depth()
	}

	return res / float64(len(p.points))
}

// Rectangle returns a rectangle 'path' with
// the bottom-left corner in the origin.
func Rectangle(width, height float64) *Path {
	return &Path{
		points: []*Point{
			{0, 0, 0},
			{width, 0, 0},
			{width, height, 0},
			{0, height, 0},
		},
	}
}

// Circle returns a circle 'path' centered
// at origin with a given radius and number of vertices.
func Circle(radius float64, vertices int) *Path {
	if vertices < 5 {
		vertices = 5
	}

	res := &Path{
		points: make([]*Point, vertices),
	}

	tot := float64(vertices)
	for i := 0; i < vertices; i++ {
		res.points[i] = &Point{
			x: radius * math.Cos(float64(i)*2.0*math.Pi/tot),
			y: radius * math.Sin(float64(i)*2.0*math.Pi/tot),
			z: 0,
		}
	}

	return res
}

// Star returns a star 'path' centered at origin
// with a given outer radius, inner radius, and number of points.
func Star(outerRadius, innerRadius float64, points int) *Path {
	res := &Path{
		points: make([]*Point, 2*points),
	}

	tot := float64(points)

	for i := 0; i < points*2; i++ {
		r := innerRadius
		if i%2 == 0 {
			r = outerRadius
		}

		res.points[i] = &Point{
			x: r * math.Cos(float64(i)*math.Pi/tot),
			y: r * math.Sin(float64(i)*math.Pi/tot),
			z: 0,
		}
	}

	return res
}

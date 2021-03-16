package giso

import (
	"math"
	"sort"
)

type Shape struct {
	paths []*Path
}

func NewShape(paths []*Path) *Shape {
	res := &Shape{paths: paths}
	if res.paths == nil {
		res.paths = make([]*Path, 0)
	}
	return res
}

// Push append a path at the end of the shape.
func (sh *Shape) Push(pat *Path) *Shape {
	sh.paths = append(sh.paths, pat)
	return sh
}

// Translate translates the shape.
func (sh *Shape) Translate(dx, dy, dz float64) *Shape {
	res := &Shape{paths: make([]*Path, len(sh.paths))}

	for i, el := range sh.paths {
		res.paths[i] = el.Translate(dx, dy, dz)
	}

	return res
}

// Scale scales the shape about the given origin.
func (sh *Shape) Scale(origin *Point, dx, dy, dz float64) *Shape {
	res := &Shape{paths: make([]*Path, len(sh.paths))}

	for i, el := range sh.paths {
		res.paths[i] = el.Scale(origin, dx, dy, dz)
	}

	return res
}

// RotateX rotates the shape about origin on the X axis.
func (sh *Shape) RotateX(origin *Point, angle float64) *Shape {
	res := &Shape{paths: make([]*Path, len(sh.paths))}

	for i, el := range sh.paths {
		res.paths[i] = el.RotateX(origin, angle)
	}

	return res
}

// RotateY rotates the shape about origin on the X axis.
func (sh *Shape) RotateY(origin *Point, angle float64) *Shape {
	res := &Shape{paths: make([]*Path, len(sh.paths))}

	for i, el := range sh.paths {
		res.paths[i] = el.RotateY(origin, angle)
	}

	return res
}

// RotateZ rotates the shape about origin on the X axis.
func (sh *Shape) RotateZ(origin *Point, angle float64) *Shape {
	res := &Shape{paths: make([]*Path, len(sh.paths))}

	for i, el := range sh.paths {
		res.paths[i] = el.RotateZ(origin, angle)
	}

	return res
}

func (sh *Shape) orderedPaths() []*Path {
	sort.Slice(sh.paths, func(i, j int) bool {
		return sh.paths[i].Depth() > sh.paths[j].Depth()
	})

	return sh.paths
}

// Extrude creates a 3D object by raising a 2D path along the z-axis.
func Extrude(p *Path, height float64) *Shape {
	top := p.Translate(0, 0, height)

	shape := &Shape{
		paths: make([]*Path, 0),
	}
	/* Push the top and bottom faces, top face must be oriented correctly */
	shape.Push(p.Reverse())
	shape.Push(top)

	topPathLen := len(top.points)
	patLen := len(p.points)

	for i := 0; i < patLen; i++ {
		shape.Push(&Path{
			points: []*Point{
				top.points[i],
				p.points[i],
				p.points[(i+1)%patLen],
				top.points[(i+1)%topPathLen],
			},
		})
	}

	return shape
}

func Prism(dx, dy, dz float64) *Shape {
	shape := &Shape{
		paths: make([]*Path, 6),
	}

	// Squares parallel to the x-axis
	face1 := &Path{
		points: []*Point{
			{0, 0, 0},
			{dx, 0, 0},
			{dx, 0, dz},
			{0, 0, dz},
		},
	}

	// Push this face and its opposite
	shape.paths[0] = face1
	shape.paths[1] = face1.Reverse().Translate(0, dy, 0)

	// Square parallel to the y-axis
	face2 := &Path{
		points: []*Point{
			{0, 0, 0},
			{0, 0, dz},
			{0, dy, dz},
			{0, dy, 0},
		},
	}
	shape.paths[2] = face2
	shape.paths[3] = face2.Reverse().Translate(dx, 0, 0)

	// Square parallel to the xy-plane
	face3 := &Path{
		points: []*Point{
			{0, 0, 0},
			{dx, 0, 0},
			{dx, dy, 0},
			{0, dy, 0},
		},
	}
	/* This surface is oriented backwards, so we need to reverse the points */
	shape.paths[4] = face3.Reverse()
	shape.paths[5] = face3.Translate(0, 0, dz)

	return shape
}

func Stairs(steps int) *Shape {
	paths := make([]*Path, steps*2+2)

	zigZag := &Path{}
	points := make([]*Point, steps*2+2)
	points[0] = &Point{0, 0, 0}

	count := 1
	for i := 0; i < steps; i++ {
		stepCorner := &Point{0, float64(i) / float64(steps), float64(i+1) / float64(steps)}
		paths[count-1] = &Path{
			points: []*Point{
				stepCorner,
				stepCorner.Translate(0, 0, -1/float64(steps)),
				stepCorner.Translate(1, 0, -1/float64(steps)),
				stepCorner.Translate(1, 0, 0),
			},
		}
		points[count] = stepCorner
		count = count + 1

		paths[count-1] = &Path{
			points: []*Point{
				stepCorner,
				stepCorner.Translate(1, 0, 0),
				stepCorner.Translate(1, 1/float64(steps), 0),
				stepCorner.Translate(0, 1/float64(steps), 0),
			},
		}

		points[count] = stepCorner.Translate(0, 1/float64(steps), 0)
		count = count + 1
	}

	points[count] = &Point{0, 1, 0}
	zigZag.points = points

	paths[count-1] = zigZag
	count = count + 1

	tmp := zigZag.Reverse()
	paths[count-1] = tmp.Translate(1, 0, 0)

	return &Shape{paths}
}

func Pyramid(dx, dy, dz float64) *Shape {
	// Path parallel to the x-axis
	face1 := &Path{
		points: []*Point{
			{0, 0, 0},
			{dx, 0, 0},
			{dx / 2, dy / 2, dz},
		},
	}
	// Path parallel to the y-axis
	face2 := &Path{
		points: []*Point{
			{0, 0, 0},
			{dx / 2, dy / 2, dz},
			{0, dy, 0},
		},
	}

	centerOfRot := &Point{0.5 * dx, 0.5 * dy, 0}

	return &Shape{
		paths: []*Path{
			face1,
			face1.RotateZ(centerOfRot, math.Pi),
			face2,
			face2.RotateZ(centerOfRot, math.Pi),
		},
	}
}

func Cylinder(radius, height float64, vertices int) *Shape {
	circle := Circle(radius, vertices)
	return Extrude(circle, height)
}

func Octahedron() *Shape {
	center := &Point{0.5, 0.5, 0.5}

	upperTriangle := &Path{
		points: []*Point{
			{0.0, 0.0, 0.5},
			{0.5, 0.5, 1.0},
			{0.0, 1.0, 0.5},
		},
	}

	lowerTriangle := &Path{
		points: []*Point{
			{0.0, 0.0, 0.5},
			{0.0, 1.0, 0.5},
			{0.5, 0.5, 0.0},
		},
	}

	paths := make([]*Path, 8)
	count := 0
	for i := 0; i < 4; i++ {
		paths[count] = upperTriangle.RotateZ(center, float64(i)*math.Pi/2.0)
		count = count + 1
		paths[count] = lowerTriangle.RotateZ(center, float64(i)*math.Pi/2.0)
		count = count + 1
	}

	res := &Shape{paths: paths}
	return res.Scale(center, math.Sqrt(2)/2.0, math.Sqrt(2)/2.0, 1)
}

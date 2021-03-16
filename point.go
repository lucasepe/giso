package giso

import (
	"fmt"
	"math"
)

// Point is a position in a 3D space.
type Point struct {
	x float64
	y float64
	z float64
}

// At returns a point.
func At(x, y, z float64) *Point {
	return &Point{x, y, z}
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) Z() float64 {
	return p.z
}

// Translate translates the point by a given dx, dy, and dz.
func (p *Point) Translate(dx, dy, dz float64) *Point {
	return &Point{
		x: p.x + dx,
		y: p.y + dy,
		z: p.z + dz,
	}
}

// Scale scales the point about a given origin.
func (p *Point) Scale(origin *Point, dx, dy, dz float64) *Point {
	t := p.Translate(-origin.x, -origin.y, -origin.z)

	t.x = t.x * dx
	t.y = t.y * dy
	t.z = t.z * dz

	return t.Translate(origin.x, origin.y, origin.z)
}

// RotateX rotates the point about origin on the X axis.
func (p *Point) RotateX(origin *Point, angle float64) *Point {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	t := p.Translate(-origin.x, -origin.y, -origin.z)

	z := t.z*cos - t.y*sin
	y := t.z*sin + t.y*cos
	t.z, t.y = z, y

	return t.Translate(origin.x, origin.y, origin.z)
}

// RotateY rotates the point about origin on the Y axis.
func (p *Point) RotateY(origin *Point, angle float64) *Point {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	t := p.Translate(-origin.x, -origin.y, -origin.z)

	x := t.x*cos - t.z*sin
	z := t.x*sin + t.z*cos
	t.x, t.z = x, z

	return t.Translate(origin.x, origin.y, origin.z)
}

// RotateZ rotates the point about origin on the Z axis.
func (p *Point) RotateZ(origin *Point, angle float64) *Point {
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	t := p.Translate(-origin.x, -origin.y, -origin.z)

	x := t.x*cos - t.y*sin
	y := t.x*sin + t.y*cos
	t.x, t.y = x, y

	return t.Translate(origin.x, origin.y, origin.z)
}

// Depth computes the depth of a point in the isometric plane.
func (p *Point) Depth() float64 {
	// z is weighted slightly to accomodate |_ arrangements
	return p.x + p.y - 2*p.z
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v, %v, %v)", p.x, p.y, p.z)
}

/* NOT USED TBD remove it!
// Distance returns the distance between two points.
func Distance(p1, p2 *Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	dz := p2.z - p1.z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
*/

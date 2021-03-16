package giso

import (
	"fmt"
	"math"
)

// Vector is a quantity that has
// both magnitude and direction.
type Vector struct {
	i float64
	j float64
	k float64
}

// NewVector returns a vector with the specified coords.
func NewVector(i, j, k float64) *Vector {
	return &Vector{i, j, k}
}

// Magnitude returns the length of the vector.
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.i*v.i + v.j*v.j + v.k*v.k)
}

// Normalize returns a vector that has the same
// direction of the given vector but its length is one.
func (v *Vector) Normalize() *Vector {
	mag := v.Magnitude()
	// If the magnitude is 0 then return the zero
	// vector instead of dividing by 0
	if mag == 0 {
		return &Vector{0, 0, 0}
	}

	return &Vector{v.i / mag, v.j / mag, v.k / mag}
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%v, %v, %v>", v.i, v.j, v.k)
}

// FromTwoPoints returns a vector from two points.
func FromTwoPoints(p1, p2 *Point) *Vector {
	return &Vector{
		p2.x - p1.x,
		p2.y - p1.y,
		p2.z - p1.z,
	}
}

// CrossProduct returns a new vector that is
// perpendicular to both v1 and v2.
func CrossProduct(v1, v2 *Vector) *Vector {
	i := v1.j*v2.k - v2.j*v1.k
	j := -1 * (v1.i*v2.k - v2.i*v1.k)
	k := v1.i*v2.j - v2.i*v1.j
	return &Vector{i, j, k}
}

// DotProduct returns a vector that is the dot
// product of v1 and v2.
func DotProduct(v1, v2 *Vector) float64 {
	return v1.i*v2.i + v1.j*v2.j + v1.k*v2.k
}

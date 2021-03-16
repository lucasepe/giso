package giso

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	tolerance = .0001
)

var (
	cmpOpts = []cmp.Option{
		cmp.AllowUnexported(Vector{}),
		cmp.Comparer(func(x, y float64) bool {
			diff := math.Abs(x - y)
			mean := math.Abs(x+y) / 2.0
			if math.IsNaN(diff / mean) {
				return true
			}
			return (diff / mean) < tolerance
		}),
	}
)

func TestVectorFromTwoPoint(t *testing.T) {
	cases := []struct {
		first  *Point
		second *Point
		want   *Vector
	}{
		{
			first:  At(2, 2, 1),
			second: At(6, 3, 2),
			want:   NewVector(4, 1, 1),
		},
		{
			first:  At(2, -7, 0),
			second: At(1, -3, -5),
			want:   NewVector(-1, 4, -5),
		},
		{
			first:  At(1, -3, -5),
			second: At(2, -7, 0),
			want:   NewVector(1, -4, 5),
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s %s", tt.first, tt.second), func(t *testing.T) {
			got := FromTwoPoints(tt.first, tt.second)
			if !cmp.Equal(got, tt.want, cmpOpts...) {
				t.Fatalf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestVectorMagnitude(t *testing.T) {
	cases := []struct {
		input *Vector
		want  float64
	}{
		{NewVector(4, 1, 1), 4.2426},
		{NewVector(-1, 4, -5), 6.4807},
		{NewVector(3, -5, 10), 11.5758},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			got := tt.input.Magnitude()
			if !cmp.Equal(got, tt.want, cmpOpts...) {
				t.Fatalf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestVectorNormalize(t *testing.T) {
	cases := []struct {
		input *Vector
		want  *Vector
	}{
		{NewVector(4, 1, 1), NewVector(0.9428, 0.2357, 0.2357)},
		{NewVector(-1, 4, -5), NewVector(-0.1543, 0.6172, -0.7715)},
		{NewVector(3, -5, 10), NewVector(0.25916, -0.43193, 0.86386)},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s", tt.input), func(t *testing.T) {
			got := tt.input.Normalize()
			if !cmp.Equal(got, tt.want, cmpOpts...) {
				t.Fatalf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestVectorCrossProduct(t *testing.T) {
	cases := []struct {
		v1, v2 *Vector
		want   *Vector
	}{
		{
			v1: NewVector(4, 1, 1), v2: NewVector(-1, 9, -3),
			want: NewVector(-12, 11, 37),
		},
		{
			v1: NewVector(3, -5, 10), v2: NewVector(1, 8, -2),
			want: NewVector(-70, 16, 29),
		},
		{
			v1: NewVector(14, 11, -1), v2: NewVector(-10, 9, -2),
			want: NewVector(-13, 38, 236),
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s %s", tt.v1, tt.v2), func(t *testing.T) {
			got := CrossProduct(tt.v1, tt.v2)
			if !cmp.Equal(got, tt.want, cmpOpts...) {
				t.Fatalf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

func TestVectorDotProduct(t *testing.T) {
	cases := []struct {
		v1, v2 *Vector
		want   float64
	}{
		{
			v1: NewVector(4, 1, 1), v2: NewVector(-1, 9, -3),
			want: 2,
		},
		{
			v1: NewVector(3, -5, 10), v2: NewVector(1, 8, -2),
			want: -57,
		},
		{
			v1: NewVector(14, 11, -1), v2: NewVector(-10, 9, -2),
			want: -39,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s %s", tt.v1, tt.v2), func(t *testing.T) {
			got := DotProduct(tt.v1, tt.v2)
			if !cmp.Equal(got, tt.want, cmpOpts...) {
				t.Fatalf("got [%v] want [%v]", got, tt.want)
			}
		})
	}
}

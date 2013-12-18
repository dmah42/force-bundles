package fblib

import (
	"math"
)

const (
	EPSILON = 0.0001
)

type Point struct {
	X, Y float64
}

func (p Point) Add(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y}
}

func (p Point) Sub(p1 Point) Vector {
	return Vector{p.X - p1.X, p.Y - p1.Y}
}

type Vector struct {
	X, Y float64
}

func (v Vector) Add(v1 Vector) Vector {
	return Vector{v.X + v1.X, v.Y + v1.Y}
}

func (v Vector) Scale(f float64) Vector {
	return Vector{v.X * f, v.Y * f}
}

func (v Vector) Dot(v1 Vector) float64 {
	return v.X*v1.X + v.Y*v1.Y
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

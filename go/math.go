package main

import (
	"math"
)

const (
	EPSILON = 0.0001
)

type Point struct {
	x, y	float64
}

func (p Point) Add(v Vector) Point {
	return Point {p.x + v.x, p.y + v.y}
}

func (p Point) Sub(p1 Point) Vector {
	return Vector {p.x - p1.x, p.y - p1.y}
}

type Vector struct {
	x, y	float64
}

func (v Vector) Add(v1 Vector) Vector {
	return Vector {v.x + v1.x, v.y + v1.y}
}

func (v Vector) Scale(f float64) Vector {
	return Vector {v.x * f, v.y * f}
}

func (v Vector) Dot(v1 Vector) float64 {
	return v.x * v1.x + v.y * v1.y
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

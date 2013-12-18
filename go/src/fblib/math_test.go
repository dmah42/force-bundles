package fblib

import (
	"math"
	"testing"
)

func TestPointAdd(t *testing.T) {
	cases := []struct {
		p	Point
		v	Vector
		r	Point
	}{
		{Point{0, 0}, Vector{1, 1.4}, Point{1, 1.4}},
		{Point{-1, -3}, Vector{1, 5}, Point{0, 2}},
	}

	for _, tt := range cases {
		r := tt.p.Add(tt.v)
		if r != tt.r {
			t.Errorf("want %v, got %v\n", tt.r, r)
		}
	}
}

func TestPointSub(t *testing.T) {
	cases := []struct {
		p0,p1	Point
		r	Vector
	}{
		{Point{0, 0}, Point{1, 1.4}, Vector{-1, -1.4}},
		{Point{-1, -3}, Point{1, 5}, Vector{-2, -8}},
	}

	for _, tt := range cases {
		r := tt.p0.Sub(tt.p1)
		if r != tt.r {
			t.Errorf("want %v, got %v\n", tt.r, r)
		}
	}
}

func TestVectorAdd(t *testing.T) {
	cases := []struct {
		v0,v1,r	Vector
	}{
		{Vector{0, 0}, Vector{1, 1.4}, Vector{1, 1.4}},
		{Vector{-1, -3}, Vector{1, 5}, Vector{0, 2}},
	}

	for _, tt := range cases {
		r := tt.v0.Add(tt.v1)
		if r != tt.r {
			t.Errorf("want %v, got %v\n", tt.r, r)
		}
	}
}

func TestVectorScale(t *testing.T) {
	cases := []struct {
		v	Vector
		f	float64
		r	Vector
	}{
		{Vector{0, 0}, 1.4, Vector{0, 0}},
		{Vector{1, 1}, 1, Vector{1, 1}},
		{Vector{2, 3}, -2, Vector{-4, -6}},
	}

	for _, tt := range cases {
		r := tt.v.Scale(tt.f)
		if r != tt.r {
			t.Errorf("want %v, got %v\n", tt.r, r)
		}
	}
}

func TestVectorDot(t *testing.T) {
	cases := []struct {
		v0,v1	Vector
		r	float64
	}{
		{Vector{0, 0}, Vector{1, 1}, 0},
		{Vector{1, 1}, Vector{-1, 2}, 1},
	}

	for _, tt := range cases {
		r := tt.v0.Dot(tt.v1)
		if r != tt.r {
			t.Errorf("want %v, got %v\n", tt.r, r)
		}
	}
}

func TestVectorLength(t *testing.T) {
	cases := []struct {
		v	Vector
		r	float64
	}{
		{Vector{0, 0}, 0},
		{Vector{1, 1}, math.Sqrt(2)},
		{Vector{-3, 4}, 5},
	}

	for _, tt := range cases {
		r := tt.v.Length()
		if r != tt.r {
			t.Errorf("want %v, got %v\n", tt.r, r)
		}
	}
}

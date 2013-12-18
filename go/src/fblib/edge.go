package fblib

import (
	"fmt"
	"math"
)

const (
	K = 0.01
)

type Edge struct {
	forces     []Vector
	velocities []Vector
	vertices   []Point
}

func NewEdge(p0, p1 Point) *Edge {
	return &Edge{forces: []Vector{}, velocities: []Vector{}, vertices: []Point{p0, p1},}
}

func (e *Edge) compatibility(q Edge) float64 {
	delta_p := e.vertices[len(e.vertices)-1].Sub(e.vertices[0])
	delta_q := q.vertices[len(q.vertices)-1].Sub(q.vertices[0])

	len_p := delta_p.Length()
	len_q := delta_q.Length()

	// angle
	Ca := math.Abs(delta_p.Dot(delta_q) / (len_p * len_q))

	// scale
	len_avg := (len_p + len_q) / 2.0
	Cs := 2.0 / (len_avg*math.Min(len_p, len_q) + math.Max(len_p, len_q)/len_avg)

	// position
	mid_p := e.vertices[len(e.vertices)/2]
	mid_q := q.vertices[len(q.vertices)/2]
	Cp := len_avg / (len_avg + mid_p.Sub(mid_q).Length())

	// visibility
	// TODO
	Cv := 1.0

	return Ca * Cs * Cp * Cv
}

func (e *Edge) Subdivide(segments int) {
	delta := e.vertices[len(e.vertices)-1].Sub(e.vertices[0])
	subdelta := delta.Scale(1.0 / float64(segments))

	newVertices := make([]Point, segments+1)
	newVertices[segments] = e.vertices[len(e.vertices)-1]
	for i := 0; i < segments; i++ {
		newVertices[i] = e.vertices[0].Add(subdelta.Scale(float64(i)))
	}
	e.vertices = newVertices
	e.forces = make([]Vector, len(e.vertices))
	e.velocities = make([]Vector, len(e.vertices))

	if len(e.vertices) != len(e.forces) || len(e.vertices) != len(e.velocities) {
		fmt.Println("WTF0")
	}
}

func (e *Edge) ClearForces() {
	for i, _ := range e.forces {
		e.forces[i] = Vector{0, 0}
	}
}

func (e *Edge) AddSpringForces() {
	if len(e.vertices) != len(e.forces) || len(e.vertices) != len(e.velocities) {
		fmt.Println("WTF1")
	}
	for i := 1; i < len(e.vertices)-1; i++ {
		// spring forces
		delta0 := e.vertices[i-1].Sub(e.vertices[i])
		delta1 := e.vertices[i].Sub(e.vertices[i+1])

		// TODO: shouldn't this be the difference from the original length?
		delta0_len := delta0.Length()
		delta1_len := delta1.Length()

		delta0_dir := delta0.Scale(1.0 / delta0_len)
		delta1_dir := delta1.Scale(1.0 / delta1_len)

		Fs0 := delta0_dir.Scale(K * delta0_len)
		Fs1 := delta1_dir.Scale(K * delta1_len)

		e.forces[i] = e.forces[i].Add(Fs0).Add(Fs1)
	}
}

func (e *Edge) AddElectrostaticForces(q Edge) {
	if len(e.vertices) != len(e.forces) || len(e.vertices) != len(e.velocities) {
		fmt.Println("WTF2")
	}
	compat := e.compatibility(q)
	for i := 1; i < len(e.vertices)-1; i++ {
		// electrostatic forces
		delta_e := e.vertices[i].Sub(q.vertices[i])
		delta_e_len := delta_e.Length()
		delta_e_dir := delta_e.Scale(1.0 / delta_e_len)
		Fe := delta_e_dir.Scale(1.0 / delta_e_len)

		e.forces[i] = e.forces[i].Add(Fe.Scale(compat))
	}
}

func (e *Edge) UpdatePositions(dt float64) bool {
	if len(e.vertices) != len(e.forces) || len(e.vertices) != len(e.velocities) {
		fmt.Println("WTF3")
	}
	moved := false
	for i, _ := range e.vertices {
		// assume mass == 1
		// Euler integration (blech)
		e.velocities[i] = e.velocities[i].Add(e.forces[i].Scale(dt))
		delta_p := e.velocities[i].Scale(dt)
		e.vertices[i] = e.vertices[i].Add(delta_p)

		if delta_p.Length() > EPSILON {
			moved = true
		}
	}
	return moved
}

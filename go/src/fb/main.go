package main

import (
	"fblib"
	"fmt"
)

const (
	DT           = 1.0 / 60
	NUM_SEGMENTS = 4
)

func createGraph() (*fblib.Graph, error) {
	var N int
	fmt.Scan(&N)
	fmt.Printf("Reading %d edges\n", N)

	g := new(fblib.Graph)
	for i := 0; i < N; i++ {
		var a0, z0, a1, z1 float64
		if _, err := fmt.Scanf("%f %f %f %f", &a0, &z0, &a1, &z1); err != nil {
			return nil, err
		}
		g.Add(fblib.NewEdge(fblib.Point{a0, z0}, fblib.Point{a1, z1}))
	}
	fmt.Printf("Create\n  %+v\n", g)
	return g, nil
}

func main() {
	g, err := createGraph()
	if err != nil {
		fmt.Printf("ERROR: %q\n", err)
		return
	}
	g.Subdivide(NUM_SEGMENTS)
	fmt.Printf("Sub\n  %+v\n", g)

	done := false
	for !done {
		done = g.Step(DT)
	}

	fmt.Printf("Done\n  %+v\n", g)
}

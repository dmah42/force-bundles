package main

type Graph struct {
	edges	[]Edge
}

func (g *Graph) Add(e Edge) {
	g.edges = append(g.edges, e)
}

func (g *Graph) Subdivide(segments int) {
	for i, _ := range g.edges {
		g.edges[i].Subdivide(segments)
	}
}

func (g *Graph) Step(dt float64) bool {
	for i, _ := range g.edges {
		g.edges[i].ClearForces()
		g.edges[i].AddSpringForces()
		for j, q := range g.edges {
			if i == j {
				continue
			}

			g.edges[i].AddElectrostaticForces(q)
		}
	}

	done := true
	for i, _ := range g.edges {
		if g.edges[i].UpdatePositions(dt) {
			done = false
		}
	}
	return done
}


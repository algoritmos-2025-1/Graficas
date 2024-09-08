package algorithms

import (
	"github.com/algoritmos-2025-1/Graficas/internal/collection"
	"github.com/algoritmos-2025-1/Graficas/internal/graph"
)

// Algoritmo de búsqueda en gráficas
func GraphSearch(G graph.Graph, r int, C collection.Collection[int]) (*[]int, *[]int) {
	p := make([]int, G.Order())
	for i, _ := range p {
		p[i] = -1
	}

	t := make([]int, G.Order())

	i := 0	
	i = i + 1
	G.Colour(r)
	C.Add(r)
	p[r] = -1
	t[r] = i
	for !C.Empty() {
		x := C.Remove()
		for _, y := range *G.Neighbours(x) {
			if G.Uncoloured(y) {
				i = i + 1
				G.Colour(y)
				C.Add(y)
				p[y] = x
				t[y] = i
			}
		}
	}
	return &p, &t
}
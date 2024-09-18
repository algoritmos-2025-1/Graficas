package algorithms

import (
	"fmt"
	"github.com/algoritmos-2025-1/Graficas/internal/collection"
	"github.com/algoritmos-2025-1/Graficas/pkg/graph"
)

// Algoritmo de búsqueda en gráficas
func GraphSearch(G *graph.StaticGraph, r int, C collection.Collection[int]) (*[]int, *[]int) {
	p := make([]int, G.Order())
	for i, _ := range p {
		p[i] = -1
	}
	fmt.Println(G.Order())

	t := make([]int, G.Order())
	colors := make([]bool, G.Order())
	i := 0	
	i = i + 1
	colors[r] = true;
	C.Add(r)
	p[r] = -1
	t[r] = i
	for !C.Empty() {
		x := C.Remove()
		for _, y := range G.NeighboursList(x) {
			if !colors[y] {
				i = i + 1
				colors[y] = true;
				C.Add(y)
				p[y] = x
				t[y] = i
			}
		}
	}
	return &p, &t
}
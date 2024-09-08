package graph

import "fmt"

// Define una gráfica por listas de adyacencias
type AdjacencyListGraph struct {
	adjacencyList [][]int
	colours       []bool
}

// Crea una gráfica por listas de adyacencias
func CreateAdjacencyListGraph(n int) Graph {
	adjacencyList := make([][]int, n)
	colours := make([]bool, n)
	for i := range adjacencyList {
		adjacencyList[i] = make([]int, 0)
	}
	return &AdjacencyListGraph{
		adjacencyList: adjacencyList,
		colours:       colours,
	}
}

// Añade una arista
func (g *AdjacencyListGraph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

// Colorea un vértice
func (g *AdjacencyListGraph) Colour(v int) {
	g.colours[v] = true
}

// Revisa si un vértice no está coloreado
func (g *AdjacencyListGraph) Uncoloured(v int) bool {
	return !g.colours[v]
}

// Regresa la lista de vecions de un vértice
func (g *AdjacencyListGraph) Neighbours(v int) *[]int {
	return &g.adjacencyList[v]
}

// Regresa el orden de la gráfica
func (g *AdjacencyListGraph) Order() int {
	return len(g.adjacencyList)
}

// Regresa el tamaño de la gráfica
func (g *AdjacencyListGraph) Size() int {
	size := 0
	for _, l := range g.adjacencyList {
		size += len(l)
	}
	return size / 2
}

// Limpia los colores
func (g *AdjacencyListGraph) ResetColours() {
	clear(g.colours)
}

//Imprime una representación de la lista de adyacencias
func (g *AdjacencyListGraph) Print() {
	for x, y := range g.adjacencyList{
		fmt.Println(x, ":", y)
	} 
}
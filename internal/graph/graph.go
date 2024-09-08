package graph

type Graph interface {

	// Añade una arista
	AddEdge(u, v int)

	// Colorea un vértice
	Colour(v int)

	// Revisa si un vértice no está coloreado
	Uncoloured(v int) bool

	// Regresa la lista de vecions de un vértice
	Neighbours(v int) *[]int

	// Limpia los colores
	ResetColours()

	// Regresa el orden de la gráfica
	Order() int

	// Regresa el tamaño de la gráfica
	Size() int
}
package main

import (
	"fmt"
	"github.com/algoritmos-2025-1/Graficas/pkg/graph"
	 "github.com/algoritmos-2025-1/Graficas/internal/collection"
	 "github.com/algoritmos-2025-1/Graficas/internal/stack"
	 "github.com/algoritmos-2025-1/Graficas/pkg/algorithms"
)

func main() {
	lista:= [][]byte{
		{0,1},
		{1,0},
	}
	g,_ := graph.NewGraphFromMatrix(lista)
	s:= stack.CreateSliceStack[int]()
	s.Push(0)
	c:=collection.Collection[int](s)
	p,t:= algorithms.GraphSearch(g,0,c)
	fmt.Println(p,t)
}

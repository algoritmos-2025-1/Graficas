package main

import (
	"fmt"
	"github.com/algoritmos-2025-1/Graficas/pkg/graph"
	 //"github.com/algoritmos-2025-1/Graficas/internal/collection"
	 //"github.com/algoritmos-2025-1/Graficas/internal/stack"
	 "github.com/algoritmos-2025-1/Graficas/pkg/algorithms"
)
func Ternary(condition bool, trueValue, falseValue int) int {
    if condition {
        return trueValue
    }
    return falseValue
}



func main() {
	//Creamos matriz de adyacencia
	matrizAdy:= [][]byte{
		{0,1},
		{1,0},
	}
	//Creamos una gráfica a partir de la matriz
	g,_ := graph.NewGraphFromMatrix(matrizAdy)
	K,S,_ := algorithms.GetKS(g)
	fmt.Println(K,S)
}

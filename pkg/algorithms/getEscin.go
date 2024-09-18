package algorithms
import (
	"sort"
	"github.com/algoritmos-2025-1/Graficas/pkg/graph"
	 //"github.com/algoritmos-2025-1/Graficas/internal/collection"
	 //"github.com/algoritmos-2025-1/Graficas/internal/stack"
	 //"github.com/algoritmos-2025-1/Graficas/pkg/algorithms"
)
type Pair struct {
    Value int
    Index int
}


type ByValue []Pair

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value > a[j].Value }

func getEscin(g *graph.StaticGraph)([]int, []int, string){
	//Obtenemos una sucesión de dgrados de la gráfica
	grados := g.DegreeSequence()
	//Para poder ordenar los grados en forma descendiente guardamos el indice para saber cuál era originalmente
	pairs := make([]Pair, len(grados))
	for i,value := range grados {
		pairs[i] = Pair{Value: value, Index: i}
	}
	sort.Sort(ByValue(pairs))
	//Buscaremos separar la gráfica en el clan y el conjunto independiente
	var clan []int
	var indepen []int
	for r,_ := range pairs {
		izq:=0
		drc := r*(r+1)
		//Haremos la suma izquierda y derecha de la desigualdad, d.Value es el grado del nodo i
		for i,d := range pairs {
			if(i <= r){
				izq = izq + d.Value
			}else{
				drc = drc + min(r,d.Value)
			}
		}
		//Si son iguales, ya encontramos en qué nodo separamos el K y S como conjuntos, entonces agregamos los nodos al respectibo conjunto
		if(izq==drc){
			for i,d := range pairs {
				if(i <= r){
					clan = append(clan, d.Index)
				}else{
					indepen = append(indepen, d.Index)
				}
			}
			return clan,indepen,""
		}
	}
	return nil,nil, "No se da la igualdad."
}

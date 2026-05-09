package tree

type vertex struct {
	data string
	adjacentVertex []*vertex
	lookup map[*vertex]bool
}
type vertexMap map[*vertex]bool

func (v *vertex) dfs() {
	v.innerDfs(v, new(make(map[string]bool)))
}

func createVertex(data string) *vertex {
	return &vertex{
		data: data,
		lookup: make(map[*vertex]bool),
	}
}

func (v *vertex) addToAdjacentList(newVertex *vertex) {
	if v.lookup[newVertex] {
		return
	}
	v.adjacentVertex = append(v.adjacentVertex, newVertex)
	v.lookup[newVertex] = true

	newVertex.addToAdjacentList(v)
}

func (v *vertex) innerDfs(currV *vertex, visited *map[string]bool)  {
	(*visited)[currV.data] = true

	for _, adj := range v.adjacentVertex {
		if _,ok := (*visited)[adj.data]; !ok {
			adj.innerDfs(adj,visited)
		}
	}
}
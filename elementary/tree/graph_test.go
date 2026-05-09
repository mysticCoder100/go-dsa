package tree

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createGraph(vertices *[]*vertex) (count int) {
	path,pathEr := os.Getwd()
	nameList := make(map[string]*vertex)

	if pathEr != nil {
		panic(pathEr)
	}

	fullPath := filepath.Join(path,"list.txt")

	file, fileError := os.Open(fullPath)

	if fileError != nil {
		panic(fileError)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count = 0
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ":")

		currentVertex := getVertex(data[0], &nameList, vertices)

		adjacentList := strings.Split(data[1], ",")

		for _,d := range adjacentList {
			childVertex := getVertex(d, &nameList, vertices)
			currentVertex.addToAdjacentList(childVertex)
		}
		count++
	}
	return
}

func getVertex(data string, nameList *map[string]*vertex, vertices *[]*vertex) *vertex  {
	data = strings.Trim(data," ")

	if d, ok := (*nameList)[data]; ok {
		return d
	}

	newVertex := createVertex(data)
	(*nameList)[data] = newVertex
	*vertices = append(*vertices, newVertex)
	return newVertex
}

func TestInsertion(t *testing.T) {
	var allVertices []*vertex
	count := createGraph(&allVertices)

	if len(allVertices) != count {
		t.Errorf("Incorrect number of vertices, expected %d, got %d", count, len(allVertices))
	}
}

func TestDfs(t *testing.T) {
	var allVertices []*vertex
	createGraph(&allVertices)

	allVertices[0].dfs()
}



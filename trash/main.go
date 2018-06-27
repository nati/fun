package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

const (
	retry = 1000
)

type vertex struct {
	index int
	edges []*vertex
}

func (v *vertex) String() string {
	return fmt.Sprintf("[%d]", v.index)
}

type Graph struct {
	vertex []*vertex
}

//print dots file
func (g *Graph) String() string {
	var buf bytes.Buffer
	buf.WriteString("graph{\n")
	drawed := map[int]bool{}
	for _, v := range g.vertex {
		for _, e := range v.edges {
			i := v.index
			j := e.index
			if j > i {
				t := i
				i = j
				j = t
			}
			h := i + j*len(g.vertex)
			if !drawed[h] {
				buf.WriteString(fmt.Sprintf("\t%d -- %d;\n", v.index, e.index))
				drawed[h] = true
			}
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func (g *Graph) Visit(i int, visited map[int]bool, accept func(v *vertex) bool) {
	if i > len(g.vertex) {
		panic("no vertex")
	}
	v := g.vertex[i]
	if visited[i] {
		return
	}
	visited[i] = true
	if accept(v) {
		for _, e := range v.edges {
			g.Visit(e.index, visited, accept)
		}
	}
}

func (g *Graph) AddEdge(left, right int) bool {
	if left == right {
		return false
	}
	if left >= len(g.vertex) || right >= len(g.vertex) {
		return false
	}
	l := g.vertex[left]
	r := g.vertex[right]
	if l.Connected(right) {
		return false
	}
	l.edges = append(l.edges, r)
	r.edges = append(r.edges, l)
	return true
}

func (v *vertex) Connected(to int) bool {
	for _, e := range v.edges {
		if e.index == to {
			return true
		}
	}
	return false
}

func makeGraph(seed int64, vertexCount, edgeCount int) (g *Graph) {
	r := rand.New(rand.NewSource(seed))
	g = &Graph{}
	for i := 0; i < vertexCount; i++ {
		g.vertex = append(g.vertex, &vertex{
			index: i,
		})
	}
	// for i := 0; i < vertexCount; i++ {
	// 	for j := 0; j < edgeCount; j++ {
	// 		connected := r.Intn(vertexCount)
	// 		g.AddEdge(i, connected)
	// 	}
	// }
	for i := 0; i < retry; i++ {
		visited := map[int]bool{}
		g.Visit(0, visited, func(*vertex) bool {
			return true
		})
		if len(visited) == vertexCount {
			return g
		}
		for _, v := range g.vertex {
			if !visited[v.index] {
				connected := r.Intn(vertexCount)
				g.AddEdge(v.index, connected)
			}
		}
	}
	fmt.Println("not visited")
	return g
}

func main() {
	g := makeGraph(333, 10, 3)
	fmt.Println(g)
}

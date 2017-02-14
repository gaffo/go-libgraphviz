package main

import (
	"fmt"
	"github.com/gaffo/go-libgraphviz"
)

func main() {
	gv := graphviz.NewGraphviz()
	defer gv.Close()

	fmt.Println("GV Initted")

	g := graphviz.NewGraph("g", graphviz.AGDIRECTED)
	n := graphviz.NewNode(g, "n", 1)
	m := graphviz.NewNode(g, "m", 1)
	graphviz.NewEdge(g, n, m, "label", 1)

	gv.LayoutJobs(g)

	gv.RenderJobs(g)

	gv.FreeLayout(g)

	g.Close()
}

package main

import (
	"fmt"
	"github.com/gaffo/go-libgraphviz"
)

func main() {
	gv := graphviz.NewGraphviz()
	defer gv.Close()

	fmt.Println("GV Initted")
}

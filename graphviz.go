package graphviz

import ()

// #cgo CFLAGS: -I${SRCDIR}/vendor/graphviz_install/include/graphviz
// #cgo LDFLAGS: ${SRCDIR}/libgvc_C.a ${SRCDIR}/libcgraph_C.a ${SRCDIR}/libcdt_C.a ${SRCDIR}/libpathplan_C.a -lexpat -lz -lm -lltdl
// #include "gvc.h"
import "C"

var AGDIRECTED = C.Agdirected

type Graphviz struct {
	gv *C.GVC_t
}

func NewGraphviz() *Graphviz {
	gv := C.gvContext()
	C.gvParseArgs(gv, 0, nil)
	return &Graphviz{
		gv: gv,
	}
}

func (this *Graphviz) Close() {
	C.gvFreeContext(this.gv)
}

func (this *Graphviz) LayoutJobs(graph *Graph) {
	C.gvLayoutJobs(this.gv, graph.g)
}

func (this *Graphviz) RenderJobs(graph *Graph) {
	C.gvRenderJobs(this.gv, graph.g)
}

func (this *Graphviz) FreeLayout(graph *Graph) {
	C.gvFreeLayout(this.gv, graph.g)
}

type Graph struct {
	g *C.Agraph_t
}

func NewGraph(text string, desc C.Agdesc_t) *Graph {
	g := C.agopen(C.CString(text), desc, nil)
	return &Graph{
		g: g,
	}
}

func (this *Graph) Close() {
	C.agclose(this.g)
}

type Node struct {
	n *C.Agnode_t
}

func NewNode(graph *Graph, text string, create int) *Node {
	n := C.agnode(graph.g, C.CString(text), C.int(create))
	return &Node{
		n: n,
	}
}

type Edge struct {
	e *C.Agedge_t
}

func NewEdge(graph *Graph, node1, node2 *Node, text string, create int) *Edge {
	e := C.agedge(graph.g, node1.n, node2.n, C.CString(text), C.int(create))
	return &Edge{
		e: e,
	}
}

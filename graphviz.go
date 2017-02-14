package graphviz

import ()

// #cgo CFLAGS: -I${SRCDIR}/vendor/graphviz_install/include/graphviz
// #cgo LDFLAGS: ${SRCDIR}/libgvc_C.a ${SRCDIR}/libcgraph_C.a ${SRCDIR}/libcdt_C.a ${SRCDIR}/libpathplan_C.a -lexpat -lz -lm -lltdl
// #include "gvc.h"
import "C"

type Graphviz struct {
	gv *C.GVC_t
}

func NewGraphviz() *Graphviz {
	gv := C.gvContext()
	return &Graphviz{
		gv: gv,
	}
}

func (this *Graphviz) Close() {
	C.gvFreeContext(this.gv)
}

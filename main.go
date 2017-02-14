package graphviz

import (
	"fmt"
)

// #cgo CFLAGS: -I${SRCDIR}/vendor/graphviz_install/include/graphviz
// #cgo LDFLAGS: ${SRCDIR}/libgvc_C.a ${SRCDIR}/libcgraph_C.a ${SRCDIR}/libcdt_C.a -lexpat -lzlib -lm
// #include "gvc.h"
import "C"

func main() {
	gv := C.gvContext()
	fmt.Println(gv)
}
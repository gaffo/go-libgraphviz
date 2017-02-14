GRAPHVIZ_COMMIT ?= 27e3785

all: libgvc_C.a install

install:
	go install -x github.com/gaffo/go-libgraphviz/gvtest/...

test: fmt
	go test -v

fmt:
	go fmt *.go > /dev/null

libgvc_C.a: vendor/graphviz
	find vendor/graphviz/ -name "*.a" | xargs -I{} cp {} .

vendor/graphviz:
	mkdir -p vendor
	git clone https://github.com/ellson/graphviz.git vendor/graphviz
	cd vendor/graphviz && git reset --hard && git clean -fdx
	cd vendor/graphviz && git checkout ${GRAPHVIZ_COMMIT}
	cd vendor/graphviz && ./autogen.sh && ./configure --prefix $(CURDIR)/vendor/graphviz_install --enable-perl=no
	cd vendor/graphviz && mkdir $(CURDIR)/vendor/graphviz_install
	cd vendor/graphviz && make && make install

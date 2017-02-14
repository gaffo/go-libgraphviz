GRAPHVIZ_COMMIT ?= 27e3785

all: libgvc.la cgo

cgo:
	go build -x main.go

test: fmt
	go test -v

fmt:
	go fmt *.go > /dev/null

libgvc.la: vendor/graphviz
	find vendor/graphviz/ -name "*.a" | xargs -I{} cp {} .

vendor/graphviz:
	mkdir -p vendor
	git clone https://github.com/ellson/graphviz.git vendor/graphviz
	cd vendor/graphviz && git reset --hard && git clean -fdx
	cd vendor/graphviz && git checkout ${GRAPHVIZ_COMMIT}
	cd vendor/graphviz && ./autogen.sh && ./configure --prefix $(CURDIR)/vendor/graphviz_install
	cd vendor/graphviz && mkdir $(CURDIR)/vendor/graphviz_install
	cd vendor/graphviz && make && make install
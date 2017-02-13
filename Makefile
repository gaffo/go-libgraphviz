GRAPHVIZ_COMMIT ?= 27e3785

all: libgvc.la

libgvc.la: vendor/graphviz
	cp vendor/graphviz_install/lib/*.la .
	cp vendor/graphviz_install/lib/graphviz/*.la .

vendor/graphviz:
	mkdir -p vendor
	git clone https://github.com/ellson/graphviz.git vendor/graphviz
	cd vendor/graphviz && git reset --hard && git clean -fdx
	cd vendor/graphviz && git checkout ${GRAPHVIZ_COMMIT}
	cd vendor/graphviz && ./autogen.sh && ./configure --prefix $(CURDIR)/vendor/graphviz_install
	cd vendor/graphviz && mkdir $(CURDIR)/vendor/graphviz_install
	cd vendor/graphviz && make && make install
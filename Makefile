GOPATH=$(PWD)/.go
SRC_IN_GOPATH=$(GOPATH)/src/github.com/heyLu/echo-chamber

build: echo-chamber load

echo-chamber: setup echo-chamber.go
	cd $(SRC_IN_GOPATH) && go build github.com/heyLu/echo-chamber

load: setup cmd/load/load.go
	cd $(SRC_IN_GOPATH) && go build github.com/heyLu/echo-chamber/cmd/load

setup: $(SRC_IN_GOPATH)

$(SRC_IN_GOPATH):
	mkdir -p $$(dirname $@)
	ln -s $(PWD) $@

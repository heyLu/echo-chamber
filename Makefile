GOPATH=$(PWD)/.go
SRC_IN_GOPATH=$(GOPATH)/src/github.com/heyLu/echo-chamber

build: echo-chamber load

echo-chamber: setup echo-chamber.go
	cd $(SRC_IN_GOPATH) && go install github.com/heyLu/echo-chamber
	@cp $(GOPATH)/bin/$@ $(PWD)

load: setup cmd/load/load.go
	cd $(SRC_IN_GOPATH) && go install github.com/heyLu/echo-chamber/cmd/load
	@cp $(GOPATH)/bin/$@ $(PWD)

setup: $(SRC_IN_GOPATH) $(PWD)/vendor

$(SRC_IN_GOPATH):
	mkdir -p $$(dirname $@)
	ln -s $(PWD) $@

$(PWD)/vendor:
	cd $(SRC_IN_GOPATH) && dep ensure

clean:
	rm -f echo-chamber load

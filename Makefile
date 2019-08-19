name = auction
img = golang:1.12
src = github.com/iuryfukuda/$(name)
workdir = /go/src/$(src)
run = docker run --net=host -it -v $(PWD):$(workdir) -w $(workdir) --rm $(img)

build:
	$(run) go build .

shell: image
	$(run) sh

run:
	$(run) ./start.sh $(args)

test:
ifeq ($(origin args), undefined)
	$(run) go test ./...
else
	$(run) go test $(args)
endif

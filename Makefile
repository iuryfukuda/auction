name = auction
img = golang:1.12
src = github.com/iuryfukuda/$(name)
workdir = /go/src/$(src)
run = docker run -it -v $(PWD):$(workdir) -w $(workdir) --rm $(img)

build:
	$(run) go build .

shell: image
	$(run) sh

run:
	$(run) ./start.sh $(args)

check:
ifeq ($(origin args), undefined)
	$(run) go test ./...
else
	$(run) go test $(args)
endif

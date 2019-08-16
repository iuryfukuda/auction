name = auction
img = golang:1.11
src = github.com/iuryfukuda/$(name)
workdir = /go/src/$(src)
run = docker run -v $(PWD):$(workdir) -w $(workdir) --rm $(img)

build:
	$(run) go build .

shell: image
	$(run) sh

run:
	$(run) ./start.sh $(args)

check:
	$(run) go test ./...

check-args:
	$(run) go test $(args)

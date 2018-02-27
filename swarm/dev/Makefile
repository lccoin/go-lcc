.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/LCC/go-LCC/cmd/swarm
	go build -o bin/geth     github.com/LCC/go-LCC/cmd/geth
	go build -o bin/bootnode github.com/LCC/go-LCC/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/LCC/go-LCC/swarm/...

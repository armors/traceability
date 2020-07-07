# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test build
build:
	rm -rf target/
	mkdir target/
	cp config.json target/config.json
	$(GOBUILD) -o target/admin cmd/admin.go

test:
	$(GOTEST) -v ./...

clean:
	rm -rf target/

run:
	nohup target/admin 2>&1 > target/admin.log &

stop:
	pkill -f target/admin

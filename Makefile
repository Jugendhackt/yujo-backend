GOCMD=go
GOBUILD=$(GOCMD) build -v
GORUN=$(GOCMD) run -v
PKGER=${HOME}/go/bin/pkger

all: run
build:
	$(PKGER)
	$(GOBUILD)
run:
	$(PKGER)
	$(GORUN) .

# Cross compile
build-server:
	$(PKGER)
	GOOS=linux $(GOBUILD) --ldflags '-linkmode external -extldflags "-static"'

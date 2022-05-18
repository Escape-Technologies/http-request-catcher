NAME	:= http-request-catcher
DEL		:= /bin/rm
ENTRY   := cmd/http-request-catcher/main.go


GIT_VERSION	:= $(shell git describe --tags --abbrev=0)
BUILD_TIME	:= $(shell date -u +%Y-%m-%dT%H:%M:%S)
GIT_HASH	:= $(shell git rev-parse --short HEAD)
GIT_BRANCH	:= $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: all
all: $(NAME)

$(NAME):
	go build -ldflags "-X main.Version=$(GIT_VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitHash=$(GIT_HASH) -X main.GitBranch=$(GIT_BRANCH)" -o $(NAME) $(ENTRY)

.PHONY: clean
clean:
	$(DEL) $(NAME)

.PHONY: re
re: clean all

.PHONY: run
run:
	go run $(ENTRY)

.PHONY: test
test:
	golangci-lint run
	gocyclo -over 15 $(ENTRY)

.PHONY: release
release:
	goreleaser release --snapshot --rm-dist

.PHONY: compose
compose:
	docker-compose up --build

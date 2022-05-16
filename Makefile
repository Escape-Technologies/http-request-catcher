NAME	= http-request-catcher
DEL		= /bin/rm
ENTRY	= cmd/http-request-catcher/main.go

.PHONY: all
all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRCS) -w -s -ldflags "-X main.Version=`git describe --tags` -X main.BuildTime=`date -u +%Y-%m-%dT%H:%M:%SZ` -X main.GitHash=`git rev-parse --short HEAD` -X main.GitBranch=`git rev-parse --abbrev-ref HEAD`"

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

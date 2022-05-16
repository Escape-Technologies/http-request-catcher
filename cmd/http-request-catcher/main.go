package main

import (
	"fmt"

	"github.com/Escape-Technologies/http-request-catcher/api"
	"github.com/Escape-Technologies/http-request-catcher/internal"
	"github.com/Escape-Technologies/http-request-catcher/pkg/database"
	"github.com/Escape-Technologies/http-request-catcher/pkg/server"
)

var (
	Version   string
	BuildTime string
	GitHash   string
	GitBranch string
)

func main() {
	fmt.Printf("Starting http-request-catcher %s@%s-%s (%s)\n", Version, GitHash, GitBranch, BuildTime)

	r := server.CreateRouter()

	database.Connect(api.REDIS)

	server.SetupRouter(r)
	server.StartRouter(r, internal.CATCHER_PORT)
}

package main

import (
	"os"

	"github.com/MarkTBSS/go-logger/config"
	"github.com/MarkTBSS/go-logger/modules/servers"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	servers.NewServer(cfg).Start()
}

package main

import (
	"flag"
	"github.com/fastid/fastid/internal/app/command"
	"github.com/fastid/fastid/internal/app/http"
	"os"
)

func main() {
	var help bool
	var runServer bool
	var createSuperUser bool

	flag.BoolVar(&help, "help", false, "List available commands")
	flag.BoolVar(&runServer, "run", false, "Starts the FastID server")
	flag.BoolVar(&createSuperUser, "createsuperuser", false, "Create superuser")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if createSuperUser {
		command.CreateSuperUser()
		os.Exit(0)
	}

	if runServer {
		http.HTTP()
	}
}

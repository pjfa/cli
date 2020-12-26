package main

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"

	"git.arvan.me/arvan/cli/pkg/cli"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	command := cli.NewCommandCLI()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

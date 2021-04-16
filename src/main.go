package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aodpi/hrm-go-graphql/v2/config"
	"github.com/aodpi/hrm-go-graphql/v2/server"
)

func main() {

	environment := flag.String("e", "dev", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*environment)
	server.Init()
}

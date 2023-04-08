package main

import (
	"flag"
	"fmt"
	"ita/monitoring/config"
	"ita/monitoring/server"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	// db.Init()
	server.Init()
}

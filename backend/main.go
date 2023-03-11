package main

import (
	"flag"
	"fmt"
	"os"

	"lilien-server/config"
	"lilien-server/database"
	"lilien-server/routers"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Initlization(*environment)
	database.Initlization()

	r := routers.NewRouter()
	r.Run(config.GetConfig().GetString("server.port"))
}

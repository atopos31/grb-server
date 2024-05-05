package main

import (
	"flag"
	"gvb/config"
)

func main() {
	configDir := flag.String("config", "./config/dev.yaml", "config dir")
	flag.Parse()
	config.Init(*configDir)
}

package main

import (
	"backend/internal/config"
	"flag"
	"fmt"
)

var configFile = flag.String("f", "etc/dev.yaml", "the config file")

func main() {
	flag.Parse()
	serverCfg := config.MustLoadCfg(*configFile, "YML")
	fmt.Printf(serverCfg.Name)
}

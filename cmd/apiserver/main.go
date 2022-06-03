package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/ozon_test/internal/app/apiserver"
)

var (
	configPath  string
	configStore string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to  config file")
	flag.StringVar(&configStore, "store-type", "sql", "available stores: sql, internal")
}
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, &config)
	if isFlagPassed("store-type") {
		config.StoreType = configStore
	}

	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

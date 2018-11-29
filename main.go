package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/short"
	"github.com/rasa/shortme/web"
)

const LOCAL_JSON = "local.json"

var (
	cfgFile = flag.String("c", conf.DEFAULT_CONFIG_JSON, "configuration file")
	version = flag.Bool("v", false, "show version")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Version %s (%s%s)\n", conf.Version, conf.GitCommit, conf.Tags)
		os.Exit(0)
	}

	conf.ParseDefaultConfig()
	conf.ParseConfig(*cfgFile)

	_, err := os.Stat(LOCAL_JSON)
	if err == nil {
		conf.ParseConfig(LOCAL_JSON)
	}

	// short service
	short.Start()

	// api
	web.Start()
}

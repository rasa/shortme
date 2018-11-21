package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/short"
	"github.com/rasa/shortme/web"
)

var (
	cfgFile = flag.String("c", "config.json", "configuration file")
	version = flag.Bool("v", false, "show version")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Version %s (%s%s)\n", conf.Version, conf.GitCommit, conf.Tags)
		os.Exit(0)
	}

	// parse config
	conf.MustParseConfig(*cfgFile)

	// short service
	short.Start()

	// api
	web.Start()
}

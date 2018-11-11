package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yougg/shortme/conf"
	"github.com/yougg/shortme/short"
	"github.com/yougg/shortme/web"
)

var (
	cfgFile = flag.String("c", "config.json", "configuration file")
	version = flag.Bool("v", false, "show version")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println(conf.Version)
		os.Exit(0)
	}

	// parse config
	conf.MustParseConfig(*cfgFile)

	// short service
	short.Start()

	// api
	web.Start()
}

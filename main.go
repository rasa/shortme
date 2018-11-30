package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/short"
	"github.com/rasa/shortme/web"
)

var (
	cfgFile = flag.String("c", conf.DEFAULT_CONFIG_JSON, "configuration file")
	version = flag.Bool("v", false, "show version")
)

func main() {
	basename := filepath.Base(os.Args[0])
	progname := strings.TrimSuffix(basename, filepath.Ext(basename))

	dash := "-"
	if conf.Tags == "" {
		dash = ""
	}

	flag.Parse()

	if *version {
		fmt.Printf("Version %s (%s%s%s)\n", conf.Version, conf.GitCommit, dash, conf.Tags)
		os.Exit(0)
	}

	log.Printf("%s: Version %s (%s%s%s)\n", progname, conf.Version, conf.GitCommit, dash, conf.Tags)
	log.Printf("Built with %s for %s/%s (%d CPUs/%d GOMAXPROCS)\n",
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		runtime.GOMAXPROCS(-1))

	conf.ParseConfigs(*cfgFile)

	// short service
	short.Start()

	// api
	web.Start()

	short.Close()
}

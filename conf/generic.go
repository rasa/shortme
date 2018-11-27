package conf

import (
	"log"
)

var (
	Tags      string
	Version   string
	GitCommit string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

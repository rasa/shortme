package conf

import (
	"log"
)

var Version string

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

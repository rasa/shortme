package conf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type sequenceDB struct {
	DSN          string `json:"dsn"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

type http struct {
	Listen string `json:"listen"`
}

type shortDB struct {
	ReadDSN      string `json:"read_dsn"`
	WriteDSN     string `json:"write_dsn"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

type common struct {
	BlackShortURLs    []string        `json:"black_short_urls"`
	BlackShortURLsMap map[string]bool `json:"-"`
	BaseString        string          `json:"base_string"`
	BaseStringLength  uint64          `json:"-"`
	DomainName        string          `json:"domain_name"`
	Schema            string          `json:"schema"`
	Title             string          `json:"title"`
	ShortURL          string          `json:"short_url"`
}

type config struct {
	Http       http       `json:"http"`
	SequenceDB sequenceDB `json:"sequence_db"`
	ShortDB    shortDB    `json:"short_db"`
	Common     common     `json:"common"`
}

var Conf config

func MustParseConfig(configFile string) {
	if fileInfo, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			log.Panicf("configuration file %v does not exist.", configFile)
		} else {
			log.Panicf("configuration file %v can not be stated. %v", configFile, err)
		}
	} else {
		if fileInfo.IsDir() {
			log.Panicf("%v is a directory name", configFile)
		}
	}

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("read configuration file error. %v", err)
	}
	content = bytes.TrimSpace(content)

	err = json.Unmarshal(content, &Conf)
	if err != nil {
		log.Panicf("unmarshal json object error. %v", err)
	}

	// short url black list
	Conf.Common.BlackShortURLsMap = make(map[string]bool)
	for _, blackShortURL := range Conf.Common.BlackShortURLs {
		Conf.Common.BlackShortURLsMap[blackShortURL] = true
	}

	// base string
	Conf.Common.BaseStringLength = uint64(len(Conf.Common.BaseString))
}

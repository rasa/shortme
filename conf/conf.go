package conf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	ValidURLChars = "0-9A-Za-z$_.+!*'(),-"
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
	ShortURLMax       uint64          `json:"short_url_max"`
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
			log.Panicf("File not found: %v", configFile)
		} else {
			log.Panicf("File not readable: %v: %v", configFile, err)
		}
	} else {
		if fileInfo.IsDir() {
			log.Panicf("File is a directory: %v", configFile)
		}
	}

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("Cannot read %v", err)
	}
	content = bytes.TrimSpace(content)

	err = json.Unmarshal(content, &Conf)
	if err != nil {
		log.Panicf("Read error unmarshaling %v: %v", configFile, err)
	}

	// short url black list
	// @todo add files: index.html, favicon.ico, robots.txt, etc
	Conf.Common.BlackShortURLsMap = make(map[string]bool)
	for _, blackShortURL := range Conf.Common.BlackShortURLs {
		Conf.Common.BlackShortURLsMap[blackShortURL] = true
	}

	// base string
	Conf.Common.BaseStringLength = uint64(len(Conf.Common.BaseString))

	s := fmt.Sprintf("[^%v]+", regexp.QuoteMeta(ValidURLChars))
	log.Printf(s)
	re := regexp.MustCompile(s)

	if re.MatchString(Conf.Common.BaseString) {
		log.Panicf("base_string in %s contains invalid URL characters: %q", configFile, re.FindString(Conf.Common.BaseString))
	}

	for _, char := range Conf.Common.BaseString {
		s := string(char)
		if strings.Count(Conf.Common.BaseString, s) > 1 {
			log.Panicf("base_string in %s contains %d instances of %s", configFile, strings.Count(Conf.Common.BaseString, s), s)
		}
	}
}

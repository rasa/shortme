package conf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
)

const (
	DEFAULT_CONFIG_JSON = "config.json"
	// See https://perishablepress.com/stop-using-unsafe-characters-in-urls/#character-encoding-chart
	ValidURLChars = "0-9A-Za-z$_.+!*'(),-"
)

type Backend string

const (
	BADGER Backend = "badger"
	MYSQL  Backend = "mysql"
	REDIS  Backend = "redis"
)

type sequenceBadger struct {
	Dir        string `json:"dir"`
	ValueDir   string `json:"value_dir"`
	SyncWrites bool   `json:"sync_writes"`
	KeyName    string `json:"key_name"`
	Bandwidth  uint64 `json:"bandwidth"`
}

type sequenceDB struct {
	DSN          string `json:"dsn"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

type sequenceRedis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	PoolSize int    `json:"pool_size"`
	KeyName  string `json:"key_name"`
	//DialTimeout  uint8    `json:"dial_timeout"`
	//PoolTimeout  uint8    `json:"pool_timeout"`
	//ReadTimeout  uint8    `json:"read_timeout"`
	//WriteTimeout uint8    `json:"write_timeout"`
}

type _http struct {
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
	ShortURLMax       uint16          `json:"short_url_max"`
}

type config struct {
	Http            _http          `json:"http"`
	SequenceBackend string         `json:"sequence_backend"`
	SequenceBadger  sequenceBadger `json:"sequence_badger"`
	SequenceDB      sequenceDB     `json:"sequence_db"`
	SequenceRedis   sequenceRedis  `json:"sequence_redis"`
	ShortDB         shortDB        `json:"short_db"`
	Common          common         `json:"common"`
}

var Conf config

func ParseDefaultConfig() {
	if reflect.TypeOf(Assets).String() == "http.Dir" {
		return
	}
	filename := "internal config"
	fh, err := Assets.Open(DEFAULT_CONFIG_JSON)
	if err != nil {
		log.Fatalf("Failed to open %v: %v", filename, err)
	}
	defer fh.Close()
	var data []byte
	data, err = ioutil.ReadAll(fh)
	if err != nil {
		log.Fatalf("Failed to read %v: %v", filename, err)
	}
	parseConfig(filename, data)
}

func ParseConfig(configFile string) {
	if strings.EqualFold(configFile, os.DevNull) {
		return
	}
	log.Printf("Reading %v\n", configFile)
	if fileInfo, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			// Don't fail now if the file isn't found, as we have a default config burned in
			log.Printf("File not found: %v\n", configFile)
			return
		} else {
			log.Fatalf("File not readable: %v: %v\n", configFile, err)
		}
	} else {
		if fileInfo.IsDir() {
			// Don't fail now if the file isn't found, as we have a default config burned in
			log.Printf("File is a directory: %v\n", configFile)
			return
		}
	}

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Cannot read %v\n", err)
	}
	parseConfig(configFile, content)
}

func parseConfig(configFile string, content []byte) {
	content = bytes.TrimSpace(content)

	err := json.Unmarshal(content, &Conf)
	if err != nil {
		log.Fatalf("Read error unmarshaling %v: %v", configFile, err)
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
	re := regexp.MustCompile(s)

	if re.MatchString(Conf.Common.BaseString) {
		log.Fatalf("base_string in %s contains invalid URL characters: %q", configFile, re.FindString(Conf.Common.BaseString))
	}

	for _, char := range Conf.Common.BaseString {
		s := string(char)
		if strings.Count(Conf.Common.BaseString, s) > 1 {
			log.Fatalf("base_string in %s contains %d instances of %s", configFile, strings.Count(Conf.Common.BaseString, s), s)
		}
	}
}

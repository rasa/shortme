package base_test

import (
	"math"
	"testing"

	"github.com/rasa/shortme/base"
	"github.com/rasa/shortme/conf"
)

func init() {
	conf.Conf.Common.BaseString = "Ds3K9ZNvWmHcakr1oPnxh4qpMEzAye8wX5IdJ2LFujUgtC07lOTb6GYBQViSfR"
	conf.Conf.Common.BaseStringLength = uint64(len(conf.Conf.Common.BaseString))
}

func Test_Int2String_Min_UInt64(t *testing.T) {
	var uint64Value uint64 = 0
	var expectedShortURL string = string(conf.Conf.Common.BaseString[0])
	if shortURL := base.Int2String(uint64Value); shortURL == expectedShortURL {
		t.Logf("Min Integer %v passes Int2String", uint64Value)
	} else {
		t.Errorf("Min Integer %v does not pass Int2String. "+
			"Expected short	URL: %s, got: %s",
			uint64Value, expectedShortURL, shortURL)
	}
}

func Test_Int2String_1000_Uint64(t *testing.T) {
	var (
		uint64Value      uint64 = 1000
		expectedShortURL string = "oW"
	)

	if shortURL := base.Int2String(uint64Value); shortURL == expectedShortURL {
		t.Logf("Integer %v passes Int2String", uint64Value)
	} else {
		t.Errorf("Integer %v does not pass Int2String. "+
			"Expected short URL: %s, got: %s",
			uint64Value, expectedShortURL, shortURL)
	}
}

func Test_Int2String_10248794872232232_Uint64(t *testing.T) {
	var (
		uint64Value      uint64 = 10248794872232232
		expectedShortURL string = "0i1lzEeCY"
	)

	if shortURL := base.Int2String(uint64Value); shortURL == expectedShortURL {
		t.Logf("Integer %v passes Int2String", uint64Value)
	} else {
		t.Errorf("Integer %v does not pass Int2String. "+
			"Expected short URL: %s, got: %s",
			uint64Value, expectedShortURL, shortURL)
	}
}

func Test_Int2String_Max_UInt64(t *testing.T) {
	var expectedShortURL string = "4fUPJsNHPI1"
	if shortURL := base.Int2String(uint64(math.MaxUint64)); shortURL == expectedShortURL {
		t.Logf("Max Integer %v passes Int2String", uint64(math.MaxUint64))
	} else {
		t.Errorf("Max Integer %v does not pass Int2String. "+
			"Expected short URL: %s, got: %s",
			uint64(math.MaxUint64), expectedShortURL, shortURL)
	}
}

func Test_String2Int_With_First_Char(t *testing.T) {
	var (
		shortURL    string = string(conf.Conf.Common.BaseString[0])
		expectedSeq uint64 = 0
	)

	if seq := base.String2Int(shortURL); seq == expectedSeq {
		t.Logf("Short URL %v passes String2Int", shortURL)
	} else {
		t.Errorf("Short URL %v do not pass String2Int. "+
			"Expected Sequence: %v, got: %v",
			shortURL, expectedSeq, seq)
	}
}

func Test_String2Int_With_oW(t *testing.T) {
	var (
		shortURL    string = "oW"
		expectedSeq uint64 = 1000
	)

	if seq := base.String2Int(shortURL); seq == expectedSeq {
		t.Logf("Short URL %v passes String2Int", shortURL)
	} else {
		t.Errorf("Short URL %v do not pass String2Int. "+
			"Expected Sequence: %v, got: %v",
			shortURL, expectedSeq, seq)
	}
}

func Test_String2Int_0i1lzEeCY_Uint64(t *testing.T) {
	var (
		shortURL    string = "0i1lzEeCY"
		expectedSeq uint64 = 10248794872232232
	)

	if seq := base.String2Int(shortURL); seq == expectedSeq {
		t.Logf("Short URL %v passes String2Int", shortURL)
	} else {
		t.Errorf("Short URL %v do not pass String2Int. "+
			"Expected Sequence: %v, got: %v",
			shortURL, expectedSeq, seq)
	}
}

func Test_String2Int_With_4fUPJsNHPI1(t *testing.T) {
	var (
		shortURL    string = "4fUPJsNHPI1"
		expectedSeq uint64 = math.MaxUint64
	)

	if seq := base.String2Int(shortURL); seq == expectedSeq {
		t.Logf("Short URL %v passes String2Int", shortURL)
	} else {
		t.Errorf("Short URL %v do not pass String2Int. "+
			"Expected Sequence: %v, got: %v",
			shortURL, expectedSeq, seq)
	}
}

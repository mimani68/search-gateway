package scraping

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/imroc/req"
	"golang.org/x/net/html/charset"
	"market.ir/config"
	"market.ir/pkg/logHandler"
)

func ScrapingHandler(keyword string, where string) bool {
	var result interface{}
	header := req.Header{
		"Accept": "application/json",
		// "Authorization": "Basic YWRtaW46YWRtaW4=",
	}
	param := req.Param{}
	query := keyword
	resultNumber := config.GetConfig("resultNumber")
	url := config.GetConfig("SCRAPER") + "/?q=" + url.QueryEscape(query) + "&source=" + where + "&itemNumber=" + resultNumber
	logHandler.Log(url)
	r, err := req.Get(url, header, param)
	if err != nil {
		log.Fatal(err)
		return false
	}
	r.ToJSON(&result)
	return true
}

func convrtToUTF8(str string, origEncoding string) string {
	strBytes := []byte(str)
	byteReader := bytes.NewReader(strBytes)
	reader, _ := charset.NewReaderLabel(origEncoding, byteReader)
	strBytes, _ = ioutil.ReadAll(reader)
	return string(strBytes)
}

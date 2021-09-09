package sdk

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/imroc/req"
	"golang.org/x/net/html/charset"
	"market.ir/dto"
)

func Search(keyword string, tag string) dto.TaggerArticleList {
	var result interface{}
	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic YWRtaW46YWRtaW4=",
	}
	param := req.Param{}
	query := convrtToUTF8(keyword, "ascii")
	url := "http://scraper:3000/?q=" + query + "&source=zoomit&itemNumber=20"
	r, err := req.Get(url, header, param)
	if err != nil {
		log.Fatal(err)
	}
	r.ToJSON(&result) // response => struct/map
	//
	// DEBUG
	//
	// log.Printf("%+v", r) // print info (try it, you may surprise)

	return dto.TaggerArticleList{
		Tag:  tag,
		Data: result,
	}
}

func convrtToUTF8(str string, origEncoding string) string {
	strBytes := []byte(str)
	byteReader := bytes.NewReader(strBytes)
	reader, _ := charset.NewReaderLabel(origEncoding, byteReader)
	strBytes, _ = ioutil.ReadAll(reader)
	return string(strBytes)
}

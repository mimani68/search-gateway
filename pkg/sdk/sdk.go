package sdk

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

// func Search(keyword string, where string, tag string) dto.TaggerArticleList {
// 	var result interface{}
// 	header := req.Header{
// 		"Accept": "application/json",
// 		// "Authorization": "Basic YWRtaW46YWRtaW4=",
// 	}
// 	param := req.Param{}
// 	query := keyword
// 	resultNumber := config.GetConfig("resultNumber")
// 	url := config.GetConfig("SCRAPER") + "/?q=" + url.QueryEscape(query) + "&source=" + where + "&itemNumber=" + resultNumber
// 	logHandler.Log(url)
// 	r, err := req.Get(url, header, param)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	r.ToJSON(&result)
// 	//
// 	// DEBUG
// 	//
// 	return dto.TaggerArticleList{
// 		Tag:  tag,
// 		Data: result,
// 	}
// }

func Search(keyword string, where string, tag string) interface{} {
	var result interface{}
	//
	// DEBUG
	//
	// return dto.TaggerArticleList{
	// 	Tag:  tag,
	// 	Data: result,
	// }
	return result
}

func convrtToUTF8(str string, origEncoding string) string {
	strBytes := []byte(str)
	byteReader := bytes.NewReader(strBytes)
	reader, _ := charset.NewReaderLabel(origEncoding, byteReader)
	strBytes, _ = ioutil.ReadAll(reader)
	return string(strBytes)
}

package textminipulator

import (
	"fmt"
	"strings"
)

func ShowText(text string, query string) map[string]interface{} {
	text = CleanFarsiText(text)
	paragraphList := strings.Split(text, "\n")

	result := map[string]interface{}{}
	result["success"] = "false"

	for _, value := range paragraphList {
		doc := ExtractSentences(value)
		for _, sent := range doc {
			if strings.Contains(sent[0], query) {
				result["selectItem"] = sent[0]
				result["selectItemIndex"] = strings.Index(text, sent[0])
				result["success"] = "true"
			}
		}
	}

	if result["success"] != "true" {
		return result
	}

	result["pre"] = text[:result["selectItemIndex"].(int)]
	postStartingIndex := result["selectItemIndex"].(int) + len(fmt.Sprintf("%s", result["selectItem"]))
	result["post"] = text[postStartingIndex:] + "..."
	result["complete"] = text
	return result
}

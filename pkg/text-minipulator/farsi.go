package textminipulator

import (
	"regexp"
)

func CleanFarsiText(text string) string {
	removeHtmlTagsPattern, _ := regexp.Compile("(?mi)<[^>]*>")
	result := removeHtmlTagsPattern.ReplaceAllString(text, "")

	removeScriptTagsPattern, _ := regexp.Compile("(?mi)<script[\\s\\S]*?>[\\s\\S]*?<\\/script>")
	result = removeScriptTagsPattern.ReplaceAllString(result, "")

	removeNumberTagsPattern, _ := regexp.Compile("(?mi).\\s{1,}[۰۱۲۳۴۵۶۷۸۹]")
	result = removeNumberTagsPattern.ReplaceAllString(result, "")

	removeInvalidCharsPattern, _ := regexp.Compile("(?mi)(&#13;)")
	result = removeInvalidCharsPattern.ReplaceAllString(result, "")

	removeDoubleSpacePattern, _ := regexp.Compile("(?mi)\\s{2,}")
	result = removeDoubleSpacePattern.ReplaceAllString(result, " ")

	return result
}

func ExtractSentences(text string) [][]string {
	onlyValidString, _ := regexp.Compile("[a-zA-Zضصثقفغعهخحجچشسیبلاتنمکگظطزرذدپو].*?(\\.\\s|\\?\\s|\\!\\s|؟\\s|؛\\s)")
	list := onlyValidString.FindAllStringSubmatch(text, -1)
	return list
}

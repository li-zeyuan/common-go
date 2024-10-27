package utils

import (
	"html"
	"regexp"
)

var htmlRe = regexp.MustCompile("<.*?>")

func TrimHtml(htmlStr string) string {
	if len(htmlStr) == 0 {
		return ""
	}
	return htmlRe.ReplaceAllString(htmlStr, "")
}

func HtmlUnescapeAndTrim(htmlStr string) string {
	if len(htmlStr) == 0 {
		return ""
	}
	return TrimHtml(html.UnescapeString(htmlStr))
}

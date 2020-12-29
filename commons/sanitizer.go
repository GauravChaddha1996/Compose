package commons

import "github.com/microcosm-cc/bluemonday"

var InputSanitizerStrictPolicy = bluemonday.StrictPolicy()
var InputSanitizerAllowedHtmlPolicy = bluemonday.UGCPolicy()

func StrictSanitizeString(input string) string {
	return InputSanitizerStrictPolicy.Sanitize(input)
}

func UgcSanitizeString(input string) string {
	return InputSanitizerAllowedHtmlPolicy.Sanitize(input)
}

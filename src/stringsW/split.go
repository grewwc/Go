package stringsW

import (
	"bytes"
	"strings"
)

// SplitNoEmpty remove empty strings
func SplitNoEmpty(str, sep string) []string {
	var res []string
	if str == "" {
		return res
	}
	for _, s := range strings.Split(str, sep) {
		if s == "" {
			continue
		}
		res = append(res, s)
	}
	return res
}

// SplitNoEmptyKeepQuote keep content in quote intact
func SplitNoEmptyKeepQuote(str string, sep byte) []string {
	inQuote := false
	var res []string
	var word bytes.Buffer
	if str == "" {
		return res
	}

	for i := range str {
		s := str[i]
		if s == '"' {
			inQuote = !inQuote
		} else if s != sep || inQuote {
			word.WriteByte(s)
		} else if word.Len() != 0 {
			res = append(res, word.String())
			word.Reset()
		}
	}
	if word.Len() != 0 {
		res = append(res, word.String())
	}
	return res
}

func GetLastItem(slice []string) string {
	if len(slice) < 1 {
		return ""
	}
	return slice[len(slice)-1]
}

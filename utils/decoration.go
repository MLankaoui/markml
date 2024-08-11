package utils

import "strings"

func Decoration(word string) string {
	if strings.HasPrefix(word, "**") && strings.HasSuffix(word, "**") {
		word = strings.TrimPrefix(word, "**")
		word = strings.TrimSuffix(word, "**")

		return "<bold>" + strings.TrimSpace(word) + "</bold>"
	}

	if strings.HasPrefix(word, "~~") && strings.HasSuffix(word, "~~") {
		word = strings.TrimPrefix(word, "~~")
		word = strings.TrimSuffix(word, "~~")

		return "<s>" + strings.TrimSpace(word) + "</s>"
	}
	if strings.HasPrefix(word, "*") && strings.HasSuffix(word, "*") {
		word = strings.TrimPrefix(word, "*")
		word = strings.TrimSuffix(word, "*")

		return "<em>" + strings.TrimSpace(word) + "</em>"
	}

	return word
}

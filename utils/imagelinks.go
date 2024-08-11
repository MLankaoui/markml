package utils

import "strings"

func ImageLinks(text string) string {
	if strings.HasPrefix(text, "[A link](") && strings.HasSuffix(text, ")") {
		text = strings.TrimPrefix(text, "[A link](")
		text = strings.TrimSuffix(text, ")")

		return "<a href='" + strings.TrimSpace(text) + "'>"
	}
	if strings.HasPrefix(text, "![") && strings.Contains(text, "](") && strings.HasSuffix(text, ")") {
		// Extract the alt text
		altTextStart := strings.Index(text, "![") + len("![")
		altTextEnd := strings.Index(text, "](")
		altText := text[altTextStart:altTextEnd]

		// Extract the URL
		urlStart := altTextEnd + len("](")
		urlEnd := len(text) - 1
		url := text[urlStart:urlEnd]

		return "<img src='" + strings.TrimSpace(url) + "' alt='" + strings.TrimSpace(altText) + "'>"
	}
	return ""
}

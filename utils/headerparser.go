package utils

import "strings"

func HeaderParser(line string) string {
	if strings.HasPrefix(line, "# ") {
		line = strings.TrimPrefix(line, "# ")
		return "<h1>" + strings.TrimSpace(line) + "</h1>"
	}
	if strings.HasPrefix(line, "## ") {
		line = strings.TrimPrefix(line, "## ")
		return "<h2>" + strings.TrimSpace(line) + "</h2>"
	}
	if strings.HasPrefix(line, "### ") {
		line = strings.TrimPrefix(line, "### ")
		return "<h3>" + strings.TrimSpace(line) + "</h3>"
	}
	if strings.HasPrefix(line, "#### ") {
		line = strings.TrimPrefix(line, "#### ")
		return "<h4>" + strings.TrimSpace(line) + "</h4>"
	}
	if strings.HasPrefix(line, "##### ") {
		line = strings.TrimPrefix(line, "##### ")
		return "<h5>" + strings.TrimSpace(line) + "</h5>"
	}
	if strings.HasPrefix(line, "###### ") {
		line = strings.TrimPrefix(line, "###### ")
		return "<h6>" + strings.TrimSpace(line) + "</h6>"
	}
	return ""
}

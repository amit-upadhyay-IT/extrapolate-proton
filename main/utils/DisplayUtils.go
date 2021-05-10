package utils

import "strings"

func GetNoteSegment(text string) string {
	begIndex := strings.Index(text, "/* NOTE BEGIN") + 13
	endIndex := strings.Index(text, "NOTE END */")
	if begIndex == -1 || endIndex == -1 {
		return ""
	}
	return text[begIndex:endIndex]
}

func RemoveNoteSegmentFromCode(text string) string {
	begIndex := strings.Index(text, "/* NOTE BEGIN")
	endIndex := strings.Index(text, "NOTE END */") + 11
	if begIndex == -1 || endIndex == -1 {
		return text
	}
	return text[:begIndex] + text[endIndex:]
}

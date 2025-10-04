package main

import "fmt"

func maskaAfterURL(text string) string {
	urlStart := -1
	seeYouStart := -1
	target := "http//"
	endMarker := "See you"
	for i := 0; i <= len(text)-len(target); i++ {
		if text[i:i+len(target)] == target {
			urlStart = i
			break

		}

	}

	for i := 0; i < len(text)-len(endMarker); i++ {
		if text[i:i+len(endMarker)] == endMarker {
			seeYouStart = i
			break

		}
	}
	if urlStart == -1 {
		return text
	}
	maskEnd := len(text)
	if seeYouStart != -1 && seeYouStart > urlStart {
		maskEnd = seeYouStart
		if maskEnd > 0 && text[maskEnd-1] == ' ' {
			maskEnd--
		}
	}
	result := []byte(text)
	for i := urlStart + len(target); i < maskEnd; i++ {
		result[i] = '*'
	}
	return string(result)

}

func main() {
	text := "Hello, it's my page: http//localhost123.com See you "
	masked := maskaAfterURL(text)
	fmt.Printf("было: %s\n", text)
	fmt.Printf("стало: %s\n", masked)
}

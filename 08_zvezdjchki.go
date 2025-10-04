package main

import "fmt"

func maskaAfterURL(text string) string {
	result := []byte(text)
	target := "http//"

	for i := 0; i <= len(text)-len(target); i++ {
		if string(result[i:i+len(target)]) == target {

			for j := i + len(target); j < len(result); j++ {

				if result[j] == ' ' {
					break
				}
				result[j] = '*'
			}
			i += len(target)
		}
	}
	return string(result)
}

func main() {
	text := "Hello, it's my page: http//localhost123.com See you"
	masked := maskaAfterURL(text)
	fmt.Printf("было: %s\n", text)
	fmt.Printf("стало: %s\n", masked)
}


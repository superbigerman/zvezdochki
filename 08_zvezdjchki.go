package main

import "fmt"

func maskBytesByPosition(data []byte, start, length int) []byte {
	result := make([]byte, len(data))
	copy(result, data)

	if start < 0 || start >= len(result) {
		return result
	}

	end := start + length
	if end > len(result) {
		end = len(result)
	}

	for i := start; i < end; i++ {
		result[i] = '*'

	}
	return result
}
func main() {
	text := []byte("Hello, it's my page: http//localhost123.com See you ")
	masked := maskBytesByPosition(text, 27, 16)
	fmt.Printf("Было: %s\n", text)
	fmt.Printf("Стало: %s\n", masked)

}

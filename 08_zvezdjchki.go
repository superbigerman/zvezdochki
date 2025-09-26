package main

import "fmt"

func maskBytesExactMatch(data []byte, target []byte) []byte {
	result := make([]byte, len(data))
	copy(result, data)

	targetLen := len(target)
	if targetLen == 0 {
		return result
	}
	for i := 0; i <= len(result)-targetLen; i++ {
		match := true
		for j := 0; j < targetLen; j++ {
			if result[i+j] != target[j] {
				match = false
				break
			}
		}

		if match {
			for j := 0; j < targetLen; j++ {
				result[i+j] = '*'
			}
		}
	}

	return result
}

func main() {
	text := []byte("Hello, it's my page:\nhttp://localhost123.com See you")
	target := []byte("localhost123.com")

	masked := maskBytesExactMatch(text, target)
	fmt.Printf("Было: %s\n", text)
	fmt.Printf("Стало:%s\n", masked)

}

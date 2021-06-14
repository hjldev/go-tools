package logger

import "testing"

func TestLog(t *testing.T) {
	NewLogger("")
	for {
		Println("test")
	}
}

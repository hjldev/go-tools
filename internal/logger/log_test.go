package logger

import "testing"

func TestLog(t *testing.T) {
	InitLogger("")
	for {
		Println("test")
	}
}

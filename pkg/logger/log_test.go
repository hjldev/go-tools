package logger

import "testing"

func TestLog(t *testing.T) {
	NewLogger("")
	WriteStr("test")
}

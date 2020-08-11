package providers

import (
	"testing"
)

func TestLoggerTest(test *testing.T) {
	_, err := LoggerProvider()

	if err != nil {
		test.Error("Failed to create logger")
	}
}

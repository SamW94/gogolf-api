package helpers

import (
	"log"
	"strings"
	"testing"
)

func LogString(str string) {
	log.Printf("Provided string: %s", str)
}

func TestCaptureLogOutput(t *testing.T) {
	logOutput := CaptureLogOutput(func() {
		LogString("whoopdedoo")
	})

	if !strings.Contains(logOutput, "Provided string: whoopdedoo") {
		t.Errorf("expected log output to contain message, got %q", logOutput)
	}
}

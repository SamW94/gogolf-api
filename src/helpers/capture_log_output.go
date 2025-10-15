package helpers

import (
	"bytes"
	"log"
)

// Helper function to temporarily set the logging output destination as a temporary
// buffer, then returns the message in the log as a string. The logging output destination
// is set back to the original logging output destination by the deferred function.

func CaptureLogOutput(functionToLog func()) string {
	var tempLogBuffer bytes.Buffer
	originalLogOutputDestination := log.Writer()
	log.SetOutput(&tempLogBuffer)
	defer log.SetOutput(originalLogOutputDestination)
	functionToLog()
	return tempLogBuffer.String()
}

package adapter

import (
	"errors"
	"strings"
	"testing"
)

func TestStdLoggerErrorHandling(t *testing.T) {
	logger := NewStdLogger()
	
	testErr := errors.New("test error message")
	result := logger.buildContext("error", F("error", testErr))
	
	if !strings.Contains(result, "error=test error message") {
		t.Errorf("Expected error message to be included in log output, got: %s", result)
	}
	
	result = logger.buildContext("info", F("count", 42))
	
	if !strings.Contains(result, "count=42") {
		t.Errorf("Expected non-error value to be included in log output, got: %s", result)
	}
}

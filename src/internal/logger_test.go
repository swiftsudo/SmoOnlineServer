package internal

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test Log levels function
func TestLogLevels(t *testing.T) {
	// Test Log level strings
	assert.Equal(t, "Debug", LevelDebug.String())
	assert.Equal(t, "Info", LevelInfo.String())
	assert.Equal(t, "Warn", LevelWarn.String())
	assert.Equal(t, "Error", LevelError.String())

	// Test Log level colors
	assert.Equal(t, LevelInfoColor, LevelInfo.Color())
	assert.Equal(t, LevelWarnColor, LevelWarn.Color())
	assert.Equal(t, LevelErrorColor, LevelError.Color())
}

// Test PrefixNewLines function
func TestPrefixNewLines(t *testing.T) {
	testLogTime := time.Now().Format(time.RFC3339)
	testLogLevel := LevelInfo.String()
	testLogSource := "test"

	// Test no new lines
	testMessage := "testing"
	expectedOutput := fmt.Sprintf(LogFormat, testLogTime, testLogLevel, testLogSource, testMessage)
	assert.Equal(t, expectedOutput, PrefixNewLines(testLogTime, testLogLevel, testLogSource, testMessage))

	// Test multiple new lines
	testMessage = strings.Repeat("testing\n", 3)
	expectedOutput = strings.Repeat(fmt.Sprintf(LogFormat, testLogTime, testLogLevel, testLogSource, "testing"), 3)

	assert.Equal(t, expectedOutput, PrefixNewLines(testLogTime, testLogLevel, testLogSource, testMessage))
}

// TestLogger functions
func TestLogger(t *testing.T) {
	testLogger := NewLogger("test")

	// Ensure logger is not nil
	assert.NotNil(t, testLogger)
	// Ensure we have a valid internal logger
	assert.NotNil(t, testLogger.logger)

	// Ensure Logger name is set correctly
	assert.Equal(t, "test", testLogger.Name(), "Logger name is not set correctly")

	// Change Logger name
	testLogger.SetName("test2")
	assert.Equal(t, "test2", testLogger.Name(), "Logger name is not set correctly")
}

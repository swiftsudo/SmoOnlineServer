package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Level int

const (
	// Level Constants
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError

	// Color Constants
	LevelInfoColor  = "\x1b[37m" // White
	LevelWarnColor  = "\x1b[33m" // Yellow
	LevelErrorColor = "\x1b[31m" // Red
	ResetColor      = "\x1b[0m"  // Resets terminal color

	// Log Format
	LogFormat = "%s %s [%s] %s\n" // {logtime} {level} [{source}] {message}
)

// Logger struct
type Logger struct {
	logger *log.Logger
	name   string
}

// Get the string for the given log level
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "Debug"
	case LevelInfo:
		return "Info"
	case LevelWarn:
		return "Warn"
	case LevelError:
		return "Error"
	default:
		return "Unknown"
	}
}

// Get the color for the given log level
func (l Level) Color() string {
	switch l {
	case LevelInfo:
		return LevelInfoColor
	case LevelWarn:
		return LevelWarnColor
	case LevelError:
		return LevelErrorColor
	default:
		return ""
	}
}

// Prefix all messages with the standard log format prefix
func PrefixNewLines(logtime, level, source, message string) string {
	result := []string{}
	message = strings.Trim(message, "\n")
	lines := strings.Split(message, "\n")

	for _, line := range lines {
		outputMessage := fmt.Sprintf(LogFormat, logtime, level, source, line)
		result = append(result, outputMessage)
	}

	return strings.Join(result, "")
}

// Initialize a new logger
func NewLogger(name string) *Logger {
	defaultLogger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &Logger{
		logger: defaultLogger,
		name:   name,
	}
}

// Get Logger Name
func (l *Logger) Name() string {
	return l.name
}

// Set Logger Name
func (l *Logger) SetName(name string) {
	l.name = name
}

// Log an info message
func (l *Logger) Info(message string) {
	l.HandleLog(LevelInfo, message)
}

// Log a warn message
func (l *Logger) Warn(message string) {
	l.HandleLog(LevelWarn, message)
}

// Log an error message
func (l *Logger) Error(message string) {
	l.HandleLog(LevelError, message)
}

// Log a debug message
func (l *Logger) Debug(message string) {
	l.HandleLog(LevelDebug, message)
}

// Handle the log message
func (l *Logger) HandleLog(level Level, message string) {
	logtime := time.Now().Format(time.RFC3339)

	outputMessage := PrefixNewLines(logtime, level.String(), l.Name(), message)

	coloredMessage := fmt.Sprintf("%s%s%s", level.Color(), outputMessage, ResetColor)

	l.logger.Print(coloredMessage)
}

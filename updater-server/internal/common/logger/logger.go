package logger

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

type Logger struct {
	Output 			io.Writer
	MinLogLevel 	LogLevel
}

type LogLevel int

const (
	LogDebug LogLevel = iota
	LogInfo
	LogWarn
	LogError
)

type LogEntry struct {
	Level     LogLevel
	Message   string
	Fields    map[string]any
	Timestamp time.Time
}

func (l LogLevel) String() string {
	switch l {
		case LogDebug:
			return "DEBUG"
		case LogInfo:
			return "INFO"
		case LogWarn:
			return "WARN"
		case LogError:
			return "ERROR"
		default:
			return "UNKNOWN"
	}
}



func (logger *Logger) log(level LogLevel, message string, fields ...any) {
	if level < logger.MinLogLevel {
		return
	}

	entry := LogEntry{
		Level:     level,
		Message:   message,
		Fields:    make(map[string]interface{}),
		Timestamp: time.Now(),
	}

	// Add extra data fields (key-value pairs); similar to slog
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			if key, ok := fields[i].(string); ok {
				entry.Fields[key] = fields[i+1]
			}
		}
	}

	if err := logger.write(entry); err != nil {
		// Try to write directly to the output if logging fails
		fmt.Fprintf(logger.Output, "Error writing log entry: %s\n", err)
		return
	}
}

func (logger *Logger) write(entry LogEntry) error {
	timestamp := entry.Timestamp.Format("2006-01-02 15:04:05")

	// Build the main log line
	logLine := fmt.Sprintf("[%s] %s %s", timestamp, entry.Level.String(), entry.Message)

	// Add fields if any
	if len(entry.Fields) > 0 {
		fieldsStr := formatFields(entry.Fields)
		logLine += " " + fieldsStr
	}

	_, err := fmt.Fprintln(logger.Output, logLine)
	return err
}

func (logger *Logger) Debug(message string, fields ...any) {
	logger.log(LogDebug, message, fields...)
}

func (logger *Logger) Info(message string, fields ...any) {
	logger.log(LogInfo, message, fields...)
}

func (logger *Logger) Warn(message string, fields ...any) {
	logger.log(LogWarn, message, fields...)
}

func (logger *Logger) Error(message string, fields ...any) {
	logger.log(LogError, message, fields...)
}

func formatFields(fields map[string]any) string {
	if len(fields) == 0 {
		return ""
	}

	// Sort keys for consistent output
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	pairs := make([]string, 0, len(fields))
	for _, key := range keys {
		value := fields[key]
		pairs = append(pairs, fmt.Sprintf("%s=%v", key, value))
	}

	return strings.Join(pairs, " ")
}

// Mock Logger for testing purposes
func MockLogger() (Logger, *os.File, error) {
    reader, writer, err := os.Pipe()
    if err != nil {
        return Logger{}, nil, err
    }

    return Logger{Output: writer, MinLogLevel: 0}, reader, nil
}
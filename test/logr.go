package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-logr/logr"
)

// Log returns a valid Logger that discards all messages logged to it.
// It can be used whenever the caller is not interested in the logs.
func Log() logr.Logger {
	sb := bytes.NewBufferString("")

	return StringLogger{sb}
}

// StringLogger is a Logger that discards all messages.
type StringLogger struct {
	sb *bytes.Buffer
}

func (l StringLogger) Enabled() bool {
	return true
}

func (l StringLogger) Info(msg string, keysAndValues ...interface{}) {
	kv := []interface{}{
		time.Now().String(),
		"[INFO]",
	}

	kv = append(kv, keysAndValues...)

	fmt.Fprintln(l.sb, kv...)
}

func (l StringLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	kv := []interface{}{
		time.Now().String(),
		"[ERROR]",
	}

	kv = append(kv, keysAndValues...)

	fmt.Fprintln(l.sb, kv...)
}

func (l StringLogger) V(level int) logr.Logger {
	return l
}

func (l StringLogger) WithValues(keysAndValues ...interface{}) logr.Logger {
	return l
}

func (l StringLogger) WithName(name string) logr.Logger {
	return l
}

func (l StringLogger) String() string {
	return l.sb.String()
}

// Verify that it actually implements the interface
var _ logr.Logger = StringLogger{}

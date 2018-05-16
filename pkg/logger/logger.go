package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// DebugLogger is a LogWriter that writes to StdErr if enabled by user
	debugLogger = log.New(os.Stderr, "\x1B[1;36mDEBUG: ", log.LstdFlags|log.Lshortfile)
	// InfoLogger is a LogWriter that writes to StdOut
	infoLogger = log.New(os.Stdout, "\x1B[1;31mINFO: ", log.LstdFlags|log.Lshortfile)
)

// LogWriter Exposes our customer LogWriter type for enabling and disabling
type LogWriter struct {
	io.Writer
}

// ShowDebug shows debug messages via Stderr
func ShowDebug() {
	debugLogger.SetOutput(os.Stderr)
}

// HideDebug causes debug messages to be discarded before being displayed
func HideDebug() {
	debugLogger.SetOutput(ioutil.Discard)
}

// Debug logs a debug message to STDERR terminated with newline
func Debug(m ...interface{}) {
	o := fmt.Sprint(m...)
	debugLogger.Output(2, fmt.Sprintf("\x1B[0m%s\n", o))
}

// Debugf logs a user-formatted debug message to STDERR
func Debugf(m ...interface{}) {
	o := fmt.Sprint(m...)
	debugLogger.Output(2, fmt.Sprintf("\x1B[0m%s", o))
}

// Info logs a message to STDOUT terminated with newline
func Info(m ...interface{}) {
	o := fmt.Sprint(m...)
	infoLogger.Output(2, fmt.Sprintf("\x1B[0m%s\n", o))
}

// Infof logs a user-formatted message to STDOUT
func Infof(m ...interface{}) {
	o := fmt.Sprint(m...)
	infoLogger.Output(2, fmt.Sprintf("\x1B[0m%s", o))
}

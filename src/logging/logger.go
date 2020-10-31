package logging

import (
	"fmt"
)

// Log writes log files, logs a string into a file
func (s Self) Log(msg string) {
	println(msg)
	s.writeLogLine(msg)
}

// Logf writes log, but formated
func (s Self) Logf(msg string, a ...interface{}) {
	str := fmt.Sprintf(msg, a...)
	s.Log(str)
}

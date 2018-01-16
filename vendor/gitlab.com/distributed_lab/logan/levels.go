package logan

import "github.com/sirupsen/logrus"

type Level logrus.Level

const (
	// DEPRECATED
	PanicLevel Level = iota
	// DEPRECATED
	FatalLevel
	// DEPRECATED
	ErrorLevel
	// DEPRECATED
	WarnLevel
	// DEPRECATED
	InfoLevel
	// DEPRECATED
	DebugLevel
)

// DEPRECATED
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
}

// DEPRECATED
func ParseLevel(level string) (Level, error) {
	l, err := logrus.ParseLevel(level)
	return Level(l), err
}

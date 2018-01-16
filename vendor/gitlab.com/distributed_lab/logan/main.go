package logan

import (
	"os"

	"github.com/sirupsen/logrus"
	"io"
)

// DEPRECATED
type F logrus.Fields

// DEPRECATED
func New() *Entry {
	lastLogLevel := AllLevels[len(AllLevels)-1]
	return NewWithLevel(lastLogLevel)
}

// DEPRECATED
func NewWithLevel(level Level) (result *Entry) {
	return NewWithLevelFormatter(level, nil)
}

// DEPRECATED
func NewWithJSONFormatter() (result *Entry) {
	lastLogLevel := AllLevels[len(AllLevels)-1]
	return NewWithLevelJSONFormatter(lastLogLevel)
}

// DEPRECATED
func NewWithLevelJSONFormatter(level Level) (result *Entry) {
	return NewWithLevelFormatter(level, &logrus.JSONFormatter{})
}

// DEPRECATED
func NewWithLevelFormatter(level Level, formatter Formatter) (result *Entry) {
	return NewWithLevelFormatterOut(level, formatter, nil)
}

// DEPRECATED
func NewWithLevelFormatterOut(level Level, formatter Formatter, out io.Writer) (result *Entry) {
	logger := logrus.New()
	logger.Level = logrus.Level(level)
	if formatter != nil {
		logger.Formatter = logrus.Formatter(formatter)
	}
	if out != nil {
		logger.Out = out
	}

	result = &Entry{
		logrus.NewEntry(logger).WithField("pid", os.Getpid()),
	}
	return
}

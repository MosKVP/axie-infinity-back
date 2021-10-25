package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = &logrus.Logger{
	Out: os.Stderr,
	Formatter: &logrus.TextFormatter{
		DisableQuote:     true,
		QuoteEmptyFields: true,
	},
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
}

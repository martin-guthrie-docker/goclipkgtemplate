package log

import (
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

// Term is the terminal logger for the Genesis application.
var Term *logrus.Logger

func setupTerm() {
	Term = logrus.New()
	Term.Out = os.Stdout

	Term.Formatter = new(prefixed.TextFormatter)

	filenameHook := filename.NewHook()
	filenameHook.Field = "src"
	Term.AddHook(filenameHook)

	// Default starting log level
	Term.Level = logrus.WarnLevel
	//Term.Level = logrus.DebugLevel
}

func init() {
	setupTerm()
}
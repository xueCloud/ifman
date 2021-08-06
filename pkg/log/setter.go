package log

import (
	"github.com/sirupsen/logrus"
)

func SetLog(level string, w Writer, useSyslog bool) error {
	levelInt, err := LevelStr2Int(level)
	if err != nil {
		return err
	}

	logrusLevel, err := Level2Logrus(levelInt)
	if err != nil {
		return err
	}

	logrus.SetLevel(logrus.Level(logrusLevel))

	//logrus.SetFormatter(&logrus.JSONFormatter{
	//	TimestampFormat:   "2006-01-02 15:04:05",
	//	DisableTimestamp:  false,
	//	DisableHTMLEscape: true,
	//	DataKey:           "",
	//	FieldMap:          nil,
	//	CallerPrettyfier:  nil,
	//	PrettyPrint:       false,
	//})

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})

	logrus.SetOutput(w.GetWriter())

	if useSyslog {
		err = enableSyslog()
		if err != nil {
			return err
		}
	}

	return nil
}

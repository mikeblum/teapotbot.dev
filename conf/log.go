package conf

import (
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/sirupsen/logrus"
)

// logging configuration

const (
	jsonLog      = "JSON"
	envLogLevel  = "LOG_LEVEL"
	envLogFormat = "LOG_FORMAT"
	// timestamp with millisecond precision
	timestampFormat = "Jan _2 15:04:05.00"
)

// NewLog - configure logging
func NewLog(confName string) *logrus.Entry {
	var conf *koanf.Koanf
	var err error
	if conf, err = NewConf(confName); err != nil {
		conf = koanf.New(PropDelimiter)
	}
	logFormat := GetEnv(envLogFormat, conf.String(envLogFormat))
	if strings.EqualFold(logFormat, jsonLog) {
		logrus.SetFormatter(&logrus.JSONFormatter{
			DisableHTMLEscape: true,
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timestampFormat,
			ForceColors:     true,
		})
	}
	logrus.SetOutput(os.Stdout)
	var logLevel logrus.Level
	if logLevel, err = logrus.ParseLevel(GetEnv(envLogLevel, conf.String(envLogLevel))); err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)

	return logrus.WithFields(logrus.Fields{})
}

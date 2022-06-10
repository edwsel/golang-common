package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func (l *Logger) LevelType() (logrus.Level, error) {
	return logrus.ParseLevel(l.Level)
}

func (l *Logger) FormatType() (logrus.Formatter, error) {
	if l.Format == "text" {
		return &logrus.TextFormatter{}, nil
	} else if l.Format == "json" {
		return &logrus.JSONFormatter{}, nil
	}

	return nil, fmt.Errorf("not a valid logrus Format: %q", l.Format)
}

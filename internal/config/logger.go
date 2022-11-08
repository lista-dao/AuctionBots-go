package config

import (
	"github.com/namsral/flag"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

type logger struct {
	loggerParams `json:"log"`
	log          *logrus.Logger
	o            sync.Once
}

type loggerParams struct {
	level string `json:"level"`
	lvl   logrus.Level
}

func (l *logger) validate() error {
	return errors.Wrap(l.check(), "failed to validate logger")
}

func (l *logger) populate() {
	flag.StringVar(&l.loggerParams.level, "LOG_LEVEL", "debug", "log level")
	flag.Parse()
}

func (l *logger) check() error {
	lvl, err := logrus.ParseLevel(l.level)
	if err != nil {
		return err
	}

	l.lvl = lvl
	return nil
}

func (l *logger) Log() *logrus.Logger {
	l.o.Do(func() {
		if l.log == nil {
			log := logrus.New()
			log.Level = l.lvl
			l.log = log
		}
	})

	return l.log
}

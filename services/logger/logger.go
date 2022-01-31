package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerInstance struct {
	Logger zerolog.Logger
}

func CreateLogger(componentName string) LoggerInstance {
	l := log.With().
		Str("component", componentName).
		Logger()
	return LoggerInstance{Logger: l}
}

func send(event *zerolog.Event, msg string, str []string) {
	for i := 0; i < len(str); i += 2 {
		event.
			Str(str[i], str[i+1])
	}

	event.Msg(msg)
}

func (li LoggerInstance) Info(msg string, str ...string) {
	l := li.Logger.Info()
	send(l, msg, str)
}

func (li LoggerInstance) Trace(msg string, str ...string) {
	l := li.Logger.Trace()
	send(l, msg, str)
}

func (li LoggerInstance) Debug(msg string, str ...string) {
	l := li.Logger.Debug()
	send(l, msg, str)
}

func (li LoggerInstance) Warn(msg string, str ...string) {
	l := li.Logger.Warn()
	send(l, msg, str)
}

func (li LoggerInstance) Error(msg string, err error, str ...string) {
	l := li.Logger.Error().Err(err)
	send(l, msg, str)
}

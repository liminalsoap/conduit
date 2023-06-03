package logger

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

type Logger struct {
	logger *zerolog.Logger
}

func NewLogger(level string) *Logger {
	var l zerolog.Level
	switch strings.ToLower(level) {
	case "debug":
		l = zerolog.DebugLevel
	case "info":
		l = zerolog.InfoLevel
	case "warn":
		l = zerolog.WarnLevel
	case "error":
		l = zerolog.ErrorLevel
	case "fatal":
		l = zerolog.FatalLevel
	}
	zerolog.SetGlobalLevel(l)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &Logger{logger: &logger}
}

func (l Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l Logger) Debugf(msg string, args ...interface{}) {
	l.logger.Debug().Msgf(msg, args...)
}

func (l Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l Logger) Infof(msg string, args ...interface{}) {
	l.logger.Info().Msgf(msg, args...)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l Logger) Warnf(msg string, args ...interface{}) {
	l.logger.Warn().Msgf(msg, args...)
}

func (l Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l Logger) Errorf(msg string, args ...interface{}) {
	l.logger.Error().Msgf(msg, args...)
}

func (l Logger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}

func (l Logger) Fatalf(msg string, args ...interface{}) {
	l.logger.Fatal().Msgf(msg, args...)
}

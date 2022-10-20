package cuslog

import (
	"io"
	"os"
)

const FmtEmptySeparate = ""

type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

type options struct {
	output        io.Writer
	level         Level
	stdLevel      Level
	formatter     Formatter
	disableCaller bool
}

type Option func(*options)

func initOptions(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}
	if o.output == nil {
		o.output = os.Stderr
	}
	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}
	return
}

func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdLevel = level
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

func WithDisableCaller(caller bool) Option {
	return func(o *options) {
		o.disableCaller = caller
	}
}

package cuslog

import (
	"fmt"
	"os"
	"sync"
)

var std = New()

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

func New(opt ...Option) *logger {
	logger := &logger{opt: initOptions(opt...)}
	logger.entryPool = &sync.Pool{New: func() any { return entry(logger) }}
	return logger
}

func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *logger) Debug(args ...any) {
	l.entry().write(DebugLevel, FmtEmptySeparate, args...)
}

func (l *logger) Info(args ...any) {
	l.entry().write(InfoLevel, FmtEmptySeparate, args...)
}

func (l *logger) Warn(args ...any) {
	l.entry().write(WarnLevel, FmtEmptySeparate, args...)
}

func (l *logger) Error(args ...any) {
	l.entry().write(ErrorLevel, FmtEmptySeparate, args...)
}

func (l *logger) Panic(args ...any) {
	l.entry().write(PanicLevel, FmtEmptySeparate, args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Fatal(args ...any) {
	l.entry().write(FatalLevel, FmtEmptySeparate, args...)
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.entry().write(DebugLevel, format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.entry().write(InfoLevel, format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.entry().write(WarnLevel, format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.entry().write(ErrorLevel, format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.entry().write(PanicLevel, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.entry().write(FatalLevel, format, args...)
	os.Exit(1)
}

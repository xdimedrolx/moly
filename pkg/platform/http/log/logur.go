package log

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"logur.dev/logur"
	"os"
)

type LogurAdapter struct {
	logur logur.Logger
}

func NewLogurAdapter(logger logur.LoggerFacade) *LogurAdapter {
	return &LogurAdapter{logur: logger}
}

func (l *LogurAdapter) Output() io.Writer {
	return os.Stdout
}

func (l *LogurAdapter) SetOutput(w io.Writer) {

}

// Log format is completely determined by logur
func (l *LogurAdapter) Prefix() string {
	return ""
}

// Log format is completely determined by logur
func (l *LogurAdapter) SetPrefix(p string) {

}

func (l *LogurAdapter) Level() log.Lvl {
	// Fallback to Info level
	return log.INFO
}

// Ignore level as Logger shouldn't be configured by echo.Echo anyway
func (l *LogurAdapter) SetLevel(v log.Lvl) {
}

// Log format is completely determined by logur
func (l *LogurAdapter) SetHeader(h string) {

}

func (l *LogurAdapter) Print(i ...interface{}) {
	l.logur.Debug(fmt.Sprint(i...))
}

func (l *LogurAdapter) Printf(format string, args ...interface{}) {
	l.logur.Info(fmt.Sprintf(format, args...))
}

func (l *LogurAdapter) Printj(j log.JSON) {
	l.logur.Info("", j)
}

func (l *LogurAdapter) Debug(i ...interface{}) {
	l.logur.Debug(fmt.Sprint(i...))
}

func (l *LogurAdapter) Debugf(format string, args ...interface{}) {
	l.logur.Debug(fmt.Sprintf(format, args...))
}

func (l *LogurAdapter) Debugj(j log.JSON) {
	l.logur.Debug("", j)
}

func (l *LogurAdapter) Info(i ...interface{}) {
	l.logur.Info(fmt.Sprint(i...))
}

func (l *LogurAdapter) Infof(format string, args ...interface{}) {
	l.logur.Info(fmt.Sprintf(format, args...))
}

func (l *LogurAdapter) Infoj(j log.JSON) {
	l.logur.Info("", j)
}

func (l *LogurAdapter) Warn(i ...interface{}) {
	l.logur.Warn(fmt.Sprint(i...))
}

func (l *LogurAdapter) Warnf(format string, args ...interface{}) {
	l.logur.Error(fmt.Sprintf(format, args...))
}

func (l *LogurAdapter) Warnj(j log.JSON) {
	l.logur.Warn("", j)
}

func (l *LogurAdapter) Error(i ...interface{}) {
	l.logur.Error(fmt.Sprint(i...))
}

func (l *LogurAdapter) Errorf(format string, args ...interface{}) {
	l.logur.Error(fmt.Sprintf(format, args...))
}

func (l *LogurAdapter) Errorj(j log.JSON) {
	l.logur.Error("", j)
}

func (l *LogurAdapter) Fatal(i ...interface{}) {
	l.logur.Error(fmt.Sprint(i...))
	os.Exit(1)
}

func (l *LogurAdapter) Fatalj(j log.JSON) {
	l.logur.Error("", j)
	os.Exit(1)
}

func (l *LogurAdapter) Fatalf(format string, args ...interface{}) {
	l.logur.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (l *LogurAdapter) Panic(i ...interface{}) {
	l.logur.Error(fmt.Sprint(i...))
	panic(fmt.Sprint(i...))
}

func (l *LogurAdapter) Panicj(j log.JSON) {
	l.logur.Error("", j)
	panic(j)
}

func (l *LogurAdapter) Panicf(format string, args ...interface{}) {
	l.logur.Error(fmt.Sprintf(format, args...))
	panic(fmt.Sprintf(format, args...))
}

var _ echo.Logger = (*LogurAdapter)(nil)

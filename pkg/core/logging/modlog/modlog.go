/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package modlog

import (
	"github.com/hellobchain/fabric-sdk-go/pkg/core/logging/api"
	"github.com/hellobchain/fabric-sdk-go/pkg/core/logging/metadata"
	"github.com/hellobchain/wswlog/wlogging"
	"sync"
)

var rwmutex = &sync.RWMutex{}
var moduleLevels = &metadata.ModuleLevels{}
var callerInfos = &metadata.CallerInfo{}

// Provider is the default logger implementation
type Provider struct {
}

//GetLogger returns SDK logger implementation
func (p *Provider) GetLogger(module string) api.Logger {
	newDefLogger := wlogging.MustGetLogger(module)
	return &Log{deflogger: newDefLogger, module: module}
}

//LoggerProvider returns logging provider for SDK logger
func LoggerProvider() api.LoggerProvider {
	return &Provider{}
}

//Log is a standard SDK logger implementation
type Log struct {
	deflogger    *wlogging.WswLogger
	customLogger api.Logger
	module       string
	custom       bool
	once         sync.Once
}

//LoggerOpts  for all logger customization options
type loggerOpts struct {
	levelEnabled      bool
	callerInfoEnabled bool
}

//getLoggerOpts - returns LoggerOpts which can be used for customization
func getLoggerOpts(module string, level api.Level) *loggerOpts {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return &loggerOpts{
		levelEnabled:      moduleLevels.IsEnabledFor(module, level),
		callerInfoEnabled: callerInfos.IsCallerInfoEnabled(module, level),
	}
}

// Fatal is CRITICAL log followed by a call to os.Exit(1).
func (l *Log) Fatal(args ...interface{}) {
	l.deflogger.Fatal(args...)
}

// Fatalf is CRITICAL log formatted followed by a call to os.Exit(1).
func (l *Log) Fatalf(format string, args ...interface{}) {
	l.deflogger.Fatalf(format, args...)
}

// Panic is CRITICAL log followed by a call to panic()
func (l *Log) Panic(args ...interface{}) {
	l.deflogger.Panic(args...)
}

// Panicf is CRITICAL log formatted followed by a call to panic()
func (l *Log) Panicf(format string, args ...interface{}) {
	l.deflogger.Panicf(format, args...)
}

// Debug calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *Log) Debug(args ...interface{}) {
	l.deflogger.Debug(args...)
}

// Debugf calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *Log) Debugf(format string, args ...interface{}) {
	l.deflogger.Debugf(format, args...)
}

// Info calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *Log) Info(args ...interface{}) {
	l.deflogger.Info(args...)
}

// Infof calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *Log) Infof(format string, args ...interface{}) {

	l.deflogger.Infof(format, args...)
}

// Warn calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *Log) Warn(args ...interface{}) {
	l.deflogger.Warn(args...)
}

// Warnf calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *Log) Warnf(format string, args ...interface{}) {
	l.deflogger.Warnf(format, args...)
}

// Error calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *Log) Error(args ...interface{}) {
	l.deflogger.Error(args...)
}

// Errorf calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *Log) Errorf(format string, args ...interface{}) {
	l.deflogger.Errorf(format, args...)
}

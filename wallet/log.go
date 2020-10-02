// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2018 The Hdfchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import "github.com/hdfchain/slog"

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log = slog.Disabled

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using slog.
func UseLogger(logger slog.Logger) {
	log = logger
}

type infoLogger struct{}

var infoLog infoLogger

func (infoLogger) Print(args ...interface{})                 { log.Info(args...) }
func (infoLogger) Printf(format string, args ...interface{}) { log.Infof(format, args...) }

type debugLogger struct{}

var debugLog debugLogger

func (debugLogger) Print(args ...interface{})                 { log.Debug(args...) }
func (debugLogger) Printf(format string, args ...interface{}) { log.Debugf(format, args...) }

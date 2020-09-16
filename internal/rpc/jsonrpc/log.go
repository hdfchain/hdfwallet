// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2018 The Hdfchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package jsonrpc

import "github.com/hdfchain/slog"

var log = slog.Disabled

// UseLogger sets the package-wide logger.  Any calls to this function must be
// made before a server is created and used (it is not concurrent safe).
func UseLogger(logger slog.Logger) {
	log = logger
}

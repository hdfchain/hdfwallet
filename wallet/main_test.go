// Copyright (c) 2020 The Hdfchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"flag"
	"os"
	"testing"

	"github.com/hdfchain/slog"
)

var logFlag = flag.Bool("log", false, "enable package logger")

func TestMain(m *testing.M) {
	flag.Parse()
	if *logFlag {
		UseLogger(slog.NewBackend(os.Stderr).Logger("WLLT"))
	}

	os.Exit(m.Run())
}

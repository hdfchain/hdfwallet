// Copyright (c) 2018 The Hdfchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/hdfchain/hdfwallet/wallet/drivers/bdb"
	"github.com/hdfchain/hdfwallet/wallet/walletdb"
	"github.com/hdfchain/hdfd/chaincfg/v3"
	"github.com/hdfchain/hdfd/dcrutil/v3"
)

var basicWalletConfig = Config{
	PubPassphrase: []byte(InsecurePubPassphrase),
	GapLimit:      20,
	RelayFee:      dcrutil.Amount(1e5),
	Params:        chaincfg.SimNetParams(),
}

func testWallet(t *testing.T, cfg *Config) (w *Wallet, teardown func()) {
	ctx := context.Background()
	f, err := ioutil.TempFile("", "hdfwallet.testdb")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	db, err := walletdb.Create("bdb", f.Name())
	if err != nil {
		t.Fatal(err)
	}
	rm := func() {
		db.Close()
		os.Remove(f.Name())
	}
	err = Create(ctx, opaqueDB{db}, []byte(InsecurePubPassphrase), []byte("private"), nil, cfg.Params)
	if err != nil {
		rm()
		t.Fatal(err)
	}
	cfg.DB = opaqueDB{db}
	w, err = Open(ctx, cfg)
	if err != nil {
		rm()
		t.Fatal(err)
	}
	teardown = func() {
		rm()
	}
	return
}

package compat

import (
	dcrec1 "github.com/hdfchain/hdfd/dcrec"
	dcrutil2 "github.com/hdfchain/hdfd/dcrutil/v2"
	hdkeychain2 "github.com/hdfchain/hdfd/hdkeychain/v2"
)

func HD2Address(k *hdkeychain2.ExtendedKey, params dcrutil2.AddressParams) (*dcrutil2.AddressPubKeyHash, error) {
	pk, err := k.ECPubKey()
	if err != nil {
		return nil, err
	}
	hash := dcrutil2.Hash160(pk.SerializeCompressed())
	return dcrutil2.NewAddressPubKeyHash(hash, params, dcrec1.STEcdsaSecp256k1)
}

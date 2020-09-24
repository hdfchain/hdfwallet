package compat

import (
	"github.com/hdfchain/hdfd/dcrec"
	"github.com/hdfchain/hdfd/dcrutil/v3"
	"github.com/hdfchain/hdfd/hdkeychain/v3"
)

func HD2Address(k *hdkeychain.ExtendedKey, params dcrutil.AddressParams) (*dcrutil.AddressPubKeyHash, error) {
	pk := k.SerializedPubKey()
	hash := dcrutil.Hash160(pk)
	return dcrutil.NewAddressPubKeyHash(hash, params, dcrec.STEcdsaSecp256k1)
}

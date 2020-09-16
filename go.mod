module github.com/hdfchain/hdfwallet

go 1.12

require (
	github.com/hdfchain/base58 v1.0.5 // indirect
	github.com/hdfchain/hdfd/addrmgr v1.1.0
	github.com/hdfchain/hdfd/blockchain/stake v1.2.1
	github.com/hdfchain/hdfd/blockchain/stake/v2 v2.0.2
	github.com/hdfchain/hdfd/blockchain/standalone v1.1.0
	github.com/hdfchain/hdfd/certgen v1.1.0
	github.com/hdfchain/hdfd/chaincfg v1.5.2
	github.com/hdfchain/hdfd/chaincfg/chainhash v1.0.2
	github.com/hdfchain/hdfd/chaincfg/v2 v2.3.0
	github.com/hdfchain/hdfd/connmgr v1.1.0
	github.com/hdfchain/hdfd/connmgr/v2 v2.1.0
	github.com/hdfchain/hdfd/dcrec v1.0.0
	github.com/hdfchain/hdfd/dcrec/secp256k1 v1.0.3 // indirect
	github.com/hdfchain/hdfd/dcrjson/v2 v2.2.0
	github.com/hdfchain/hdfd/dcrjson/v3 v3.0.1
	github.com/hdfchain/hdfd/dcrutil v1.4.0
	github.com/hdfchain/hdfd/dcrutil/v2 v2.0.1
	github.com/hdfchain/hdfd/hdkeychain/v2 v2.1.0
	github.com/hdfchain/hdfd/rpc/jsonrpc/types v1.0.1
	github.com/hdfchain/hdfd/rpcclient/v2 v2.1.0
	github.com/hdfchain/hdfd/txscript v1.1.0
	github.com/hdfchain/hdfd/txscript/v2 v2.1.0
	github.com/hdfchain/hdfd/wire v1.3.0
	github.com/hdfchain/hdfwallet/chain/v3 v3.0.1
	github.com/hdfchain/hdfwallet/errors/v2 v2.0.0
	github.com/hdfchain/hdfwallet/p2p/v2 v2.0.0
	github.com/hdfchain/hdfwallet/rpc/client/hdfd v1.0.0
	github.com/hdfchain/hdfwallet/rpc/jsonrpc/types v1.4.0
	github.com/hdfchain/hdfwallet/rpc/walletrpc v0.3.0
	github.com/hdfchain/hdfwallet/spv/v3 v3.0.1
	github.com/hdfchain/hdfwallet/ticketbuyer/v4 v4.0.1
	github.com/hdfchain/hdfwallet/version v1.0.3
	github.com/hdfchain/hdfwallet/wallet/v3 v3.1.0
	github.com/hdfchain/hdfwallet/walletseed v1.0.3
	github.com/hdfchain/go-socks v1.1.0
	github.com/hdfchain/slog v1.0.0
	github.com/gorilla/websocket v1.4.1
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/logrotate v1.0.0
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.25.1
)

replace (
	github.com/hdfchain/hdfwallet/errors/v2 => ./errors
	github.com/hdfchain/hdfwallet/lru => ./lru
	github.com/hdfchain/hdfwallet/p2p/v2 => ./p2p
	github.com/hdfchain/hdfwallet/pgpwordlist => ./pgpwordlist
	github.com/hdfchain/hdfwallet/rpc/jsonrpc/types => ./rpc/jsonrpc/types
	github.com/hdfchain/hdfwallet/rpc/walletrpc => ./rpc/walletrpc
	github.com/hdfchain/hdfwallet/validate => ./validate
	github.com/hdfchain/hdfwallet/version => ./version
	github.com/hdfchain/hdfwallet/wallet/v3 => ./wallet
	github.com/hdfchain/hdfwallet/walletseed => ./walletseed
)

replace github.com/hdfchain/hdfwallet/spv/v3 => ./spv

replace github.com/hdfchain/hdfwallet/ticketbuyer/v4 => ./ticketbuyer

replace github.com/hdfchain/hdfwallet/chain/v3 => ./chain

replace github.com/hdfchain/hdfwallet/rpc/client/hdfd => ./rpc/client/hdfd

replace github.com/hdfchain/hdfwallet/deployments/v2 => ./deployments

package command

import (
	"github.com/umbracle/ethgo"

	"github.com/nnlgsakib/neth/chain"
	"github.com/nnlgsakib/neth/server"
)

const (
	DefaultGenesisFileName  = "genesis.json"
	DefaultChainName        = "neth-smart-chain"
	DefaultChainID          = 9996
	DefaultConsensus        = server.IBFTConsensus
	DefaultGenesisGasUsed   = 458752  // 0x70000
	DefaultGenesisGasLimit  = 5242880 // 0x500000
	DefaultGenesisBaseFeeEM = chain.GenesisBaseFeeEM
)

var (
	DefaultStake          = ethgo.Ether(1e6)
	DefaultPremineBalance = ethgo.Ether(1e6)
	DefaultGenesisBaseFee = chain.GenesisBaseFee
)

const (
	JSONOutputFlag  = "json"
	GRPCAddressFlag = "grpc-address"
	JSONRPCFlag     = "jsonrpc"
)

// GRPCAddressFlagLEGACY Legacy flag that needs to be present to preserve backwards
// compatibility with running clients
const (
	GRPCAddressFlagLEGACY = "grpc"
)

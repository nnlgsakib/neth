package reward

import (
	"fmt"

	"github.com/nnlgsakib/neth/chain"
	"github.com/nnlgsakib/neth/helper/common"
	"github.com/nnlgsakib/neth/helper/hex"
	"github.com/nnlgsakib/neth/types"
)

var (
	DefaultRewardContractBalance = "0x21E19E0C9BAB2400000" // Balance to pre-fund the reward contract
	RewardContractAddress        = types.StringToAddress("0x000000000000000000000000000000000000FFf1")
	//nolint: lll
	RewardContractBytecode = "0x608060405273c6bf9487f53dee994210b7844636685b9c75bba76000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073c6bf9487f53dee994210b7844636685b9c75bba7600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801560b857600080fd5b506111b2806100c86000396000f3fe6080604052600436106100e15760003560e01c8063b7bd5b371161007f578063d0f2b54c11610059578063d0f2b54c14610332578063f2fde38b1461035d578063f8077fae14610386578063fe188184146103b157610136565b8063b7bd5b37146102b1578063bafedcaa146102dc578063cd74234e1461030757610136565b806348b92fd3116100bb57806348b92fd3146102055780636798ec2d14610232578063893d20e81461025b5780638da5cb5b1461028657610136565b8063060a726c146101865780630728ab88146101b15780631ec8bb8c146101dc57610136565b36610136573373ffffffffffffffffffffffffffffffffffffffff167fbb527541f6cc244ae00ab689f8c23d072a30a3b4176021df62ea1c2bac8aa2263460405161012c9190610b78565b60405180910390a2005b3373ffffffffffffffffffffffffffffffffffffffff167fbb527541f6cc244ae00ab689f8c23d072a30a3b4176021df62ea1c2bac8aa2263460405161017c9190610b78565b60405180910390a2005b34801561019257600080fd5b5061019b6103dc565b6040516101a89190610bd4565b60405180910390f35b3480156101bd57600080fd5b506101c6610402565b6040516101d39190610b78565b60405180910390f35b3480156101e857600080fd5b5061020360048036038101906101fe9190610c4c565b61040c565b005b34801561021157600080fd5b5061021a6106f8565b60405161022993929190610c8c565b60405180910390f35b34801561023e57600080fd5b5061025960048036038101906102549190610cc3565b610731565b005b34801561026757600080fd5b506102706108ee565b60405161027d9190610bd4565b60405180910390f35b34801561029257600080fd5b5061029b610917565b6040516102a89190610bd4565b60405180910390f35b3480156102bd57600080fd5b506102c661093b565b6040516102d39190610b78565b60405180910390f35b3480156102e857600080fd5b506102f1610941565b6040516102fe9190610b78565b60405180910390f35b34801561031357600080fd5b5061031c610947565b6040516103299190610b78565b60405180910390f35b34801561033e57600080fd5b5061034761094f565b6040516103549190610bd4565b60405180910390f35b34801561036957600080fd5b50610384600480360381019061037f9190610cc3565b610979565b005b34801561039257600080fd5b5061039b610b33565b6040516103a89190610b78565b60405180910390f35b3480156103bd57600080fd5b506103c6610b39565b6040516103d39190610bd4565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600254905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461049c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049390610d73565b60405180910390fd5b600081116104df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d690610e05565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361054e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054590610e71565b60405180910390fd5b80471015610591576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058890610edd565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff16826040516105b790610f2e565b60006040518083038185875af1925050503d80600081146105f4576040519150601f19603f3d011682016040523d82523d6000602084013e6105f9565b606091505b505090508061063d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063490610f8f565b60405180910390fd5b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160048190555042600581905550816002600082825461069e9190610fde565b925050819055508273ffffffffffffffffffffffffffffffffffffffff167fe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea6836040516106eb9190610b78565b60405180910390a2505050565b6000806000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600454600554925092509250909192565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b690611084565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361082e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610825906110f0565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fbfe69a6bb4ed6720331fd9ad1df9b69784159f6aa02da24706b528506ca659fe60405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b60045481565b600047905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a07576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109fe90611084565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6d9061115c565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60055481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000819050919050565b610b7281610b5f565b82525050565b6000602082019050610b8d6000830184610b69565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610bbe82610b93565b9050919050565b610bce81610bb3565b82525050565b6000602082019050610be96000830184610bc5565b92915050565b600080fd5b610bfd81610bb3565b8114610c0857600080fd5b50565b600081359050610c1a81610bf4565b92915050565b610c2981610b5f565b8114610c3457600080fd5b50565b600081359050610c4681610c20565b92915050565b60008060408385031215610c6357610c62610bef565b5b6000610c7185828601610c0b565b9250506020610c8285828601610c37565b9150509250929050565b6000606082019050610ca16000830186610bc5565b610cae6020830185610b69565b610cbb6040830184610b69565b949350505050565b600060208284031215610cd957610cd8610bef565b5b6000610ce784828501610c0b565b91505092915050565b600082825260208201905092915050565b7f4e6f7420617574686f72697a65643a204f6e6c79206f66662d636861696e206560008201527f6e67696e65000000000000000000000000000000000000000000000000000000602082015250565b6000610d5d602583610cf0565b9150610d6882610d01565b604082019050919050565b60006020820190508181036000830152610d8c81610d50565b9050919050565b7f52657761726420616d6f756e74206d757374206265206772656174657220746860008201527f616e207a65726f00000000000000000000000000000000000000000000000000602082015250565b6000610def602783610cf0565b9150610dfa82610d93565b604082019050919050565b60006020820190508181036000830152610e1e81610de2565b9050919050565b7f496e76616c696420726563697069656e74206164647265737300000000000000600082015250565b6000610e5b601983610cf0565b9150610e6682610e25565b602082019050919050565b60006020820190508181036000830152610e8a81610e4e565b9050919050565b7f496e73756666696369656e7420636f6e74726163742062616c616e6365000000600082015250565b6000610ec7601d83610cf0565b9150610ed282610e91565b602082019050919050565b60006020820190508181036000830152610ef681610eba565b9050919050565b600081905092915050565b50565b6000610f18600083610efd565b9150610f2382610f08565b600082019050919050565b6000610f3982610f0b565b9150819050919050565b7f526577617264207472616e73666572206661696c656400000000000000000000600082015250565b6000610f79601683610cf0565b9150610f8482610f43565b602082019050919050565b60006020820190508181036000830152610fa881610f6c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610fe982610b5f565b9150610ff483610b5f565b925082820190508082111561100c5761100b610faf565b5b92915050565b7f4f6e6c79206f776e65722063616e20706572666f726d2074686973206163746960008201527f6f6e000000000000000000000000000000000000000000000000000000000000602082015250565b600061106e602283610cf0565b915061107982611012565b604082019050919050565b6000602082019050818103600083015261109d81611061565b9050919050565b7f496e76616c6964206e657720656e67696e652061646472657373000000000000600082015250565b60006110da601a83610cf0565b91506110e5826110a4565b602082019050919050565b60006020820190508181036000830152611109816110cd565b9050919050565b7f496e76616c6964206e6577206f776e6572206164647265737300000000000000600082015250565b6000611146601983610cf0565b915061115182611110565b602082019050919050565b6000602082019050818103600083015261117581611139565b905091905056fea26469706673582212205d475c68f22233dee2a7eadcd975e5241a290ba8d6393c5c5b1a81bcac0b319964736f6c634300081a0033" // Truncated for brevity
)

// PredeployRewardContract predeploys the reward contract in the genesis block
func PredeployRewardContract() (*chain.GenesisAccount, error) {
	// Decode the bytecode
	scHex, err := hex.DecodeHex(RewardContractBytecode)
	if err != nil {
		return nil, fmt.Errorf("unable to decode reward contract bytecode: %w", err)
	}

	// Parse the default balance into a *big.Int
	bigDefaultBalance, err := common.ParseUint256orHex(&DefaultRewardContractBalance)
	if err != nil {
		return nil, fmt.Errorf("unable to parse reward contract balance: %w", err)
	}

	// Create the genesis account for the reward contract
	rewardAccount := &chain.GenesisAccount{
		Code:    scHex,
		Balance: bigDefaultBalance,
	}

	// Return the pre-configured reward contract account
	return rewardAccount, nil
}

package reward 

import "github.com/nnlgsakib/neth/types"




var (
	DefaultRewardContractBalance = "0x8AC7230489E80000" 
     Reward_contract = types.StringToAddress("ffff")
	 //nolint: lll
	 Reward_Bytecode = "0x608060405273c6bf9487f53dee994210b7844636685b9c75bba76000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073c6bf9487f53dee994210b7844636685b9c75bba7600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801560b857600080fd5b50610d58806100c86000396000f3fe6080604052600436106100c65760003560e01c80638da5cb5b1161007f578063cd74234e11610059578063cd74234e14610253578063d0f2b54c1461027e578063f8077fae146102a9578063fe188184146102d4576100cd565b80638da5cb5b146101d2578063b7bd5b37146101fd578063bafedcaa14610228576100cd565b8063060a726c146100d25780630728ab88146100fd5780631ec8bb8c1461012857806348b92fd3146101515780636798ec2d1461017e578063893d20e8146101a7576100cd565b366100cd57005b600080fd5b3480156100de57600080fd5b506100e76102ff565b6040516100f4919061081e565b60405180910390f35b34801561010957600080fd5b50610112610325565b60405161011f9190610852565b60405180910390f35b34801561013457600080fd5b5061014f600480360381019061014a91906108ca565b61032f565b005b34801561015d57600080fd5b5061016661061b565b6040516101759392919061090a565b60405180910390f35b34801561018a57600080fd5b506101a560048036038101906101a09190610941565b610654565b005b3480156101b357600080fd5b506101bc610726565b6040516101c9919061081e565b60405180910390f35b3480156101de57600080fd5b506101e761074f565b6040516101f4919061081e565b60405180910390f35b34801561020957600080fd5b50610212610773565b60405161021f9190610852565b60405180910390f35b34801561023457600080fd5b5061023d610779565b60405161024a9190610852565b60405180910390f35b34801561025f57600080fd5b5061026861077f565b6040516102759190610852565b60405180910390f35b34801561028a57600080fd5b50610293610787565b6040516102a0919061081e565b60405180910390f35b3480156102b557600080fd5b506102be6107b1565b6040516102cb9190610852565b60405180910390f35b3480156102e057600080fd5b506102e96107b7565b6040516102f6919061081e565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600254905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b6906109f1565b60405180910390fd5b60008111610402576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f990610a83565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046890610aef565b60405180910390fd5b804710156104b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ab90610b5b565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff16826040516104da90610bac565b60006040518083038185875af1925050503d8060008114610517576040519150601f19603f3d011682016040523d82523d6000602084013e61051c565b606091505b5050905080610560576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055790610c0d565b60405180910390fd5b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816004819055504260058190555081600260008282546105c19190610c5c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff167fe34918ff1c7084970068b53fd71ad6d8b04e9f15d3886cbf006443e6cdc52ea68360405161060e9190610852565b60405180910390a2505050565b6000806000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600454600554925092509250909192565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d990610d02565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b60045481565b600047905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60055481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610808826107dd565b9050919050565b610818816107fd565b82525050565b6000602082019050610833600083018461080f565b92915050565b6000819050919050565b61084c81610839565b82525050565b60006020820190506108676000830184610843565b92915050565b600080fd5b61087b816107fd565b811461088657600080fd5b50565b60008135905061089881610872565b92915050565b6108a781610839565b81146108b257600080fd5b50565b6000813590506108c48161089e565b92915050565b600080604083850312156108e1576108e061086d565b5b60006108ef85828601610889565b9250506020610900858286016108b5565b9150509250929050565b600060608201905061091f600083018661080f565b61092c6020830185610843565b6109396040830184610843565b949350505050565b6000602082840312156109575761095661086d565b5b600061096584828501610889565b91505092915050565b600082825260208201905092915050565b7f4e6f7420617574686f72697a65643a204f6e6c79206f66662d636861696e206560008201527f6e67696e65000000000000000000000000000000000000000000000000000000602082015250565b60006109db60258361096e565b91506109e68261097f565b604082019050919050565b60006020820190508181036000830152610a0a816109ce565b9050919050565b7f52657761726420616d6f756e74206d757374206265206772656174657220746860008201527f616e207a65726f00000000000000000000000000000000000000000000000000602082015250565b6000610a6d60278361096e565b9150610a7882610a11565b604082019050919050565b60006020820190508181036000830152610a9c81610a60565b9050919050565b7f496e76616c696420726563697069656e74206164647265737300000000000000600082015250565b6000610ad960198361096e565b9150610ae482610aa3565b602082019050919050565b60006020820190508181036000830152610b0881610acc565b9050919050565b7f496e73756666696369656e7420636f6e74726163742062616c616e6365000000600082015250565b6000610b45601d8361096e565b9150610b5082610b0f565b602082019050919050565b60006020820190508181036000830152610b7481610b38565b9050919050565b600081905092915050565b50565b6000610b96600083610b7b565b9150610ba182610b86565b600082019050919050565b6000610bb782610b89565b9150819050919050565b7f526577617264207472616e73666572206661696c656400000000000000000000600082015250565b6000610bf760168361096e565b9150610c0282610bc1565b602082019050919050565b60006020820190508181036000830152610c2681610bea565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c6782610839565b9150610c7283610839565b9250828201905080821115610c8a57610c89610c2d565b5b92915050565b7f4f6e6c79206f776e65722063616e2075706461746520746865206f66662d636860008201527f61696e20656e67696e6500000000000000000000000000000000000000000000602082015250565b6000610cec602a8361096e565b9150610cf782610c90565b604082019050919050565b60006020820190508181036000830152610d1b81610cdf565b905091905056fea2646970667358221220a4ee7153c464067482a687359fad4956197e6952bf1b6bb6fc72bc8e89a82a0664736f6c634300081a0033"
)
package main

import (
	"fmt"

	"github.com/skycoin/skycoin-exchange/src/coin/bitcoin"
)

func main() {

	//pubkey, seckey := cipher.GenerateKeyPair()

	//pubkeyAddr := cipher.AddressFromPubKey(pubkey)
	//seckeyAddr := cipher.AddressFromSecKey(seckey)

	pubkeyAddr := "1AHWU3jY8RL2ERBNBrMaFPo5u6NSY3v5Sf"

	//fmt.Println(pubkeyAddr, seckeyAddr)

	holdings, err := bitcoin.GetBalance([]string{pubkeyAddr})
	if err != nil {
		panic(err)
	}
	fmt.Println(holdings)

	type Output struct {
		pubAddr    string
		outputHash string
		satoshiBal uint64
	}

	outputs, err := bitcoin.GetUnspentOutputs([]string{pubkeyAddr})
	if err != nil {
		panic(err)
	}

	for _, output := range outputs {
		o := Output{
			pubAddr:    output.GetAddress(),
			outputHash: output.GetTxid(),
			satoshiBal: output.GetAmount(),
		}
		fmt.Println(o)

	}

}

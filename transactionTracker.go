package main

import (
	"fmt"
	"time"
	"github.com/skycoin/skycoin/src/api/cli"
	"github.com/skycoin/skycoin-exchange/src/coin/bitcoin"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
)

//create state struct 
//create array/struct of pending things to send 
//create a single transaction
//DO NOT CREATE new transaction until previous transactoin executed
//then once transaction is executed, do next pending transaction

type PendingTx struct {
	SkycoinAddress string
	Amount         uint32
}

type State struct {
	PendingOut []PendingTx

	PendingTransaction *coin.Transaction
}

//add pending, send request to queue
func (self *State) AddPendingTx(SkycoinAddress string, uint64 amount) error {
	
	err := 

	var tx PendingTx
	tx.SkycoinAddress = SkycoinAddress
	tx.Amount = amount

	self.State.PendingOut = append(self.State.PendingOut)
}

// all stuff happen in this loop
func (self *State) Tick() {
	//check the state of the pending transaction

	//if non pending transaction or pending transaction is "confirmed", then create next transaction
	if PendingTx == nil && len(self.PendingOut) == 0 {
		return
	}

	if PendingTx == nil && len(self.PendingOut != 0) {
		//turn pendingOut, into a new transaction
	}

	//check status of pending transaction
	//- if transaction is confirmed, then do new tranasction
	//- if transactoin is not confirmed, wait
}

//

func main() {

	for true {

		//x.Tick()
		time.Sleep(5 * time.Millisecond)
	}

	const exRate = 500 //BTC input * 500 = skycoin owed

	skyPubkey, skySeckey := cipher.GenerateKeyPair()
	skyAddr := cipher.AddressFromPubKey(skyPubkey)

	fmt.Println(skySeckey)

	err := skyAddr.Verify(skyPubkey)
	if err != nil {
		panic(err)
	}

	fmt.Println(skyPubkey, skyAddr)

	pubkey, seckey := cipher.GenerateKeyPair()

	pubkeyAddr := cipher.AddressFromPubKey(pubkey)
	seckeyAddr := cipher.AddressFromSecKey(seckey)

	fmt.Println(seckeyAddr)

	errA := pubkeyAddr.Verify(pubkey)
	if errA != nil {
		panic(err)
	}

	boundAddr := map[cipher.Address]cipher.Address{
		skyAddr: pubkeyAddr,
	}

	fmt.Println(boundAddr)

	holdings, err := bitcoin.GetBalance([]string{pubkeyAddr.BitcoinString()})
	if err != nil {
		panic(err)
	}

	holdings = 2

	type Output struct {
		pubAddr    string
		outputHash string
		satoshiBal uint64
	}

	if holdings == 0 {
		fmt.Println("Nothing received yet")
	} else {

		fmt.Println(holdings*exRate, " Skycoin is owed to address ", pubkeyAddr)
		outputs, err := bitcoin.GetUnspentOutputs([]string{pubkeyAddr.BitcoinString()})
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

	//trackTrans := coin.Transaction {
	//	Length: ,
	//	Type: ,
	//	InnerHash ,
	//
	//		Sigs: ,
	//		In: ,
	//		Out: ,
	//	}

	//fmt.Println(trackTrans)
}

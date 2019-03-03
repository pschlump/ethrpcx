package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/pschlump/ethrpc"
)

func main() {
	client := ethrpc.NewEthRPC("ws://192.168.0.158:8546") // will not work
	client := ethrpc.NewEthRPC("http://192.168.0.158:8545")

	version, err := client.Web3ClientVersion()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected, version = %s\n", version)

	// Send 1 eth
	txid, err := client.EthSendTransaction(ethrpc.T{
		From:  "0x6247cf0412c6462da2a51d05139e2a3c6c630f0a",
		To:    "0xcfa202c4268749fbb5136f2b68f7402984ed444b",
		Value: big.NewInt(1000000000000000000),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TxID is %s\n", txid)
}

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/eris-ltd/mint-client/mintx/core"
	"github.com/tendermint/tendermint/wire"
)

//TODO pass amt in from header
func UpdateNameReg(fileName, hash, amt string) (string, error) {
	//build a mintx, prepare it for POST request

	nodeAddr := os.Getenv("MINTX_NODE_ADDR")
	//signAddr := os.Getenv("MINTX_SIGN_ADDR")
	//chainID := os.Getenv("MINTX_CHAINID")
	pubkey := os.Getenv("MINTX_PUBKEY")
	addr := ""
	amtS := "100000000"
	nonceS := ""
	feeS := "0"
	name := "bombina"
	data := "firebelly"

	tx, err := core.Name(nodeAddr, pubkey, addr, amtS, nonceS, feeS, name, data)
	if err != nil {
		fmt.Printf("corename error: %v\n", err)
	}

	fmt.Printf("tx: %v\n", tx)
	txData := wire.BinaryBytes(tx)
	fmt.Printf("txdata: %v\n", txData)

	err = postTxData(nodeAddr, hash, txData)
	if err != nil {
		fmt.Printf("post error: %v\n", err)
	}

	return "", nil
}

func postTxData(nodeAddr, hash string, txData []byte) error {
	//ping node
	//send tx in body
	//hash in header -> or can be

	// post needs to be to toadserver endpoint, which'll route the TX to the node
	// this post request needs to ensure node is running; query its endpoint
	// or should the toadserver have an endpoint that does that?
	txD := bytes.NewReader(txData)
	//it can also query for th name reg to ensure things are good
	//TODO use NewRequest & DEfaultClient.Do to set headers
	endpoint := "0.0.0.0:11113" + "receiveNameTx/" + hash
	_, err := http.Post(endpoint, "", txD)
	if err != nil {
		fmt.Printf("post error: %v\n", err)
	}
	return nil
}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/eris-ltd/mint-client/mintx/core"
	"github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/types"
	"github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/wire"
)

//TODO pass amt in from url. (make estimate of length)
func UpdateNameReg(fileName, hash, amt string) (string, error) {
	nodeAddr := os.Getenv("MINTX_NODE_ADDR")
	signAddr := os.Getenv("MINTX_SIGN_ADDR")
	chainID := os.Getenv("MINTX_CHAINID")
	pubkey := os.Getenv("MINTX_PUBKEY")
	addr := ""
	amtS := "1000000"
	nonceS := ""
	feeS := "0"
	name := fileName
	data := hash

	//build and sign a nameTx, send it away for broadcasting
	nTx, err := core.Name(nodeAddr, signAddr, pubkey, addr, amtS, nonceS, feeS, name, data)
	if err != nil {
		fmt.Printf("corename error: %v\n", err)
	}
	fmt.Printf("chainsOD %v\n", chainID)
	fmt.Printf("signAddr %v\n", signAddr)
	fmt.Printf("nTX %v\n", nTx)
	//sign but don't broadcast
	_, err = core.SignAndBroadcast(chainID, nodeAddr, signAddr, nTx, true, false, false)
	if err != nil {
		fmt.Printf("sign error: %v\n", err)
	}

	n := new(int64)
	w := new(bytes.Buffer)
	wire.WriteBinary(nTx, w, n, &err)

	err = postTxData(nodeAddr, hash, w.Bytes())
	if err != nil {
		fmt.Printf("post error: %v\n", err)
	}
	return "", nil
}

func postTxData(nodeAddr, hash string, txData []byte) error {

	// post needs to be to toadserver endpoint, which'll route the TX to the node
	// or should the toadserver have an endpoint that does that?
	txD := bytes.NewReader(txData)
	//it can also query for th name reg to ensure things are good
	endpoint := "http://0.0.0.0:11113/" + "receiveNameTx/" + hash
	_, err := http.Post(endpoint, "", txD)
	if err != nil {
		fmt.Printf("post error: %v\n", err)
	}
	return nil
}

func receiveNameTx(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		str := r.URL.Path[1:]
		hash := strings.Split(str, "/")[1]

		fmt.Printf("hash in rnt %v\n", hash)

		txData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("error reading body: %v\n", err)
		}

		tx := new(types.NameTx)
		n := new(int64)
		txD := bytes.NewReader(txData)

		wire.ReadBinary(tx, txD, n, &err)
		if err != nil {
			fmt.Printf("error reading binary: %v\n", err)
		}

		rpcAddr := os.Getenv("MINTX_NODE_ADDR")
		_, err1 := core.Broadcast(tx, rpcAddr)
		if err1 != nil {
			fmt.Printf("error broadcasting: %v\n", err1)
		}

		//TODO check Name reg

		endpoint := "http://0.0.0.0:11113/" + "cacheHash/" + hash
		_, err2 := http.Post(endpoint, "", nil)
		if err2 != nil {
			fmt.Printf("cache post error: %v\n", err2)
		}
		w.Write([]byte("success!"))
	}
}

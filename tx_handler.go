package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/eris-ltd/common/go/ipfs"
	"github.com/eris-ltd/mint-client/mintx/core"
	"github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/wire"
)

func receiveNameTx(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		str := r.URL.Path[1:]
		hash := strings.Split(str, "/")[0]

		body := r.Body
		txData, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Printf("error reading body: %v\n", err)
		}

		var tx types.Tx
		n := new(int64) //, new(error)
		txD := bytes.NewReader(txData)
		wire.ReadBinary(tx, txD, n, &err)
		if err != nil {
			// ,.,,
		}
		nodeAddr := "0.0.0.0:46656"
		_, err = core.Broadcast(tx, nodeAddr)
		if err != nil {
			// ,.,,
		}

		endpoint := "0.0.0.0:11113/" + "cacheHash/" + hash
		_, err = http.Post(endpoint, "", nil)
		if err != nil {
			fmt.Printf("post error: %v\n", err)
		}
		w.Write([]byte("success!"))
	}
}

//pass hash into here, somehow!
func cacheHash(w http.ResponseWriter, r *http.Request) {
	//XXX this endpoint should also require authentication
	//if name reg update is legit, make post to ipfs hosts w/ hash in URL;
	//endpoint takes hash and pins it locally
	hash := "212"
	pinned, err := ipfs.PinToIPFS(hash, bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Printf("error pinning to IPFS: %v\n", err)
	}
	fmt.Printf("%v has been added succesfully\n", pinned)
}

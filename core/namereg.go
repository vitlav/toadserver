package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/eris-ltd/mint-client/mintx/core"
	//"github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/wire"
)

func UpdateNameReg(fileName, hash string) error {
	nodeAddr := os.Getenv("MINTX_NODE_ADDR")
	signAddr := os.Getenv("MINTX_SIGN_ADDR")
	chainID := os.Getenv("MINTX_CHAINID")
	pubkey := strings.TrimSpace(os.Getenv("MINTX_PUBKEY")) //because bash
	addr := ""
	amtS := "10000"
	nonceS := ""
	feeS := "0"
	name := fileName
	data := hash

	log.WithFields(log.Fields{
		"MINTX_NODE_ADDR": nodeAddr,
		"MINTX_CHAINID":   chainID,
		"MINTX_PUBKEY":    pubkey,
		"MINTX_SIGN_ADDR": signAddr,
		"name":            fileName,
		"data":            data,
		"amount":          amtS,
	}).Warn("Building a nameTx:")

	nTx, err := core.Name(nodeAddr, signAddr, pubkey, addr, amtS, nonceS, feeS, name, data)
	if err != nil {
		return errors.New(fmt.Sprintf("corename error: %v\n", err))
	}
	log.WithField("=>", fmt.Sprintf("%v", nTx)).Warn("Success, nameTx created:")

	//sign but don't broadcast

	log.Warn("Signing transaction...")
	_, err = core.SignAndBroadcast(chainID, nodeAddr, signAddr, nTx, true, false, false)
	if err != nil {
		return errors.New(fmt.Sprintf("sign error: %v\n", err))

	}

	n := new(int64)
	w := new(bytes.Buffer)
	wire.WriteBinary(nTx, w, n, &err)

	// post needs to be to toadserver endpoint, which'll
	// eventually route the TX to the nodes using mindy
	txD := bytes.NewReader(w.Bytes())
	//it can also query for the name reg to ensure things are good
	endpoint := "http://0.0.0.0:11113/" + "receiveNameTx/" //+ hash
	_, err = http.Post(endpoint, "", txD)
	if err != nil {
		return errors.New(fmt.Sprintf("post error: %v\n", err))
	}
	if err := os.Remove(fileName); err != nil {
		return errors.New(fmt.Sprintf("remove file error: %v\n", err))
	}
	return nil

}

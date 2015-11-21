package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	cclient "github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/rpc/core_client"
)

var (
	DefaultNodeRPCHost = "0.0.0.0"
	DefaultNodeRPCPort = "46657"
	DefaultNodeRPCAddr = "http://" + DefaultNodeRPCHost + ":" + DefaultNodeRPCPort

	DefaultChainID string

	REQUEST_TYPE = "HTTP"
)

// override the hardcoded defaults with env variables if they're set
func init() {
	nodeAddr := os.Getenv("MINTX_NODE_ADDR")
	if nodeAddr != "" {
		DefaultNodeRPCAddr = nodeAddr
	}

	chainID := os.Getenv("MINTX_CHAINID")
	if chainID != "" {
		DefaultChainID = chainID
	}

}

//TODO better errors
func getInfos(fileName string) (string, error) {
	c := cclient.NewClient(DefaultNodeRPCAddr, REQUEST_TYPE)
	if fileName == "" {
		//to eventually support an endpoint that lists available files
		_, err := c.ListNames()
		ifExit(err)
		/*res := make([]string, len(names.Names))
		i := 0
		for n := range names.Names {
			res[i] = n.Entry.Name
			i += 1
		}
		result := string.Join(res, "\n")*/
		return "", nil //result of format output
	} else {
		n, err := c.GetName(fileName)
		ifExit(err)

		name := n.Entry.Data

		//formatOutput(r)
		return name, nil //result of format output
	}
}

//this func is just a check
func checkAddr(addr string, w io.Writer) error {
	c := cclient.NewClient(DefaultNodeRPCAddr, REQUEST_TYPE)
	if addr == "" {
		_, err := c.ListAccounts()
		ifExit(err)
		//formatOutput(r)
		return nil //result of format output
	} else {
		addrBytes, err := hex.DecodeString(addr)
		if err != nil {
			exit(fmt.Errorf("Addr %s is improper hex: %v", addr, err))
		}
		r, err := c.GetAccount(addrBytes)
		ifExit(err)
		if r == nil {
			exit(fmt.Errorf("Account %X does not exist", addrBytes))
		}
		r2 := r.Account
		if r2 == nil {
			exit(fmt.Errorf("Account %X does not exist", addrBytes))
		}
		//formatOutput(c, 1, r2)
	}

	//TODO deal with this gracefully
	//	w.Write([]byte("Permission denied, invalid address\n"))
	return nil //errors.New("Permission denied, invalid address")

	//get more infos (like check if they have perms!)

}

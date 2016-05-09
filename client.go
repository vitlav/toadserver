package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

<<<<<<< HEAD
	cclient "github.com/tendermint/tendermint/rpc/core_client"
=======
	cclient "github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/rpc/core_client"
	. "github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
>>>>>>> godep /common & use IfExit
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

func getInfos(fileName string) (string, error) {
	c := cclient.NewClient(DefaultNodeRPCAddr, REQUEST_TYPE)
	if fileName == "" {
		//to support an endpoint that lists available files
		_, err := c.ListNames()
		IfExit(err)
		/*res := make([]string, len(names.Names))
		i := 0
		for n := range names.Names {
			res[i] = n.Entry.Name
			i += 1
		}
		result := string.Join(res, "\n")*/
		return "", nil
	} else {
		n, err := c.GetName(fileName)
		IfExit(err)

		name := n.Entry.Data
		return name, nil
	}
}

//this func is just a check
func checkAddr(addr string, w io.Writer) error {
	c := cclient.NewClient(DefaultNodeRPCAddr, REQUEST_TYPE)
	if addr == "" {
		_, err := c.ListAccounts()
		IfExit(err)
		//formatOutput(r)
		return nil //result of format output
	} else {
		addrBytes, err := hex.DecodeString(addr)
		if err != nil {
			IfExit(fmt.Errorf("Addr %s is improper hex: %v", addr, err))
		}
		r, err := c.GetAccount(addrBytes)
		IfExit(err)
		if r == nil {
			IfExit(fmt.Errorf("Account %X does not exist", addrBytes))
		}
		r2 := r.Account
		if r2 == nil {
			IfExit(fmt.Errorf("Account %X does not exist", addrBytes))
		}
	}
	//TODO get more infos (like check if they have perms!)
	//something like: w.Write([]byte("Permission denied, invalid address\n"))
	return nil
}

package main

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	//"net"
	"net/http"
	"os"
	//"reflect"
	"strings"
	"time"

	//"github.com/eris-ltd/mint-client/Godeps/_workspace/src/github.com/tendermint/tendermint/types"
	//"github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/wire"

	//cclient "github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/rpc/core_client"

	"github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/eris-ltd/common/go/ipfs"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Receiving POST request")

		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[1]

		fmt.Printf("File name to register:\t%s\n", fn)

		body := r.Body
		b, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Printf("error reading file body: %v\n", err)
		}

		err = ioutil.WriteFile(fn, b, 0666)
		if err != nil {
			fmt.Printf("error writing temp file: %v\n", err)
		}
		//should just put on whoever is doing the sending's gateway; since cacheHash won't send it there anyways
		fmt.Println("Sending File to eris' IPFS gateway")

		// because IPFS is testy, we retry the put up to 5 times.
		//TODO move this functionality to /common
		var hash string
		passed := false
		for i := 0; i < 5; i++ {
			hash, err = ipfs.SendToIPFS(fn, "eris", bytes.NewBuffer([]byte{}))
			if err != nil {
				time.Sleep(2 * time.Second)
				continue
			} else {
				passed = true
				break
			}
		}
		if !passed {
			// final time will throw
			hash, err = ipfs.SendToIPFS(fn, "eris", bytes.NewBuffer([]byte{}))
			if err != nil {
				fmt.Printf("error sending to IPFS: %v\n", err)
			}
		}

		fmt.Printf("Hash received:\t%s\n", hash)

		fmt.Printf("Sending name registry transaction:\t%s:%s\n", fn, hash)
		err = UpdateNameReg(fn, hash)
		if err != nil {
			fmt.Printf("Error updating the name registry:\n%v\n", err)
			//return err
		} else {
			fmt.Println("Success updating the name registry")
			fmt.Println("Caching hash now")
			if err := cacheHashAll(hash); err != nil {
				fmt.Printf("error caching hash: %v\n", err)
			} else {
				fmt.Println("Congratulations, you have successfully added your file to the toadserver")
			}
		}
	}
}

func cacheHashAll(hash string) error {

	//TODO handle errors to prevent getting here...
	fmt.Println("Pinning hash to your local IPFS node")

	endpoint := fmt.Sprintf("http://0.0.0.0:11113/cacheHash/%s", hash)
	_, err := http.Post(endpoint, "", nil)
	if err != nil {
		fmt.Printf("cache post error: %v\n", err)
		//return err
	}

	// IPaddrs, _ := getTheNames() -> use mindy to get ipAddrs
	IPaddrs := os.Getenv("TOADSERVER_IPFS_NODES")
	IPs := strings.Split(IPaddrs, ",")
	for _, ip := range IPs {
		endpoint := fmt.Sprintf("http://%s:11113/cacheHash/%s", ip, hash)
		fmt.Printf("Pinning hash to remote IPFS/toadserver node: %s\n", endpoint)
		_, err := http.Post(endpoint, "", nil)
		if err != nil {
			fmt.Printf("cache post error to endpoint: %s\n%v\n", endpoint, err)
			continue
			//TODO return err?
		}
	}
	return nil
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Receiving GET request")
		//take filename & send ask chain for hash
		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[1]

		fmt.Printf("Looking for filename:\t%s\n", fn)
		hash, err := getInfos(fn)
		if err != nil {
			fmt.Printf("error getting name reg info: %v\n", err)
		}
		fmt.Printf("Found corresponding hash:\t%s\n", hash)

		// because IPFS is testy, we retry the put up to
		// 5 times.
		fmt.Println("Getting file from IPFS")
		passed := false
		//TODO move this to common
		for i := 0; i < 9; i++ {
			err = ipfs.GetFromIPFS(hash, fn, "", bytes.NewBuffer([]byte{}))
			if err != nil {
				time.Sleep(2 * time.Second)
				continue
			} else {
				passed = true
				break
			}
		}

		if !passed {
			// final time will throw
			err = ipfs.GetFromIPFS(hash, fn, "", bytes.NewBuffer([]byte{}))
			if err != nil {
				fmt.Printf("error getting file from IPFS: %v\n", err)
			}
		}

		contents, err := ioutil.ReadFile(fn)
		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
		}
		w.Write(contents) //outputfile

		err = os.Remove(fn)
		if err != nil {
			fmt.Printf("error removing file: %v\n", err)
		}

		fmt.Println("Congratulations, you have successfully retreived you file from the toadserver")
	}
}

//XXX this endpoint should require authentication
func cacheHash(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Path[1:]
	hash := strings.Split(str, "/")[1]

	pinned, err := ipfs.PinToIPFS(hash, bytes.NewBuffer([]byte{}))
	if err != nil {
		strang := fmt.Sprintf("error pinning to local IPFS node: %v\n", err)
		w.Write([]byte(strang))
	} else {
		w.Write([]byte(fmt.Sprintf("Caching succesful:\t%s\n", pinned)))
	}
}

//XXX for getting IPaddrs that will do all the caching
// (rather than pass in TOADSERVER_IPFS_NODES as a csv string of addrs)
// need to think of the name spacing carefullly
// this'll block on the tendermint refactor
/*func getTheNames() ([]string, error) {
	mindy := os.Getenv("TOADSERVER_IPFS_NODES")
	mindyUrl := mindy + ":46657/"
	//fmt.Printf("Pinning hash to you mindy nodes at:\t%s\n", mindyUrl)
	c := cclient.NewClient(mindyUrl, "HTTP")

	res, err := c.ListNames()
	ifExit(err)

	allTheNames, err := formatOutput([]string{}, 1, res)
	if err != nil {
		fmt.Printf("error formating output: %v\n")
	}

	byteNames := []byte(allTheNames)

	rln := struct {
		BlockHeight int                   `json:"block_height"`
		Names       []*types.NameRegEntry `json:"names"`
	}{}

	type NameRegEntry struct {
		Name    string `json:"name"`    // registered name for the entry
		Owner   []byte `json:"owner"`   // address that created the entry
		Data    string `json:"data"`    // data to store under this name
		Expires int    `json:"expires"` // block at which this entry expires
	}

	if err := json.Unmarshal(byteNames, &rln); err != nil {
		return []string{""}, err
	}

	names := make([]string, len(rln.Names))

	for i, name := range rln.Names {
		names[i] = name.Name
	}

	return names, nil
}
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/eris-ltd/mint-client/Godeps/_workspace/src/github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/wire"

	cclient "github.com/eris-ltd/toadserver/Godeps/_workspace/src/github.com/tendermint/tendermint/rpc/core_client"

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
		} else {
			fmt.Println("Success updating the name registry")
		}

		if err := os.Remove(fn); err != nil {
			fmt.Printf("remove file error: %v\n", err)
		}

		//TODO handle errors to prevent getting here...
		fmt.Println("Pinning hash to your local IPFS node")

		endpoint := "http://0.0.0.0:11113/" + "cacheHash/" + hash
		_, err2 := http.Post(endpoint, "", nil)
		if err2 != nil {
			fmt.Printf("cache post error: %v\n", err2)
		}

		//		names, _ := getTheNames()
		names := []string{"toadserver1.interblock.io", "toadserver2.interblock.io", "toadserver3.interblock.io"}
		for _, name := range names {
			endpoint := "http://" + name + ":11113/cacheHash/" + hash
			fmt.Printf("Pinning hash to remote IPFS/toadserver node: %s\n", name)
			_, err3 := http.Post(endpoint, "", nil)
			if err3 != nil {
				fmt.Printf("cache post error to endpoint: %s\n%v\n", endpoint, err3)
				continue
			}
		}

		//TODO catch things before this stmt
		fmt.Println("Congratulations, you have successfully added your file to the toadserver")
	}
}

func getTheNames() ([]string, error) {
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

	/*type NameRegEntry struct {
		Name    string `json:"name"`    // registered name for the entry
		Owner   []byte `json:"owner"`   // address that created the entry
		Data    string `json:"data"`    // data to store under this name
		Expires int    `json:"expires"` // block at which this entry expires
	}*/

	if err := json.Unmarshal(byteNames, &rln); err != nil {
		return []string{""}, err
	}

	names := make([]string, len(rln.Names))

	for i, name := range rln.Names {
		names[i] = name.Name
	}

	return names, nil
}

func prettyPrint(o interface{}) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, wire.JSONBytes(o), "", "\t")
	if err != nil {
		return "", err
	}
	return string(prettyJSON.Bytes()), nil
}

func FieldFromTag(v reflect.Value, field string) (string, error) {
	iv := v.Interface()
	st := reflect.TypeOf(iv)
	for i := 0; i < v.NumField(); i++ {
		tag := st.Field(i).Tag.Get("json")
		if tag == field {
			return st.Field(i).Name, nil
		}
	}
	return "", fmt.Errorf("Invalid field name")
}

func formatOutput(args []string, i int, o interface{}) (string, error) {
	if len(args) < i+1 {
		return prettyPrint(o)
	}
	arg0 := args[i]
	v := reflect.ValueOf(o).Elem()
	name, err := FieldFromTag(v, arg0)
	if err != nil {
		return "", err
	}
	f := v.FieldByName(name)
	return prettyPrint(f.Interface())
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
//if name reg update is legit, make post to ipfs hosts (using mindy!) w/ hash in URL;
func cacheHash(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Path[1:]
	hash := strings.Split(str, "/")[1]

	pinned, err := ipfs.PinToIPFS(hash, bytes.NewBuffer([]byte{}))
	if err != nil {
		strang := fmt.Sprintf("error pinning to local IPFS node: %v\n", err)
		w.Write([]byte(strang))
	} else {
		fmt.Printf("Caching succesful:\t%s\n", pinned)
	}
}

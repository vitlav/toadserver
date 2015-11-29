package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

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
		//[zr] also, odd that pinning rarely hangs, only post & get from IPFS
		fmt.Println("Pinning hash to your local IPFS node")
		endpoint := "http://0.0.0.0:11113/" + "cacheHash/" + hash
		_, err2 := http.Post(endpoint, "", nil)
		if err2 != nil {
			fmt.Printf("cache post error: %v\n", err2)
		}

		fmt.Println("Congratulations, you have successfully added your file to the toadserver")
	}
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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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

		//TODO write to temp file, rename file path
		// & pass that into SendToIPFS
		// then rm file
		err = ioutil.WriteFile(fn, b, 0666)
		if err != nil {
			fmt.Printf("error writing temp file: %v\n", err)
		}
		//should just put on whoever is doing the sending's gateway; since cacheHash won't send it there anyways
		fmt.Println("Sending File to eris' IPFS gateway")
		hash, err := ipfs.SendToIPFS(fn, "eris", bytes.NewBuffer([]byte{}))
		if err != nil {
			fmt.Printf("error sending to IPFS: %v\n", err)
		}
		fmt.Printf("Hash received:\t%s\n", hash)

		fmt.Println("Sending name registry transaction:\t%s:%s", fn, hash)
		err = UpdateNameReg(fn, hash)
		if err != nil {
			fmt.Printf("Error updating the name registry:\n%v\n", err)
		} else {
			fmt.Println("Success updating the name registry")
		}

		//if no error:
		fmt.Println("Pinning hash to your local IPFS node")
		endpoint := "http://0.0.0.0:11113/" + "cacheHash/" + hash
		_, err2 := http.Post(endpoint, "", nil)
		if err2 != nil {
			fmt.Printf("cache post error: %v\n", err2)
		}
		//success msg for pin in endpoint

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
		hash, err := getInfos(fn) //TODO make proper errors
		if err != nil {
			fmt.Printf("error getting name reg info: %v\n", err)
		}
		fmt.Printf("Found corresponding hash:\t%s\n", hash)

		fmt.Println("Getting file from IPFS")
		err = ipfs.GetFromIPFS(hash, fn, "", bytes.NewBuffer([]byte{}))
		if err != nil {
			fmt.Printf("error getting file from IPFS: %v\n", err)
		}

		contents, err := ioutil.ReadFile(fn)
		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
		}
		w.Write(contents) //outputfile

		//or don't remove? & just output like above
		err = os.Remove(fn)
		if err != nil {
			fmt.Printf("error removing file: %v\n", err)
		}

		fmt.Println("Congratulations, you have successfully retreived you file from the toadserver")
	}
}

//XXX this endpoint should also require authentication
//if name reg update is legit, make post to ipfs hosts w/ hash in URL;
//ultimately, it should query mindy to know where to go pinning...
//(i.e., looks for toadserver nodes that expose this endpoint & hit it
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

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
		//TODO get amt from url
		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[1] //the value that comes after endpoint

		fmt.Printf("fielname: %s\n", fn)

		body := r.Body
		b, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Printf("error reading body: %v\n", err)
		}

		// which file perms?
		err = ioutil.WriteFile(fn, b, 0666)
		if err != nil {
			fmt.Printf("error writing temp file: %v\n", err)
		}
		//should just put on whoever is doing the sending's gateway; since cacheHash won't send it there anyways
		hash, err := ipfs.SendToIPFS(fn, "eris", bytes.NewBuffer([]byte{}))
		if err != nil {
			fmt.Printf("error sending to IPFS: %v\n", err)
		}
		//TODO rm file

		_, _ = UpdateNameReg(fn, hash, "")

	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//take filename & send ask chain for hash
		str := r.URL.Path[1:]

		fn := strings.Split(str, "/")[1]
		hash := getInfos(fn)

		err := ipfs.GetFromIPFS(hash, fn, "", bytes.NewBuffer([]byte{}))
		if err != nil {
			fmt.Printf("error getting file from IPFS: %v\n", err)
		}
		contents, err := ioutil.ReadFile(fn)
		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
		}
		w.Write(contents)

		//or don't remove? & just output like above
		err = os.Remove(fn)
		if err != nil {
			fmt.Printf("error removing file: %v\n", err)
		}
	}
}

func cacheHash(w http.ResponseWriter, r *http.Request) {
	//XXX this endpoint should also require authentication
	//if name reg update is legit, make post to ipfs hosts w/ hash in URL;
	//endpoint takes hash and pins it locally
	str := r.URL.Path[1:]
	hash := strings.Split(str, "/")[1]

	pinned, err := ipfs.PinToIPFS(hash, bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Printf("error pinning to IPFS: %v\n", err)
	}
	fmt.Printf("%v has been added succesfully\n", pinned)
}

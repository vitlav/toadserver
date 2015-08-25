package main

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"net/http"
	"strings"

	"github.com/eris-ltd/common/go/ipfs"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[1]

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

		hash, err := ipfs.SendToIPFS(fn, "eris", bytes.NewBuffer([]byte{}))
		if err != nil {
			fmt.Printf("error sending to IPFS: %v\n", err)
		}
		//TODO rm file

		amt := r.Header.Get("amt") //put in URL
		_, _ = UpdateNameReg(fn, hash, amt)

	}
}

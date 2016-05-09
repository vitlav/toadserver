package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/eris-ltd/mint-client/mintx/core"
	"github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/wire"

	"github.com/eris-ltd/common/go/ipfs"
)

// todo clean this up!
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Warn("Receiving POST request")

		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[1]

		log.WithField("=>", fn).Warn("File to register:")

		body := r.Body
		b, err := ioutil.ReadAll(body)
		if err != nil {
			log.Warn("error reading file body:")
			log.Error(err)
			return
		}

		if err := ioutil.WriteFile(fn, b, 0666); err != nil {
			log.Warn("error writing temp file:")
			log.Error(err)
			return
		}
		//should just put on whoever is doing the sending's gateway; since cacheHash won't send it there anyways
		log.Warn("Sending File to eris' IPFS gateway")

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
				log.Warn("error sending to IPFS:")
				log.Error(err)
				return
			}
		}
		log.WithField("=>", hash).Warn("Hash received:")

		log.WithFields(log.Fields{
			"filename": fn,
			"hash":     hash,
		}).Warn("Sending name registry transaction:")

		if err := UpdateNameReg(fn, hash); err != nil {
			log.Warn("error updating name registry:")
			log.Error(err)
			return
		}

		if err := cacheHashAll(hash); err != nil {
			log.Warn("error caching hash:")
			log.Error(err)
			return
		}
		log.Warn("Congratulations, you have successfully added your file to the toadserver")
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Warn("Receiving GET request")
		//take filename & send ask chain for hash
		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[1]

		log.WithField("=>", fn).Warn("Looking for filename:")
		hash, err := getInfos(fn)
		if err != nil {
			log.Warn("error getting name reg info:")
			log.Error(err)
			return
		}

		log.WithField("=>", hash).Warn("Found corresponding hash:")
		log.Warn("Getting it from IPFS...")

		// because IPFS is testy, we retry the put up to
		// 5 times.
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
			if err := ipfs.GetFromIPFS(hash, fn, "", bytes.NewBuffer([]byte{})); err != nil {
				log.Warn("error getting file from IPFS:")
				log.Error(err)
				return
			}
		}

		contents, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Warn("error reading file:")
			log.Error(err)
			return
		}
		w.Write(contents) //outputfile

		if err := os.Remove(fn); err != nil {
			log.Warn("error removing file:")
			log.Error(err)
			return
		}

		log.Warn("Congratulations, you have successfully retreived you file from the toadserver")
	}
}

// TODO this endpoint should require authentication
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

func receiveNameTx(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//TODO check valid Name reg
		//str := r.URL.Path[1:]
		//hash := strings.Split(str, "/")[1]

		txData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Warn("error reading body:")
			log.Error(err)
		}

		tx := new(types.NameTx)
		n := new(int64)
		txD := bytes.NewReader(txData)

		wire.ReadBinary(tx, txD, n, &err)
		if err != nil {
			log.Warn("error reading binary:")
			log.Error(err)
		}

		rpcAddr := os.Getenv("MINTX_NODE_ADDR")
		_, err1 := core.Broadcast(tx, rpcAddr)
		if err1 != nil {
			log.Warn("error broadcasting:")
			log.Error(err)
		}
	}
}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	tscore "github.com/eris-ltd/toadserver/core"
	log "github.com/Sirupsen/logrus"
	"github.com/eris-ltd/mint-client/mintx/core"
	"github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/wire"

	"github.com/eris-ltd/common/go/ipfs"
)

// todo clean this up!
func postHandler(w http.ResponseWriter, r *http.Request) *toadError {
	if r.Method == "POST" {
		log.Warn("Receiving POST request")

		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[1]

		log.WithField("=>", fn).Warn("File to register:")

		body := r.Body
		b, err := ioutil.ReadAll(body)
		if err != nil {
			return &toadError{err, "error reading body", 400}
		}

		if err := ioutil.WriteFile(fn, b, 0666); err != nil {
			return &toadError{err, "error writing file", 400}
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
				return &toadError{err, "error sending to IPFS", 400}
			}
		}
		log.WithField("=>", hash).Warn("Hash received:")

		log.WithFields(log.Fields{
			"filename": fn,
			"hash":     hash,
		}).Warn("Sending name registry transaction:")

		if err := tscore.UpdateNameReg(fn, hash); err != nil {
			return &toadError{err, "error updating namereg", 400}
		}

		if err := tscore.CacheHashAll(hash); err != nil {
			return &toadError{err, "error cashing hashes", 400}
		}
		log.Warn("Congratulations, you have successfully added your file to the toadserver")
	}
	return nil
}

func getHandler(w http.ResponseWriter, r *http.Request) *toadError {
	if r.Method == "GET" {
		log.Warn("Receiving GET request")
		//take filename & send ask chain for hash
		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[1]

		log.WithField("=>", fn).Warn("Looking for filename:")
		hash, err := tscore.GetInfos(fn)
		if err != nil {
			return &toadError{err, "error getting namereg info", 400}
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
				return &toadError{err, "error getting file from IPFS", 400}
			}
		}

		contents, err := ioutil.ReadFile(fn)
		if err != nil {
			return &toadError{err, "error reading file", 400}
		}
		w.Write(contents) //outputfile

		if err := os.Remove(fn); err != nil {
			return &toadError{err, "error removing file", 400}
		}

		log.Warn("Congratulations, you have successfully retreived you file from the toadserver")
	}
	return nil
}

// TODO this endpoint should require authentication
func cacheHash(w http.ResponseWriter, r *http.Request) *toadError {
	str := r.URL.Path[1:]
	hash := strings.Split(str, "/")[1]

	pinned, err := ipfs.PinToIPFS(hash, bytes.NewBuffer([]byte{}))
	if err != nil {
		return &toadError{err, "error pinning to local IPFS node", 400}
	}
	w.Write([]byte(fmt.Sprintf("Caching succesful:\t%s\n", pinned)))
	return nil
}

func receiveNameTx(w http.ResponseWriter, r *http.Request) *toadError {
	if r.Method == "POST" {
		//TODO check valid Name reg
		//str := r.URL.Path[1:]
		//hash := strings.Split(str, "/")[1]

		txData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return &toadError{err, "error reading body", 400}
		}

		tx := new(types.NameTx)
		n := new(int64)
		txD := bytes.NewReader(txData)

		wire.ReadBinary(tx, txD, n, &err)
		if err != nil {
			return &toadError{err, "error reading binary", 400}
		}

		rpcAddr := os.Getenv("MINTX_NODE_ADDR")
		_, err1 := core.Broadcast(tx, rpcAddr)
		if err1 != nil {
			return &toadError{err, "error broadcasting", 400}
		}
	}
	return nil
}

// ---------------- error handler ------------------

type toadError struct {
	Error   error
	Message string
	Code    int
}

type toadHandler func(http.ResponseWriter, *http.Request) *toadError

func (endpoint toadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := endpoint(w, r); err != nil {
		http.Error(w, err.Message, err.Code)
	}
}

/* status codes
StatusBadRequest = 400
StatusNotFound = 404
StatusInternalServerError = 500
*/

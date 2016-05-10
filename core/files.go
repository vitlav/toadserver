package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/eris-ltd/common/go/ipfs"
)

func PutFile(fileName string, body []byte) (string, error) {

	//should just put on whoever is doing the sending's gateway; since cacheHash won't send it there anyways
	//log.Warn("Sending File to eris' IPFS gateway")

	// because IPFS is testy, we retry the put up to 5 times.
	//TODO move this functionality to /common
	var hash string
	var err error
	passed := false
	for i := 0; i < 5; i++ {
		hash, err = ipfs.SendToIPFS(fileName, "", bytes.NewBuffer([]byte{}))
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
		hash, err = ipfs.SendToIPFS(fileName, "", bytes.NewBuffer([]byte{}))
		if err != nil {
			return "", err
			//return &toadError{err, "error sending to IPFS", 400}
		}
	}
	log.WithField("=>", hash).Warn("Hash received:")
	return hash, nil
}

func GetFile(fileName, hash string) ([]byte, error) {
	// because IPFS is testy, we retry the put up to
	// 5 times.
	passed := false
	//TODO move this to common
	for i := 0; i < 9; i++ {
		if err := ipfs.GetFromIPFS(hash, fileName, "", bytes.NewBuffer([]byte{})); err != nil {
			log.Warn(fmt.Sprintf("Trying: %v\n", i))
			time.Sleep(2 * time.Second)
			continue
		} else {
			passed = true
			break
		}
	}

	if !passed {
		// final time will throw
		if err := ipfs.GetFromIPFS(hash, fileName, "", bytes.NewBuffer([]byte{})); err != nil {
			return nil, err
			//return &toadError{err, "error getting file from IPFS", 400}
		}
	}

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
		//return &toadError{err, "error reading file", 400}
	}

	return contents, nil
}

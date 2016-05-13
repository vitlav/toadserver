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

	// because IPFS is testy, we retry the put up to 5 times.
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
		}
	}
	log.WithField("=>", hash).Warn("Hash received:")
	return hash, nil
}

func GetFile(fileName, hash string) ([]byte, error) {
	passed := false
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
		if err := ipfs.GetFromIPFS(hash, fileName, "", bytes.NewBuffer([]byte{})); err != nil {
			return nil, err
		}
	}

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

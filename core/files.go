package core

import (
	"bytes"
	//"fmt"
	"io/ioutil"
	//"time"

	log "github.com/Sirupsen/logrus"
	"github.com/eris-ltd/common/go/ipfs"
)

func PutFile(fileName string, body []byte) (string, error) {
	hash, err := ipfs.SendToIPFS(fileName, "", bytes.NewBuffer([]byte{}))
	if err != nil {
		return "", err
	}
	log.WithField("=>", hash).Warn("Hash received:")
	return hash, nil
}

func GetFile(fileName, hash string) ([]byte, error) {
	if err := ipfs.GetFromIPFS(hash, fileName, "", bytes.NewBuffer([]byte{})); err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

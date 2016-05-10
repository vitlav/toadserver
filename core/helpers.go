package main

import (
	//"bytes"
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
	"strings"
	//"time"

	log "github.com/Sirupsen/logrus"

)

func cacheHashAll(hash string) error {

	//TODO handle errors to prevent getting here...
	//	log.Warn("Pinning hash to your local IPFS node")

	//pinned, err := ipfs.PinToIPFS(hash, bytes.NewBuffer([]byte{}))
	/*	if err != nil {
			log.WithField("=>", fmt.Sprintf("%s", err)).Warn("")
			return fmt.Errorf("error pinning to local IPFS node: %v\n", err)
		}
		log.WithField("=>", pinned).Warn("Hash pinned to you local node")

		//XXX the problem is here!
		endpoint := fmt.Sprintf("http://0.0.0.0:11113/cacheHash/%s", hash)
		_, err := http.Post(endpoint, "", nil)
		if err != nil {
			log.Warn("error making post request:")
			log.Error(err)
			//return err
		}*/

	// IPaddrs, _ := getTheNames() -> use mindy to get ipAddrs
	IPaddrs := os.Getenv("TOADSERVER_IPFS_NODES")
	IPs := strings.Split(IPaddrs, ",")
	for _, ip := range IPs {
		endpoint := fmt.Sprintf("http://%s:11113/cacheHash/%s", ip, hash)
		log.WithField("=>", endpoint).Warn("Pinning hash to remote IPFS/toadserver node:")
		_, err := http.Post(endpoint, "", nil)
		if err != nil {
			log.WithField("=>", endpoint).Warn("error making post request to:")
			log.Error(err)
			continue
			//TODO return err?
		}
	}
	return nil
}

//XXX for getting IPaddrs that will do all the caching
// (rather than pass in TOADSERVER_IPFS_NODES as a csv string of addrs)
// need to think of the name spacing carefullly
// this'll block on the tendermint refactor
/*func getTheNames() ([]string, error) {
	mindy := os.Getenv("TOADSERVER_IPFS_NODES")
	mindyUrl := mindy + ":46657/"
	//fmt.Printf("Pinning hash to you mindy nodes at:\t%s\n", mindyUrl)
	c := cclient.NewClient(mindyUrl, "HTTP")

	res, err := c.ListNames()
	IfExit(err) (common)

	allTheNames, err := formatOutput([]string{}, 1, res)
	if err != nil {
		fmt.Printf("error formating output: %v\n")
	}

	byteNames := []byte(allTheNames)

	rln := struct {
		BlockHeight int                   `json:"block_height"`
		Names       []*types.NameRegEntry `json:"names"`
	}{}

	type NameRegEntry struct {
		Name    string `json:"name"`    // registered name for the entry
		Owner   []byte `json:"owner"`   // address that created the entry
		Data    string `json:"data"`    // data to store under this name
		Expires int    `json:"expires"` // block at which this entry expires
	}

	if err := json.Unmarshal(byteNames, &rln); err != nil {
		return []string{""}, err
	}

	names := make([]string, len(rln.Names))

	for i, name := range rln.Names {
		names[i] = name.Name
	}

	return names, nil
}
*/

//XXX will likely change post edb/tmint refactor

/*
import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/tendermint/tendermint/wire"
)

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
}*/

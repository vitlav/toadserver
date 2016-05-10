package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/eris-ltd/common/go/common"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

func startServer(cmd *cobra.Command, args []string) {
	log.Warn("Initializing toadserver...")

	mux := http.NewServeMux()
	mux.Handle("/postfile", toadHandler(postHandler)) //post a file with its contents to gateway, returns hash
	mux.Handle("/getfile", toadHandler(getHandler))   // request by name, receive contents

	//ts makes & signs a nameTx, then posts to a node, which does the broadcasting
	mux.Handle("/receiveNameTx", toadHandler(receiveNameTx)) //unpack tx, if valid, add to chain
	mux.Handle("/cacheHash", toadHandler(cacheHash))         //also if valid, pin hash on all hosts (except one that sent it :~( )

	handler := cors.Default().Handler(mux)

	log.WithField("port", ToadPort).Warn("Toadserver started at:")
	port := fmt.Sprintf(":%s", ToadPort)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Warn(err)
	}
}

func putFiles(cmd *cobra.Command, args []string) {
	fileName := args[0]

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		common.IfExit(err)
	}
	url := "http://192.168.99.100:" + ToadPort + "/postfile?fileName=" + fileName
	req, err := http.Post(url, "", file)
	if err != nil {
		common.IfExit(err)
	}
	fmt.Printf("YOYOYO %v", req)
}

func getFiles(cmd *cobra.Command, args []string) {
	fileName := args[0]

	url := "http://192.168.99.100:" + ToadPort + "/getfile?fileName=" + fileName
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		common.IfExit(err)
	}
	fmt.Printf("YOYOYO %v", resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.IfExit(err)
	}
	if err := ioutil.WriteFile(fileName, body, 0777); err != nil {
		common.IfExit(err)
	}
}

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
	mux.Handle("/postfile", toadHandler(postHandler))  //post a file with its contents to gateway, returns hash
	mux.Handle("/getfile", toadHandler(getHandler))    // request by name, receive contents
	mux.Handle("/listfiles", toadHandler(listHandler)) // request by name, receive contents

	// these are used under the hood
	mux.Handle("/receiveNameTx", toadHandler(receiveNameTx)) //unpack tx, if valid, add to chain
	mux.Handle("/cacheHash", toadHandler(cacheHash))         //also if valid, pin hash on all hosts (except one that sent it :~( )

	handler := cors.Default().Handler(mux)

	toads := fmt.Sprintf("%s:%s", ToadHost, ToadPort)
	log.WithField("=>", toads).Warn("Toadserver started at:")
	if err := http.ListenAndServe(toads, handler); err != nil {
		log.Warn(err)
	}
}

func urlHandler(host, port, endpoint, arg string) string {
	if arg == "" {
		return fmt.Sprintf("http://%s:%s%s", host, port, endpoint)
	} else {
		return fmt.Sprintf("http://%s:%s%s%s", host, port, endpoint, arg)
	}
}

func listFiles(cmd *cobra.Command, args []string) {
	url := urlHandler(ToadHost, ToadPort, "/listfiles", "")
	resp, err := http.Get(url)
	if err != nil {
		common.IfExit(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.IfExit(err)
	}
	log.Warn(string(b))
}

func putFiles(cmd *cobra.Command, args []string) {
	fileName := args[0]

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		common.IfExit(err)
	}
	url := urlHandler(ToadHost, ToadPort, "/postfile?fileName=", fileName)
	_, err = http.Post(url, "", file)
	if err != nil {
		common.IfExit(err)
	}
	log.Warn("success. file added to toadserver")
}

func getFiles(cmd *cobra.Command, args []string) {
	fileName := args[0]

	url := urlHandler(ToadHost, ToadPort, "/getfile?fileName=", fileName)
	resp, err := http.Get(url)
	if err != nil {
		common.IfExit(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.IfExit(err)
	}
	if err := ioutil.WriteFile(fileName, body, 0777); err != nil {
		common.IfExit(err)
	}
}

package main

import (
	"net/http"

	"github.com/rs/cors"

	log "github.com/Sirupsen/logrus"

	"github.com/spf13/cobra"
)

func startServer(cmd *cobra.Command, args []string) {
	log.Warn("Initializing toadserver...")

	mux := http.NewServeMux()
	mux.Handle("/postfile/", toadHandler(postHandler)) //post a file with its contents to gateway, returns hash
	//ts makes & signs a nameTx, then posts to a node, which does the broadcasting
	mux.Handle("/receiveNameTx/", toadHandler(receiveNameTx)) //unpack tx, if valid, add to chain
	mux.Handle("/cacheHash/", toadHandler(cacheHash))         //also if valid, pin hash on all hosts (except one that sent it :~( )

	mux.Handle("/getfile/", toadHandler(getHandler)) // request by name, receive contents

	handler := cors.Default().Handler(mux)

	if err := http.ListenAndServe(":11113", handler); err != nil {
		log.Warn(err)
	}
}

func putFiles(cmd *cobra.Command, args []string) {
}

func getFiles(cmd *cobra.Command, args []string) {
}

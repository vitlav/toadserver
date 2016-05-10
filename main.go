package main

import (
	"net/http"

	//"github.com/rs/cors" TODO

	log "github.com/Sirupsen/logrus"
	logger "github.com/eris-ltd/common/go/log"
)

func main() {
	log.SetFormatter(logger.ErisFormatter{})
	log.Warn("Initializing toadserver...")

	mux := http.NewServeMux()
	mux.Handle("/postfile/", toadHandler(postHandler)) //post a file with its contents to gateway, returns hash
	//ts makes & signs a nameTx, then posts to a node, which does the broadcasting
	mux.Handle("/receiveNameTx/", toadHandler(receiveNameTx)) //unpack tx, if valid, add to chain
	mux.Handle("/cacheHash/", toadHandler(cacheHash))         //also if valid, pin hash on all hosts (except one that sent it :~( )

	mux.Handle("/getfile/", toadHandler(getHandler)) // request by name, receive contents

	//handler := cors.Default().Handler(mux) TODO

	if err := http.ListenAndServe(":11113", mux); err != nil {
		log.Warn(err)
	}
}

type toadError struct {
	Error error
	Message string
	Code int
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

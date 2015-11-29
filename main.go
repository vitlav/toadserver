package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/postfile/", postHandler) //post a file with its contents to gateway, returns hash
	//ts makes & signs a nameTx, then posts to a node, which does the broadcasting
	mux.HandleFunc("/receiveNameTx/", receiveNameTx) //unpack tx, if valid, add to chain
	mux.HandleFunc("/cacheHash/", cacheHash)         //also if valid, pin hash on all hosts (except one that sent it :~( )

	mux.HandleFunc("/getfile/", getHandler) // request by name, receive contents

	http.ListenAndServe(":11113", mux)
}

//-------------------------------------------------------
//--helpers
func ifExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

package main

import (
	"demochain/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.BlockChain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandleFunc)
	http.HandleFunc("/blockchain/add", blockchainAddHandleFunc)
	http.ListenAndServe("localhost:8888", nil)
}

func blockchainGetHandleFunc(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchainAddHandleFunc(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandleFunc(w, r)
}

func main() {
	blockchain = core.NewBlockChain()
	run()
}

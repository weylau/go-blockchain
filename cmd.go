package main

import "demochain/core"

func main() {
	blockChain := core.NewBlockChain()
	err := blockChain.SendData("block 01")
	if err != nil {
		panic(err)
	}
	err = blockChain.SendData("block 02")
	if err != nil {
		panic(err)
	}
	blockChain.Print()
}

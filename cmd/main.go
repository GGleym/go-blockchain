package main

import (
	"fmt"

	"github.com/ggleym/go-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Current hash: %x\n", block.Hash)
	}
}

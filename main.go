package main

import (
	"fmt"

	"github.com/jisuhan3201/coian/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\nHash: %s\nPrevHash: %s\n\n",
			block.Data, block.Hash, block.PrevHash)
	}
}

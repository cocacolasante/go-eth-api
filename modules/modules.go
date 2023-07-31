package modules

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"socialverfication/web3context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestBlock(client ethclient.Client) *web3context.Block {
	defer func(){
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumer := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumer)

	if err != nil {
		log.Fatal(err)
	}

	_block := &web3context.Block{
		BlockNumber: block.Number().Int64(),
		Timestamp: block.Time(),
		Difficulty: block.Difficulty().Uint64(),
		Hash: block.Hash().String(),
		TransactionsCount: len(block.Transactions()),
		Transactions: []web3context.Transaction{},
	}

	for _, tx := range block.Transactions(){
		_block.Transactions = append(_block.Transactions, 
		web3context.Transaction{
			Hash: tx.Hash().String(),
			Value: tx.Value().String(),
			Gas: tx.GasPrice().Uint64(),
			Nonce: tx.Nonce(),
			To: tx.To().String(),
		})
	}
	return _block
}
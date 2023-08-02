package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	profile "socialverfication/gen"
	"socialverfication/web3context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main(){
	key, add, client, chainId, _, gasPrice := getWalletKeyClientChainNonceAndGas()
	
	// contract address from mumbai test net from actual project
	contractAdd := common.HexToAddress("0x4E539d65ff12913f380CBD7DdE1e7f5B7b5feCD1")
	deployerAdd := common.HexToAddress("0x11273F391609BF4C05CA23c6aD29D919a71dc37E")
	profileContract, err := profile.NewProfile(contractAdd, &client)
	if err != nil {
		log.Fatal("Could not fetch contract", err)
	}
	
	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		log.Fatal("Could not fetch contract", err)
	}
	
	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice
	
	
	if err != nil {
		log.Fatal("Could not fetch owner", err)
	}
	
	owner, err := profileContract.Owner(&bind.CallOpts{
		From: add,
		
	})
	
	if err != nil {
		log.Fatal("Could not fetch owner", err)
	}
	fmt.Println(owner)

	// fetching profile
	profile1, err := profileContract.Profiles(&bind.CallOpts{
		From: add,
	}, deployerAdd)
	if err != nil {
		log.Fatal("Could not fetch profile", err)
	}

	fmt.Println("Deployer profile", profile1)

	
}

func getWalletKeyClientChainNonceAndGas() (*keystore.Key, common.Address, ethclient.Client, *big.Int, uint64, *big.Int){
	b, err := os.ReadFile("wallet/UTC--2023-08-02T01-56-49.352251000Z--176e55443da73752238ccb7cc3542d8f75637396")
	if err != nil{
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, "password")
	if err != nil{
		log.Fatal(err)
	}
	client, chainId := web3context.ConnectToClient()

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil{
		log.Fatal(err)
	}
	
	
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil{
		log.Fatal(err)
	}
	return key, add, client, chainId, nonce, gasPrice
}


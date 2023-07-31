package web3context

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)



func ConnectToClient() ethclient.Client{
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	alchemyUrl := os.Getenv("ALCHEMY_URL")

	client, err := ethclient.DialContext(context.Background(), alchemyUrl)

	if err != nil {
		log.Fatal("Cannot connect")
	}

	defer client.Close()

	chainId, err := client.ChainID(context.Background())

	fmt.Printf("Chain id: %v \n", chainId)


	return *client
	  
	

}

func GetBalance(addr string) *big.Int{
	client := ConnectToClient()
	address := common.HexToAddress(addr)

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal("error getting block")
	}
	
	
	balance, err := client.BalanceAt(context.Background(), address, block.Number())
	if err != nil{
		log.Fatal("error getting balance")
	}

	fmt.Println(balance)

	return balance

}

func GenerateWallet(){
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(pvk)
	hexutil.Encode(pData)
	fmt.Println("Private Key", hexutil.Encode(pData))

	pubData := crypto.FromECDSAPub(&pvk.PublicKey)
	hexutil.Encode(pubData)
	fmt.Println(hexutil.Encode(pubData))


	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())

}
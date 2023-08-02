package main
// this file is used to deploy the contract
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
	"github.com/ethereum/go-ethereum/crypto"
)

func main(){
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
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil{
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = 3000000
	auth.Nonce = big.NewInt(int64(nonce))
	addy, tx, Profile, err := profile.DeployProfile(auth, &client)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("---------------------------------")
	fmt.Println(addy)
	fmt.Println("---------------------------------")
	fmt.Println(tx.Hash().Hex())
	fmt.Println("---------------------------------")
	fmt.Println(Profile)


}
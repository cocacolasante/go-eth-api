package main
// generate the wallet and check the wallet balance and address

import (
	"fmt"

	"log"
	"socialverfication/web3context"

	"github.com/ethereum/go-ethereum/common"
)





func main(){
	
	// addr := "0x11273F391609BF4C05CA23c6aD29D919a71dc37E"
	
	
	newWallet := web3context.GenerateWallet()
	balance := web3context.GetBalance(newWallet)

	log.Println("Balance", balance)


	fmt.Println("get test ether", common.HexToAddress("176e55443da73752238ccb7cc3542d8f75637396"))

	newBalance := web3context.GetBalance("0x176e55443da73752238Ccb7Cc3542d8F75637396")
	fmt.Println("Current balance for keystore wallet", newBalance)
	
}
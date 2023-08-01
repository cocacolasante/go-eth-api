package main

import (

	"log"
	"socialverfication/web3context"
)





func main(){
	
	// addr := "0x11273F391609BF4C05CA23c6aD29D919a71dc37E"
	
	
	newWallet := web3context.GenerateWallet()
	balance := web3context.GetBalance(newWallet)

	log.Println("Balance", balance)

	
	
	
}
package main

import (
	"log"
	"socialverfication/web3context"
)





func main(){
	
	addr := "0x11273F391609BF4C05CA23c6aD29D919a71dc37E"
	
	balance := web3context.GetBalance(addr)

	log.Println(balance)

	web3context.GenerateWallet()

	
}
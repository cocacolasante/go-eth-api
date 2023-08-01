package web3context

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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

func GetBlock(client ethclient.Client) *types.Block{
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal("error getting block")
	}
	return block
}

func GetBalance(addr string) *big.Int{
	client := ConnectToClient()
	address := common.HexToAddress(addr)

	block := GetBlock(client)
	
	
	balance, err := client.BalanceAt(context.Background(), address, block.Number())
	if err != nil{
		log.Fatal("error getting balance")
	}

	fmt.Println(balance)

	return balance

}

func GenerateWallet() (string) {
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


	return crypto.PubkeyToAddress(pvk.PublicKey).Hex()

}

func Keystore(password string){
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	a, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.Address)
}

func ReadKeystoreFile(password string){
	file, err := os.ReadFile("./wallet/UTC--2023-07-31T23-34-29.896779000Z--0042fd68b151e9c47ef045d41d55bd779ff2fb27")
	if err != nil{
		log.Fatal(err)
	}
	
	key, err := keystore.DecryptKey(file, password)

	if err != nil{
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	hexutil.Encode(pData)
	fmt.Println(hexutil.Encode(pData))
	
	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	hexutil.Encode(pubData)
	fmt.Println(hexutil.Encode(pubData))
	
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()

	fmt.Println(address)

	
}



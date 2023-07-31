package web3context

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"strings"

// 	// "fmt"
// 	"log"
// 	// "math/big"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/accounts/abi"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func FetchData(address string){
// 	ethAddress := common.HexToAddress(address)
// 	ethereumNodeUrl := "https://polygon-mumbai.g.alchemy.com/v2/42TFe2ZA21PCJgFwqy2IoD3bjNP5QJdE"

// 	contractAddress := common.HexToAddress("0x4E539d65ff12913f380CBD7DdE1e7f5B7b5feCD1")

	
// 	fileAbi := GetAbi()

// 	contractABI, err := abi.JSON(strings.NewReader(string(fileAbi)))
// 	if err != nil {
// 		log.Fatal("Cannot get contract abi", err)
// 	}
	
	

// 	contract, err := NewContract(contractAddress, ethereumNodeUrl, contractABI)
// 	if err != nil {
// 		log.Fatal("Cannot getconnet to eth")
// 	}

// 	profile, err := contract.GetProfile(context.Background(), ethAddress)

// 	if err != nil {
// 		log.Fatal("cannot fetch profile")
// 	}

// 	fmt.Print(profile)

// }

// func NewContract(address common.Address, nodeUrl string, contractAbi abi.ABI) (*Contract, error) {
// 	client, err := ethclient.Dial(nodeUrl)
// 	if err != nil {
// 		log.Fatal("Cannot fetch contract")
// 	}
// 	return &Contract{
// 		Address: address,
// 		ABI: contractAbi,
// 		Client: client,
// 	}, nil
// }


// func (c *Contract) GetProfile(ctx context.Context, address common.Address) (string, error) {
// 	data, err := c.ABI.Pack("getProfile", address)
// 	if err != nil {
// 		return "", err
// 	}

// 	msg := ethereum.CallMsg{
// 		To: &c.Address,
// 		Data: data,
// 	}

// 	result, err := c.Client.CallContract(ctx, msg, nil)
// 	if err != nil {
// 		return "", err
// 	}

// 	decodedData := string(result)

// 	return decodedData, nil


// }

// func GetAbi()ContractABI{
// 	abiFileContent, err := ioutil.ReadFile("constants/ProfileContract.json")
// 	if err != nil {
// 		log.Fatal("Cannot get contract in directory")
// 	}
// 	log.Println("got contract in directory")

// 	var contractAbi ContractABI

// 	err = json.Unmarshal(abiFileContent, &contractAbi)
// 	if err != nil {
// 		log.Fatal("Cannot marshall contract abi")
// 	}

// 	return contractAbi
// }
package main

import (
	"log"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	try "./build"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

const(
	url="http://localhost:8545"
	ContractAddress="0xA7033aa34c3c98F915E6cE60E70950f4d21C25fb"
	privatekey="4A4A8C67E9622E991F9783A41E6A84DC458F3D41F31D522BCAE6CE5556C3B8BF"
)


func clientStart() *ethclient.Client{
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Printf("Fail to connect, error: %s", err)
		log.Fatal(err)
	}
	return client
}

func loadContract(client *ethclient.Client) *try.Try{
	address := common.HexToAddress(ContractAddress)
	tryCore, err := try.NewTry(address, client)
	if err != nil {
		fmt.Printf("Fail to load contract, error: %s", err)
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
	return tryCore
}

func setAuth(client *ethclient.Client) *bind.TransactOpts{
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok{
		log.Fatal("error casting publick key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth
}

func subLogs() (ethereum.Subscription, chan types.Log){
	url_wss := "ws://127.0.0.1:8546"
	fmt.Println(url_wss)
	//url_wss := "/home/liyuankun/parity/parity-deploy/data/1/jsonrpc.ipc"
	client, err := ethclient.Dial(url_wss)
	if err != nil {
		fmt.Println(1)
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(ContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil{
		log.Fatal(err)
	}
	return sub,logs
}

func deployContract(client *ethclient.Client,auth *bind.TransactOpts) (string, string){
	address, tx, instance, err := try.DeployTry(auth, client)
	if err != nil{
		log.Fatal(err)
	}

	_ = instance
	return address.Hex(), tx.Hash().Hex()
}

func main(){
	//Connect
	client := clientStart()

	//Set auth
	auth := setAuth(client)

	//Deploy contract

	//contractAddress, tx := deployContract(client,auth)
	//fmt.Println(contractAddress,tx)


	//Load contract
	tryCore := loadContract(client)

	//Get event log
	sub,logs := subLogs()

	//Call function
	//Set Id
	tx, err := tryCore.SetId(auth,3000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	//Get logs
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
	//Get Id
	time.Sleep(time.Second*2)
	res, err := tryCore.GetId(nil)
	fmt.Println(res)

}
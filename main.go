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
)

func main(){
	// Connection
	client, err := ethclient.Dial("http://localhost:9545")
	if err != nil {
		fmt.Printf("Fail to connect, error: %s", err)
		log.Fatal(err)
	}

	//Load contract
	address := common.HexToAddress("0x337e8FE4C56f94c1b6b60eFC3F3406A4171bd177")
	tryCore, err := try.NewTry(address, client)
	if err != nil {
		fmt.Printf("Fail to load contract, error: %s", err)
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")

	//Set auth
	privateKey, err := crypto.HexToECDSA("e5ff4620232e7500b5be94a77711931e1d6b79f675767f649bd569a9c91ad9c5")
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

	//Call function
	tx, err := tryCore.SetId(auth,1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())

	res, err := tryCore.GetId(nil)

	fmt.Println(res)


}
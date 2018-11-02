package main

import (
	"encoding/json"
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
)
//curl --data '{"method":"eth_getFilterLogs","params":["0x10"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545

//curl --data '{"method":"eth_newFilter","params":[{"address":"0xA7033aa34c3c98F915E6cE60E70950f4d21C25fb","topics":["0xe7fd42cfef56f1caa721e298bd9ae0bc5ee36e3efb5f960a9f4ad0770a6ffc1e"]}],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545

type Filter struct {
	Method string `json:"method"`
	Params  []struct{
		FromBlock string `json:"fromBlock"`
		ToBlock string `json:"toBlock"`
		Address string `json:"address"`
		Topics []string `json:"topics"`
	} `json:"params"`
	ID int `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type Log struct {
	Method string `json:"method"`
	Params string `json:"params"`
	ID int `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Result string `json:"result"`
	Id string `json:"id"`
}

var(
	URL = "http://localhost:8545"
	FROM = "0x0"
	TO = "latest"
	ADDRESS = ""
	TOPICS = []string{""}
)

func eth_newFilter(url string,from string, to string, address string, topics []string) uint8{

	var fid Response

	data := &Filter{
		Method : "eth_newFilter",
		Params :[]struct{
			FromBlock string `json:"fromBlock"`
			ToBlock string `json:"toBlock"`
			Address string `json:"address"`
			Topics []string`json:"topics"`
		}{{
			FromBlock: from,
			ToBlock: to,
			Address: address,
			Topics:  topics,
		}},
		ID : 1,
		Jsonrpc : "2.0",
	}

	payloadBytes, err := json.Marshal(*data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(payloadBytes)

	fmt.Println(string(payloadBytes))

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type","application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
	json.Unmarshal(res, &fid)

	return fid.Result[0]

}

func eth_getFilterLogs(url string, filterId string) string{

	data := &Log{
		Method: "eth_getFilterLogs",
		Params: filterId,
		ID : 1,
		Jsonrpc : "2.0",
	}

	payloadBytes, err := json.Marshal(*data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(payloadBytes)

	fmt.Println(string(payloadBytes))

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type","application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)

	return string(res)
}

func main(){

	TOPICS = []string{"0xe7fd42cfef56f1caa721e298bd9ae0bc5ee36e3efb5f960a9f4ad0770a6ffc1e"}

	fid := eth_newFilter(URL, FROM, TO, "0xA7033aa34c3c98F915E6cE60E70950f4d21C25fb", TOPICS)
  	fmt.Println(fid)
	res := eth_getFilterLogs(URL, string(fid))
	fmt.Println(res)
}
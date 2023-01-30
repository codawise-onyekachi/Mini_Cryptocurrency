package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"net/http"
)

const (
	protocol = "tcp"
	port     = ":8080"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Node struct {
	NodeID     string
	Addresses  []string
	Blockchain *Blockchain
	Wallets    []*Wallet
	mtx        sync.Mutex
}

func (node *Node) Listen() {
	ln, err := net.Listen(protocol, port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go node.HandleConnection(conn)
	}
}

func (node *Node) HandleConnection(conn net.Conn) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)
	for {
		var message Message
		err := decoder.Decode(&message)
		if err != nil {
			fmt.Println(err)
			break
		}
		node.HandleMessage(message, conn)
	}
}

func (node *Node) HandleMessage(message Message, conn net.Conn) {
	node.mtx.Lock()
	defer node.mtx.Unlock()
	switch message.Type {
	case "transaction":
		var transaction Transaction
		err := json.Unmarshal(message.Data, &transaction)
		if err != nil {
			fmt.Println(err)
			return
		}
		node.Blockchain.AddTransaction(transaction)
	case "block":
		var block Block
		err := json.Unmarshal(message.Data, &block)
		if err != nil {
			fmt.Println(err)
			return
		}
		node.Blockchain.AddBlock(block)
	}
}

func (node *Node) BroadcastTransaction(transaction Transaction) {
	message := Message{Type: "transaction", Data: transaction}
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range node.Addresses {
		conn, err := net.Dial(protocol, address)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

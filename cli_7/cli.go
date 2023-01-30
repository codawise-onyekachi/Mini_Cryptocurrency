package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	Node *Node
}

func (cli *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "balance":
			cli.GetBalance(args[1:])
		case "create":
			cli.CreateWallet(args[1:])
		case "list":
			cli.ListWallets(args[1:])
		case "send":
			cli.Send(args[1:])
		case "help":
			cli.Help(args[1:])
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			cli.Help(nil)
		}
	}
}

func (cli *CLI) GetBalance(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: balance [address]")
		return
	}
	address := args[0]
	balance := 0
	for _, block := range cli.Node.Blockchain.Blocks {
		for _, transaction := range block.Transactions {
			if transaction.From == address {
				balance -= transaction.Amount
			}
			if transaction.To == address {
				balance += transaction.Amount
			}
		}
	}
	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

func (cli *CLI) CreateWallet(args []string) {
	wallet := NewWallet()
	cli.Node.Wallets = append(cli.Node.Wallets, wallet)
	fmt.Printf("Address: %s\n", wallet.Address)
}

func (cli *CLI) ListWallets(args []string) {
	for i, wallet := range cli.Node.Wallets {
		fmt.Printf("%d. Address: %s\n", i, wallet.Address)
	}
}

func (cli *CLI) Send(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: send [from_address] [to_address] [amount]")
		return
	}
	from := args[0]
	to := args[1]
	amount := args[2]
	for _, wallet := range cli.Node.Wallets {
		if wallet.Address == from {
			trans
		}
	}
}

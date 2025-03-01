package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// Transaction structure
type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Timestamp time.Time
}

// Block structure
type Block struct {
	Index        int
	Timestamp    time.Time
	PreviousHash string
	Hash         string
	Transactions []Transaction
}

// Blockchain slice
var Blockchain []Block

// Function to create a new transaction
func NewTransaction(receiver string, amount float64) Transaction {
	return Transaction{
		Sender:    "0x6187134381668E7B24F5B9f32aA3bc0EB9485b82", // Your MetaMask Address as Sender
		Receiver:  receiver,
		Amount:    amount,
		Timestamp: time.Now(),
	}
}

// Function to create a new block
func NewBlock(prevHash string, transactions []Transaction) Block {
	block := Block{
		Index:        len(Blockchain) + 1,
		Timestamp:    time.Now(),
		PreviousHash: prevHash,
		Transactions: transactions,
	}
	block.Hash = CalculateHash(block)
	return block
}

// Function to calculate the hash of a block
func CalculateHash(block Block) string {
	data, _ := json.Marshal(block)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func main() {
	fmt.Println("🚀 Welcome to Alimun Blockchain!")

	// Example: Create and add a transaction
	tx := NewTransaction("0xReceiverAddressHere", 10.0)
	fmt.Printf("✅ New Transaction Created: %+v\n", tx)

	// Create the first block (Genesis Block)
	genesisBlock := NewBlock("0", []Transaction{tx})
	Blockchain = append(Blockchain, genesisBlock)

	fmt.Printf("📌 Genesis Block Created: %+v\n", genesisBlock)
}

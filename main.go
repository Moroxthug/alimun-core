package main

import (
    "crypto/sha256"
    "fmt"
    "time"
)

// Block structure
type Block struct {
    Index        int
    Timestamp    time.Time
    PreviousHash string
    Hash         string
    Transactions []Transaction
}

// Transaction structure
type Transaction struct {
    Sender    string
    Receiver  string
    Amount    float64
    Timestamp time.Time
}

// Blockchain slice
var Blockchain []Block

func main() {
    fmt.Println("Welcome to Alimun Blockchain!")

}


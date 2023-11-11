package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

type Tx struct {
	FromID string `json:"from"`
	ToID   string `json:"to"`
	Value  uint64 `json:"value"`
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	tx := Tx{
		FromID: "Bill",
		ToID:   "Aaron",
		Value:  1000,
	}

	privateKey, err := crypto.LoadECDSA("zblock/accounts/kennedy.ecdsa")
	if err != nil {
		return fmt.Errorf("Failed to load private key: %v", err)
	}

	data, err := json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("Failed to marshal tx: %v", err)
	}

	v := crypto.Keccak256(data)

	sig, err := crypto.Sign(v, privateKey)
	if err != nil {
		return fmt.Errorf("Failed to sign tx: %v", err)
	}

	fmt.Println(hexutil.Encode(sig))

	return nil
}

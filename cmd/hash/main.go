package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run ./cmd/hash <password>")
		return
	}
	h := sha256.Sum256([]byte(os.Args[1]))
	fmt.Println(hex.EncodeToString(h[:]))
}

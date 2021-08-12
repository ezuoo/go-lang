package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Block struct {
	HashValue string
	Value     string
}

func main() {
	data := "hello world"
	hash := sha256.New()

	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	block := Block{HashValue: mdStr, Value: "It's a block"}

	jsonBytes, err := json.Marshal(block)

	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println("json String ", jsonString)
	var newBlock Block
	newErr := json.Unmarshal(jsonBytes, &newBlock)

	if newErr != nil {
		panic(newErr)
	}

	fmt.Println(newBlock.HashValue, newBlock.Value)
	// mem 구조체 필드 엑세스
	// fmt.Println(newMem.Name, newMem.Age, newMem.Active)

}

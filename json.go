package main

import (
	"encoding/json"
	"fmt"
)

//Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	// Go data
	mem := Member{"ezuoo", 58, true}

	// JSON encoding
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)

	var newMem Member
	newErr := json.Unmarshal(jsonBytes, &newMem)
	if newErr != nil {
		panic(newErr)
	}

	// mem 구조체 필드 엑세스
	fmt.Println(newMem.Name, newMem.Age, newMem.Active)

}

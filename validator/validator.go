package validator

import (
	"encoding/json"
	"fmt"
)

const (
	VALID = "valid"
	ERROR = "error"
)

func Validate(input []byte) WebData {
	var payload Payload
	err := json.Unmarshal(input, &payload)
	if err != nil {
		fmt.Println(ERROR, err)
	}

	for k, v := range payload.Auths {
		//fmt.Println("Esto es una key: ", k)
		//fmt.Println("y esto su valor: ", v.Auth)
		//fmt.Println("y esto su valor: ", v.Auth)
		return WebData{input, k, v, k}
	}
	return WebData{}
}

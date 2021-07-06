package validator

import (
	"encoding/json"
	"fmt"
)

const (
	VALID = "valid"
	ERROR = "error"
)

func Validate(input []byte) string {
	var payload Payload
	err := json.Unmarshal(input, &payload)
	if err != nil {
		fmt.Println("error:", err)
	}

	for k, v := range payload.Auths {
		fmt.Println("Esto es una key: ", k)
		fmt.Println("y esto su valor: ", v.Auth)
		//fmt.Println("y esto su valor: ", v.Auth)
		return k
	}
	return ""
}

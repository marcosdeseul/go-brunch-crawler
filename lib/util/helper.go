package util

import (
	"encoding/json"
	"fmt"
)

func PrettyPrintMap(data map[string]interface{}) {
	prettier, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(prettier))
}

func PrettyPrintStruct(data interface{}) {
	fmt.Printf("%#v", data)
}

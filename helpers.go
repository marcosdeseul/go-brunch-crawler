package main

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(data map[string]interface{}) {
	prettier, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(prettier))
}

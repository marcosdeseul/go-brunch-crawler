package util

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data map[string]interface{}) {
	prettier, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(prettier))
}

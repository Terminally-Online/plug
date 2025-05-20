package utils

import (
	"encoding/json"
	"fmt"
)

// LogObject logs an object in prettified json with a message tag for better context
func LogObject(msg string, obj any) {
	// Marshal the object to a pretty-printed JSON string
	prettyJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Printf("LogObject error: %v\n", err)
		return
	}

	fmt.Printf("[%s] %s\n", msg, string(prettyJSON))
}

package main

import (
	"encoding/base64"
	 "fmt"
	)

	func reverseStr(bt []byte) string{
		for i, j := 0, len(bt)-1; i<j; i, j = i+1, j-1 {
			bt[i], bt[j] = bt[j], bt[i]
		}
		return string(bt)
	}

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	whatIsIt = reverseStr(sd)
	fmt.Println(whatIsIt)
}
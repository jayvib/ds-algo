package main

import (
	"crypto/sha1"
	"fmt"
)

func createHash(byteStr []byte) []byte {
	hasher := sha1.New()
	hasher.Write(byteStr)
	bytes := hasher.Sum(nil)
	return bytes
}

func main() {
	input := "Luffy Monkey"
	fmt.Printf("%x\n",createHash([]byte(input)))
}

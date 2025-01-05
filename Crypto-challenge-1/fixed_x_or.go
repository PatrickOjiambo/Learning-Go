/*
Fixed XOR
Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965
... should produce:

746865206b696420646f6e277420706c6179
*/
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	result, error3 := fixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if error3 != nil{
		return
	}
	fmt.Println("Result", result)
	fmt.Println("Answer := 746865206b696420646f6e277420706c6179")
}
func fixedXOR(firstString string, secondString string) (string, error) {
	firstBytes, error1 := hex.DecodeString(firstString)
	if error1 != nil {
		fmt.Println("Decoding hex broght an error", error1)
		return "", fmt.Errorf("Decoding hex broght an error")
	}
	secondBytes, error2 := hex.DecodeString(secondString)
	if error2 != nil {
		fmt.Println("Decoding hex broght an error", error2)
		return "", fmt.Errorf("Decoding hex broght an error")
	}
	if len(firstBytes) != len(secondBytes) {
		return "", fmt.Errorf("input strings must have the same length")
	}
	xorBytes := make([]byte, len(firstBytes))
	for i := range firstBytes {
		xorBytes[i] = firstBytes[i] ^ secondBytes[i]
	}
	return hex.EncodeToString(xorBytes), nil
}

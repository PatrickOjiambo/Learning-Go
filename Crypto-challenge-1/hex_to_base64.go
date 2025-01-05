/*
Convert hex to base64
The string:

49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
Should produce:

SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
*/
package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)
func main() {
	//Convert the hex string to binary (4 bits each)
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64string, error := hexToBase64(hexString)
	if error != nil{
		fmt.Println("Error occured while converting to base64", error)
		return
	}
	fmt.Println(base64string)

}
func hexToBase64(hexString string)(string, error) {
	bytes, error := hex.DecodeString(hexString)
	if error != nil{
		fmt.Println("An error occured", error)
		return "", nil
	}
	bas64String := base64.StdEncoding.EncodeToString(bytes)
	return bas64String, nil
}
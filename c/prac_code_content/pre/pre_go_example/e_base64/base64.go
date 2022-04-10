/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/22 11:21 上午
* Description:
 */
package main

// This syntax imports the encoding/base64 package with the b64 name instead of the default base64.
// It’ll save us some space below.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible base64. Here’s how to encode using the standard encoder.
	// The encoder requires a []byte so we convert our string to that type.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoding may return an error, which you can check if you don’t already know the input to be well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64 format.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	// output:
	// YWJjMTIzIT8kKiYoKSctPUB+
	// abc123!?$*&()'-=@~

	// YWJjMTIzIT8kKiYoKSctPUB-
	// abc123!?$*&()'-=@~

	// The string encodes to slightly different values with the standard and URL base64 encoders
	// (trailing + vs -) but they both decode to the original string as desired.

}

/*
Question:

Suppose you are told that the one time pad encryption of the message "attack at
dawn" is 6c73d5240a948c86981bc294814d (the plaintext letters are encoded as
8-bit ASCII and the given ciphertext is written in hex). What would be the one
time pad encryption of the message "attack at dusk" under the same OTP key?

Solution:

	$ otp "attack at dusk"
	c2 = 6c73d5240a948c86981bc2808548
*/
package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

const (
	m1 = "attack at dawn"
	c1 = "6c73d5240a948c86981bc294814d"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: otp msg")
		os.Exit(2)
	}
	m2 := []byte(os.Args[1])

	if len(m2) != len(m1) {
		fmt.Fprintf(os.Stderr, "msg must have len=%d\n", len(m1))
		os.Exit(1)
	}

	bc1, err := hex.DecodeString(c1)
	if err != nil {
		panic(err)
	}

	if len(bc1) != len(m1) {
		panic("bc1 and m1 must have the same length")
	}

	otp := make([]byte, len(m1))
	for i := range otp {
		otp[i] = m1[i] ^ bc1[i]
	}

	bc2 := make([]byte, len(m2))
	for i := range bc2 {
		bc2[i] = m2[i] ^ otp[i]
	}

	c2 := hex.EncodeToString(bc2)
	fmt.Println("c2 =", c2)
}

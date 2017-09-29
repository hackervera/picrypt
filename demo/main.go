package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/tjgillies/picrypt"
)

func main() {
	r := rand.New(rand.NewSource(0))
	rInt := r.Intn(10000)
	str := "Test string"
	encrypted, err := picrypt.Encrypt([]byte(str), rInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encrypted '%v' as %v\n", str, encrypted)
	decrypted, err := picrypt.Decrypt(encrypted, rInt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decrypted %v as '%v'\n", encrypted, string(decrypted))
}

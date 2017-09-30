package picrypt

import (
	"fmt"
	"log"
	"math/rand"
)

func Example() {
	r := rand.New(rand.NewSource(0))
	rInt := r.Intn(10000)
	str := "Test string"
	encrypted, err := Encrypt([]byte(str), rInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encrypted)

	decrypted, err := Decrypt(encrypted, rInt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decrypted))

	// Output:
	// [18 69 16 18 58 7 58 69 4 36 58 7 58 69 58 4 16 34 16 12 16 58]
	// Test string
}

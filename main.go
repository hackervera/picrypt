package picrypt

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/claygod/PiHex"
)

func discover(b []byte, s int) int {
	pi := PiHex.New()
	var result []byte
	for {

		result = pi.Get(s, len(b))
		if result == nil {
			return 0
		}
		if bytes.Compare(result, b) == 0 {
			return s
		}
		s++
	}
}

func bytesToNibbles(buf []byte) ([]byte, error) {
	s := hex.EncodeToString(buf)
	b := make([]byte, len(s))
	for i, c := range s {
		h, err := strconv.ParseInt(string(c), 16, 0)
		if err != nil {
			return nil, err
		}
		b[i] = []byte(string(h))[0]
	}
	return b, nil
}

func nibblesToBytes(buf []byte) string {
	var results string
	for _, b := range buf {
		h := fmt.Sprintf("%x", b)
		results += h
	}
	return results
}

// Encrypt takes a byte slice and a shared secret int and encrypts each nibble against the index of pi
func Encrypt(buf []byte, rInt int) ([]byte, error) {
	b, err := bytesToNibbles(buf)
	if err != nil {
		return nil, err
	}
	var results []byte
	for _, chunk := range b {
		result := discover([]byte{chunk}, rInt)
		results = append(results, byte(result-rInt))
	}
	return results, nil
}

// Decrypt takes a byte slice encrypted with Encrypt and a shared secret int, and returns original byte slice
func Decrypt(buf []byte, rInt int) ([]byte, error) {
	pi := PiHex.New()
	var results []byte
	for _, r := range buf {
		idx := rInt + int(r)
		result := pi.Get(idx, 1)
		results = append(results, result[0])
	}
	s := nibblesToBytes(results)
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

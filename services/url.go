package services

import (
	"crypto/rand"
	"math/big"
)

func GenerateStringRandom(size int) string {
	result := ""
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < size; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(62))
		result += string(characters[index.Int64()])
	}
	return result
}

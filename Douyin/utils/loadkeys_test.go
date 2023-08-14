package utils

import (
	"fmt"
	"testing"
)

func TestLoadKeys(t *testing.T) {
	keys, err := loadKeys()
	if err != nil {
		fmt.Println(keys.PublicKey)
		fmt.Println(keys.PrivateKey)
	}
}

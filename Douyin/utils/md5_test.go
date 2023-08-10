package utils

import (
	"fmt"
	"testing"
)

func TestMd5Encryption(t *testing.T) {
	err := Md5Encryption("123456")
	fmt.Println(err)
}

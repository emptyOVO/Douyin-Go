package utils

import (
	"fmt"
	"testing"
)

func TestJwt(t *testing.T) {
	token, err := GenerateToken("feige", 1)
	if err != nil {
		fmt.Println(token)
	}

	_, c, err := ParseToken(token)
	if err != nil {
		fmt.Println(t)
		fmt.Println(c.UserId)
		fmt.Println(c.UserName)
		fmt.Println(c.ExpiresAt)
		fmt.Println(c.IssuedAt)
		fmt.Println(c.Subject)
		fmt.Println(c.UserId)
	}

}

package Controler

import (
	"crypto/rand"
	"fmt"
)
func TokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

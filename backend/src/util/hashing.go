package util

import (
	"crypto/sha512"
	"fmt"
)

func HashPassword(stringPassword string) string {
	h := sha512.New()
	h.Write([]byte(stringPassword))

	return fmt.Sprintf("%x", h.Sum(nil))
}

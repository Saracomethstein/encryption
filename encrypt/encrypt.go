package encrypt

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Encrypt256(promt string) string {
	hash := sha256.Sum256([]byte(promt))
	return fmt.Sprintf("%x", hash)
}

func Encrypt384(promt string) string {
	hash := sha512.Sum384([]byte(promt))
	return fmt.Sprintf("%x", hash)
}

func Encrypt512(promt string) string {
	hash := sha512.Sum512([]byte(promt))
	return fmt.Sprintf("%x", hash)
}

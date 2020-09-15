package main
import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)
func main() {
	a := "hello world"
	digest := sha256.Sum256([]byte(a))
	fmt.Println(digest)
	encoded := base64.StdEncoding.EncodeToString(digest[:])
	fmt.Println(encoded)
}
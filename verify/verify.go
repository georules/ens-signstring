package main

import (
	"math/big"
	"encoding/asn1"
	"flag"
	"encoding/base64"
	"fmt"
	"strings"
)

var (
	ieSig = flag.String("ieee1363", "", "IEEE 1361 signature")
)

func main() {
	flag.Parse()
	s := unpad(convertToURLEncoding(*ieSig))
	signature, err := base64.RawURLEncoding.DecodeString(s)
	signature, err = convert1363ToAsn1(signature)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(signature))
}

func unpad(s string) string {
	return strings.TrimRight(s, "=")
}

func convertToURLEncoding(s string) string {
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "/", "_")
	return s
}

func convert1363ToAsn1(b []byte) ([]byte, error) {
	rs := struct {
		R, S *big.Int
	}{
		R: new(big.Int).SetBytes(b[:len(b)/2]),
		S: new(big.Int).SetBytes(b[len(b)/2:]),
	}

	return asn1.Marshal(rs)
}
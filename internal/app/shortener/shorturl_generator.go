package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func GenerateShortLink(initialLink string) (string, error) {
	urlHashBytes := sha256Of(initialLink + time.Now().Format("2006-01-02 15:04:05.000000000"))
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return "", err
	}
	return finalString[:10], nil
}

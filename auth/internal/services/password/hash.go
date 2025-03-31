package password

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/scrypt"
)

const (
	iterationCount   = 1 << 13
	memoryCountForOp = 8
	parralelism      = 1
	passlen          = 32
)

func (s *service) hashPassword(pass string, salt string) (string, error) {
	hexHash, err := scrypt.Key(
		[]byte(pass),
		[]byte(salt),
		iterationCount,
		memoryCountForOp,
		parralelism,
		passlen,
	)
	if err != nil {
		return "", err
	}

	hash := hex.EncodeToString(hexHash)

	return hash, nil
}

func (s *service) generateSalt() string {
	salt := gofakeit.MinecraftAnimal()

	hash := sha256.New()
	hash.Write([]byte(salt))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	if len(hashString) > 64 {
		hashString = hashString[64:]
	}

	return hashString
}

package encryption

import (
	"golang.org/x/crypto/bcrypt"
)

func Encryption(key string) ([]byte, error) {

	bs, err := bcrypt.GenerateFromPassword([]byte(key), 8)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func CompareKey(bKey []byte, key string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(bKey, []byte(key))
	if err != nil {
		return false, err
	}
	return true, nil
}

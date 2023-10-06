package crypto

import "golang.org/x/crypto/bcrypt"

func GenereteHash(payload string) (hashed string, err error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashed = string(hash)
	return
}

func CompareHash(hash, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return
	}
	return
}

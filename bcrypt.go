package omnibus

import (
	"golang.org/x/crypto/bcrypt"
)

var ErrMismatchedHashAndPassword = bcrypt.ErrMismatchedHashAndPassword

func GenerateBcryptFromPassword(password string) (string, error) {
	r, err := bcrypt.GenerateFromPassword(Str2Bytes(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return Bytes2Str(r), nil
}

func CompareBcryptHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(Str2Bytes(hashedPassword), Str2Bytes(password))
}

func PasswordHash(password string) string {
	hash, _ := GenerateBcryptFromPassword(password)
	return hash
}

func PasswordVerify(password, hash string) bool {
	err := CompareBcryptHashAndPassword(hash, password)
	return err == nil
}

func PasswordHash(password string) string {
	hash, _ := GenerateBcryptFromPassword(password)
	return hash
}

func PasswordVerify(password, hash string) bool {
	err := CompareBcryptHashAndPassword(hash, password)
	return err == nil
}

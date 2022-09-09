package omnibus

import (
	"bytes"

	"golang.org/x/crypto/bcrypt"
)

var ErrMismatchedHashAndPassword = bcrypt.ErrMismatchedHashAndPassword

func GenerateBcryptFromPassword(password string) (string, error) {
	var buf bytes.Buffer
	buf.WriteString(password)
	r, err := bcrypt.GenerateFromPassword(buf.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	var buf2 bytes.Buffer
	buf2.Write(r)
	return buf2.String(), nil
}

func CompareBcryptHashAndPassword(hashedPassword, password string) error {
	var buf bytes.Buffer
	buf.WriteString(hashedPassword)

	var buf2 bytes.Buffer
	buf.WriteString(password)

	return bcrypt.CompareHashAndPassword(buf.Bytes(), buf2.Bytes())
}

func PasswordHash(password string) string {
	hash, _ := GenerateBcryptFromPassword(password)
	return hash
}

func PasswordVerify(password, hash string) bool {
	err := CompareBcryptHashAndPassword(hash, password)
	return err == nil
}

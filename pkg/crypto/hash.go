package crypto

import "golang.org/x/crypto/bcrypt"

func HashString(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashed), err
}

func DoMatch(enc, raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(enc), []byte(raw))
	return err == nil
}

package security

import (
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type PasswordEncoder interface {
	Encode(raw string) string
	Matches(raw string, encoded string) bool
}

type bcryptpasswordencoder struct {
}

var once sync.Once
var passwordEncoder PasswordEncoder

func NewBcryptPasswordEncoder() PasswordEncoder {
	once.Do(func() {
		passwordEncoder = &bcryptpasswordencoder{}
	})
	return passwordEncoder
}

func (b *bcryptpasswordencoder) Encode(rawPass string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (b *bcryptpasswordencoder) Matches(raw string, encoded string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encoded), []byte(raw))
	return err == nil
}

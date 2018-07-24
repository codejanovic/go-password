package go_password

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Password that represents a hash and a salt
type BcryptPassword struct {
	hashedPassword string
	salt           string
}

// NewBcryptPassword will create a new bcrypt password with default cost
func NewBcryptPassword(plainPassword string) (Password, error) {
	salt := uuid.Must(uuid.NewV4()).String()
	return NewBcryptPasswordWithGivenSalt(plainPassword, salt)
}

// NewBcryptPassword will create a new bcrypt password based on the given hash with default cost
func NewBcryptPasswordWithGivenSalt(plainPassword string, salt string) (Password, error) {
	password := new(BcryptPassword)

	password.salt = salt
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword+password.salt), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	password.hashedPassword = string(hash)

	return password, nil
}

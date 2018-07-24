package go_password

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Password that represents a hash and a salt
type BcryptPassword struct {
	hash string
	salt string
}

func (p *BcryptPassword) GetHash() string {
	return p.hash
}

func (p *BcryptPassword) GetSalt() string {
	return p.salt
}

func (p *BcryptPassword) Compare(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(saltedPassword(password, p.salt)))
	return err == nil
}

func saltedPassword(password string, salt string) []byte {
	return []byte(password + salt)
}


// NewBcryptPassword will create a new bcrypt password with default cost
func NewBcryptPassword(plainPassword string) (Password, error) {
	salt := uuid.Must(uuid.NewV4()).String()
	return NewBcryptPasswordWithGivenSalt(plainPassword, salt)
}

// NewBcryptPassword will create a new bcrypt password based on the given hash with default cost
func NewBcryptPasswordWithGivenSalt(plainPassword string, salt string) (Password, error) {
	var password = new(BcryptPassword)

	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword+salt), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	password.hash = string(hash)
	password.salt = salt

	return password, nil
}

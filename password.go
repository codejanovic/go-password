package go_password

type Password interface {
	GetHash() string
	GetSalt() string
}

func (p BcryptPassword) GetHash() string {
	return p.hashedPassword
}

func (p BcryptPassword) GetSalt() string {
	return p.salt
}

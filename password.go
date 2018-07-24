package go_password

type Password interface {
	GetHash() string
	GetSalt() string
	Compare(password string) bool
}


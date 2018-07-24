# go-password

this projects provides a very simple password interface and different password implementations.

Every implementation will return a `Password` interface:

```go
type Password interface {
	GetHash() string
	GetSalt() string
}
```

Supported implementations ...

### Bcrypt
Either create a new Password with a generated (UUID) salt:
```go
func NewBcryptPassword(plainPassword string) (Password, error) {
	...
}
```

or create a new Password with a given salt of your choice:
```go
func NewBcryptPasswordWithGivenSalt(plainPassword string, salt string) (Password, error) {
	...
}
```


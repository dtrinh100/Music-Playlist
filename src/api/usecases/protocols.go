package usecases

const (
	AppFaultErr = iota
	UserFaultErr
)

type M map[string]interface{}

type Logger interface {
	Log(message string) error
}

type UserRepository interface {
	OneByEmail(userEmail string) (*User, error)
	OneByUsername(username string) (*User, error)
	All() ([]User, error)
	Create(user *User) error
	Update(userEmail string, changes M) error
	Delete(userEmail string) error
	ComparePassword(hashedPass, clearTextPass []byte) error
}

type MPError interface {
	Status() int
	Error() string
}

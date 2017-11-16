package usecases

type M map[string]interface{}

type Logger interface {
	Log(message string) error
}

type UserRepository interface {
	One(userEmail string) (*User, error)
	All() ([]User, error)
	Create(user *User) error
	Update(userEmail string, changes M) error
	Delete(userEmail string) error
}

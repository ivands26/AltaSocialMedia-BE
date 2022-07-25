package domain

type User struct {
	ID       int
	Nama     string
	Username string
	Email    string
	Password string
	No_HP    string
}

type UserUseCases interface {
	Register(newUser User) (User, error)
	GetSpecificUser(userId int) (User, error)
}

type UserData interface {
	Register(newUser User) (User, error)
	GetSpecificUser(userId int) (User, error)
}

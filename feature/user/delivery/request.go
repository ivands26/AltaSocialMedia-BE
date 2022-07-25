package delivery

import "github.com/AltaProject/AltaSocialMedia/domain"

type RegisterFormat struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	No_HP    string `json:"no_hp"`
}

func (r *RegisterFormat) ToModel() domain.User {
	return domain.User{
		Nama:     r.Nama,
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		No_HP:    r.No_HP,
	}
}

type LoginFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

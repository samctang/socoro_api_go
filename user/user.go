package user

type User struct {
	Id       int    `json:"id"`
	Uuid     string `json:"uuid"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, passwd"`
	Role     int    `json:"role" validate:"required"`
}

type DTO struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, passwd"`
	Role     int    `json:"role" validate:"required"`
	isAdmin  bool   `json:"isAdmin" validate:"required"`
}
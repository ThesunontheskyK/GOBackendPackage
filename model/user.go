package model


type User struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type CreateUser struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

package model

type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	Phone    string `json:"phone"`
}

type UserDTO struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

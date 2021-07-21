package model

type User struct {
	*Model
	Name  string `json:"name"`
	Pwd   string `json:"pwd"`
	Email string `json:"email"`
}

func (u User) TableName() string {
	return "user"
}

package entity

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"` // admin or customer
}

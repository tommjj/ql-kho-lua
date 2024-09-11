package domain

type Role string

const (
	Root  Role = "root"
	Staff Role = "staff"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
}

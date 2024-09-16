package domain

type Role string

const (
	Root   Role = "root"
	Member Role = "member"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     Role   `json:"role"`
}

func (u *User) RemovePass() {
	u.Password = ""
}
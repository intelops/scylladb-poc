package models

const (
	CreateUserTable = "id text PRIMARY KEY, firstname text , lastname text, password text , email text , phone text"
)

type User struct {
	Id        string `json:"id" db:"id"`
	FirstName string    `json:"firstname" db:"firstname"`
	LastName  string    `json:"lastname,omitempty" db:"lastname"`
	Password  string    `db:"password" json:"password"`
	Email     string    `db:"email" json:"email"`
	Phone     string    `db:"phone" json:"phone"`
}

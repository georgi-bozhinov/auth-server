package storage

type User struct {
	Username  string `db:"username"`
	Password  string `db:"password"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Email     string `db:"email"`
}

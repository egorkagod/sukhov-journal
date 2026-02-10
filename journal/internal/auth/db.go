package auth

type User struct {
	ID           int64
	Login        string
	PasswordHash string
	Role         string
}

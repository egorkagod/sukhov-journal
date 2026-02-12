package auth

type User struct {
	ID           uint64 `gorm:"primaryKey"`
	Login        string `gorm:"uniqueIndex"`
	PasswordHash string
	Role         string
}

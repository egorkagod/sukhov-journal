package articles

import (
	"time"

	"journal/internal/auth"
)

type Article struct {
	ID        uint64 `gorm:"primaryKey"`
	AuthorID  uint64
	User      auth.User `gorm:"foreignKey:AuthorID"`
	Title     string
	Body      string
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

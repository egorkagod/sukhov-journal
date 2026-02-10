package articles

import "time"


type Article struct {
	ID int64
	AuthorID int64
	Title string
	Body string
	CreatedAt time.Time
	UpdatedAt time.Time
}
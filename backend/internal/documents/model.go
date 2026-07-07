package documents

import "time"

type Document struct {
	ID               string
	UserID           string
	Title            string
	OriginalFilename string
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

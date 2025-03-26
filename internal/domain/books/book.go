package books

import "library-api-go/internal/domain/authors"

type Books struct {
	ID       int             `json:"id" gorm:"primary_key"`
	Title    string          `json:"title" gorm:"not null"`
	AuthorID int             `json:"Author_ID" gorm:"not null"`
	Author   authors.Authors `json:"author" gorm:"foreignkey:AuthorID"`
}
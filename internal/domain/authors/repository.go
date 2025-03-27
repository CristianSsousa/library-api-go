package authors

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	FindAll(ctx context.Context) ([]Authors, error)
	FindById(ctx context.Context, id int) (Authors, error)
	Create(ctx context.Context, author Authors) (Authors, error)
	Update(ctx context.Context, author Authors) (Authors, error)
	Delete(ctx context.Context, author Authors) error
}

type AuthorRepository struct {
	db *gorm.DB
}


func NewAuthorsRepository(db *gorm.DB) RepositoryInterface {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) FindAll(ctx context.Context) ([]Authors, error) {
	var authors []Authors
	r.db.Find(&authors)
	return authors, nil
}

func (r *AuthorRepository) FindById(ctx context.Context, id int) (Authors, error) {
	var author Authors
	r.db.First(&author, id)
	return author, nil
}

func (r *AuthorRepository) Create(ctx context.Context, author Authors) (Authors, error) {
	r.db.Create(&author)
	return author, nil
}

func (r *AuthorRepository) Update(ctx context.Context, author Authors) (Authors, error) {
	r.db.Save(&author)
	return author, nil
}

func (r *AuthorRepository) Delete(ctx context.Context, author Authors) error {
	r.db.Delete(&author)
	return nil
}

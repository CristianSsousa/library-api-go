package authors

import "gorm.io/gorm"

type AuthorRepository struct {
	db *gorm.DB
}


func NewAuthorsRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) FindAll() ([]Authors, error) {
	var authors []Authors
	r.db.Find(&authors)
	return authors, nil
}

func (r *AuthorRepository) FindById(id int) (Authors, error) {
	var author Authors
	r.db.First(&author, id)
	return author, nil
}

func (r *AuthorRepository) Create(author Authors) (Authors, error) {
	r.db.Create(&author)
	return author, nil
}

func (r *AuthorRepository) Update(author Authors) (Authors, error) {
	r.db.Save(&author)
	return author, nil
}

func (r *AuthorRepository) Delete(author Authors) error {
	r.db.Delete(&author)
	return nil
}
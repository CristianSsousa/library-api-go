package authors

type AuthorsService struct {
	authorsRepository *AuthorRepository
}

func NewAuthorsService(authorsRepository *AuthorRepository) *AuthorsService {
	return &AuthorsService{authorsRepository: authorsRepository}
}

func (s *AuthorsService) FindAll() ([]Authors, error) {
	return s.authorsRepository.FindAll()
}

func (s *AuthorsService) FindById(id int) (Authors, error) {
	return s.authorsRepository.FindById(id)
}

func (s *AuthorsService) Create(author Authors) (Authors, error) {
	return s.authorsRepository.Create(author)
}

func (s *AuthorsService) Update(author Authors) (Authors, error) {
	return s.authorsRepository.Update(author)
}

func (s *AuthorsService) Delete(author Authors) error {
	return s.authorsRepository.Delete(author)
}
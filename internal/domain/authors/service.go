package authors

import "context"

type ServiceInterface interface {
	FindAll(ctx context.Context) ([]Authors, error)
	FindById(ctx context.Context,id int) (Authors, error)
	Create(ctx context.Context, author Authors) (Authors, error)
	Update(ctx context.Context, author Authors) (Authors, error)
	Delete(ctx context.Context, author Authors) error
}

type AuthorsService struct {
	authorsRepository RepositoryInterface
}

func NewAuthorsService(authorsRepository RepositoryInterface) *AuthorsService {
	return &AuthorsService{authorsRepository: authorsRepository}
}

func (s *AuthorsService) FindAll(ctx context.Context) ([]Authors, error) {
	return s.authorsRepository.FindAll(ctx)
}

func (s *AuthorsService) FindById(ctx context.Context, id int) (Authors, error) {
	return s.authorsRepository.FindById(ctx, id)
}

func (s *AuthorsService) Create(ctx context.Context, author Authors) (Authors, error) {
	return s.authorsRepository.Create(ctx, author)
}

func (s *AuthorsService) Update(ctx context.Context, author Authors) (Authors, error) {
	return s.authorsRepository.Update(ctx, author)
}

func (s *AuthorsService) Delete(ctx context.Context, author Authors) error {
	return s.authorsRepository.Delete(ctx, author)
}
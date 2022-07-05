package user

import (
	"context"

	"github.com/witalok2/api-rest/adapter/secondary/postgres"
	"github.com/witalok2/api-rest/internal/entity"
)

type service struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) Service {
	return service{
		repository: r,
	}
}

func (s service) CreateUser(ctx context.Context, user entity.User) error {
	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s service) ListUser(ctx context.Context) ([]entity.User, error) {
	return s.repository.ListUser(ctx)
}

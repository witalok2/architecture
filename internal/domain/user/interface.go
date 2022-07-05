package user

import (
	"context"

	"github.com/witalok2/api-rest/internal/entity"
)

type Service interface {
	CreateUser(context.Context, entity.User) error
	ListUser(context.Context) ([]entity.User, error)
}

package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/api-rest/config"
	"github.com/witalok2/api-rest/internal/entity"
)

type Repository interface {
	CreateUser(context.Context, entity.User) error
	ListUser(ctx context.Context) ([]entity.User, error)
	Close()
}

type repository struct {
	db *sqlx.DB
}

func New(ps config.DB) Repository {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", ps.User, ps.Password, ps.Host, ps.Port, ps.Database, ps.SSL))
	if err != nil {
		logger.Fatal(err)
	}

	return &repository{
		db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user entity.User) error {
	_, err := r.db.Exec(sqlCreateUser, user.Name)
	if err != nil {
		logger.WithError(err).Error("failed on create user")
		return err
	}

	return nil
}

func (r *repository) ListUser(ctx context.Context) (users []entity.User, err error) {
	err = r.db.Select(&users, sqlCreateUserAddress)
	if err != nil {
		logger.WithError(err).Error("failed on create user address")
		return nil, err
	}

	return users, nil
}

func (r *repository) Close() {
	r.db.Close()
}

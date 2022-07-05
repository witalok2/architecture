package entity

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type service string

const SERVICE service = "service"

func (u *User) Validate() error {
	return validation.ValidateStruct(u, validation.Field(&u.Name, validation.Required))
}

type User struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deletedAt" db:"deleted_at"`
}

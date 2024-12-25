package repositories

import (
	"context"

	"github.com/KalinduGandara/erp-system/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	List(ctx context.Context) ([]*entities.User, error)
}

type CustomerRepository interface {
	Create(ctx context.Context, customer *entities.Customer) error
	Update(ctx context.Context, customer *entities.Customer) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*entities.Customer, error)
	List(ctx context.Context) ([]*entities.Customer, error)
}

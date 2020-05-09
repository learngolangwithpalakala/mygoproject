package user

import (
	"context"

	"github.com/learngolangwithpalakala/mygoproject/models"
)

// Usecase represent the user's usecases
type Usecase interface {
	Store(context.Context, *models.User) error
	Update(ctx context.Context, u *models.User) error
	Delete(ctx context.Context, empNumber string) error
	GetByEmpNumber(ctx context.Context, empNumber string) (models.User, error)
	Fetch(ctx context.Context) ([]models.User, error)
	Login(ctx context.Context, u models.User) (models.User, error)
}

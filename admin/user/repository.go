package user

import (
	"context"

	"github.com/learngolangwithpalakala/mygoproject/models"
)

// Repository represent the User's repository contract
type Repository interface {
	Fetch(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmpNumber(ctx context.Context, empNumber string) (models.User, error)
	Update(ctx context.Context, ar *models.User, id int) error
	Store(ctx context.Context, a *models.User) (int, error)
	Delete(ctx context.Context, a *models.User) error
	FindAll() (res []*models.User, err error)
	GetByUserName(ctx context.Context, userName string) (models.User, error)
}

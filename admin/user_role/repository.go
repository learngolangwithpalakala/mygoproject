package user_role

import (
	"github.com/learngolangwithpalakala/mygoproject/models"
)

// Repository represent the role's repository contract
type Repository interface {
	FindAll() ([]models.UserRole, error)
	Store(ur *models.UserRole) error
	FindUserRoleByUserIdAndRoleId(userId int, roleId int) (models.UserRole, error)
	Delete(ur *models.UserRole) error
}

package role

import (
	"github.com/learngolangwithpalakala/mygoproject/models"
)

// Repository represent the role's repository contract
type Repository interface {
	FindAll() ([]models.Role, error)
	FindByRoleName(roleName string) (models.Role, error)
}

package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/learngolangwithpalakala/mygoproject/admin/role"
	"github.com/learngolangwithpalakala/mygoproject/models"
)

type mysqlRoleRepo struct {
	DB *gorm.DB
}

// NewMysqlRoleRepository will create an object that represent the role.Repository interface
func NewMysqlRoleRepository(db *gorm.DB) role.Repository {
	return &mysqlRoleRepo{
		DB: db,
	}
}

func (m *mysqlRoleRepo) FindAll() ([]models.Role, error) {
	db := m.DB
	var roles []models.Role
	err := db.Find(&roles).Error
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("{}", roles)
	return roles, nil
}

func (m *mysqlRoleRepo) FindByRoleName(roleName string) (models.Role, error) {
	var role models.Role
	err := m.DB.Where("role = ?", roleName).Find(&role).Error
	if err != nil {

	}
	return role, nil
}

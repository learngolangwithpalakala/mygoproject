package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/learngolangwithpalakala/mygoproject/admin/user_role"
	"github.com/learngolangwithpalakala/mygoproject/models"
)

type mysqlUserRoleRepo struct {
	DB *gorm.DB
}

func NewMysqlUserRoleRepository(db *gorm.DB) user_role.Repository {
	return &mysqlUserRoleRepo{
		DB: db,
	}
}

func (m *mysqlUserRoleRepo) FindAll() ([]models.UserRole, error) {
	db := m.DB
	var userRoles []models.UserRole
	err := db.Find(&userRoles).Error
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("{}", userRoles)
	return userRoles, nil
}

func (m *mysqlUserRoleRepo) Store(ur *models.UserRole) error {
	db := m.DB
	db.Create(&ur)
	return nil
}

func (m *mysqlUserRoleRepo) FindUserRoleByUserIdAndRoleId(userId int, roleId int) (models.UserRole, error) {
	var userRole models.UserRole
	err := m.DB.Where("user_id = ? AND role_id = ?", userId, roleId).Find(&userRole).Error
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("{}", userRole)
	return userRole, nil

}

func (m *mysqlUserRoleRepo) Delete(ur *models.UserRole) error {
	err := m.DB.Where("id = ?", ur.ID).Delete(&ur).Error
	if err != nil {
		return errors.New("user role failed to delete")
	}
	return nil
}

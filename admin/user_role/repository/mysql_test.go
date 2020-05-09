package repository_test

import (
	"testing"

	userRoleRepo "github.com/learngolangwithpalakala/mygoproject/admin/user_role/repository"
	"github.com/learngolangwithpalakala/mygoproject/config"
	"github.com/learngolangwithpalakala/mygoproject/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	db, err := config.GetDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	r := userRoleRepo.NewMysqlUserRoleRepository(db)
	userRolesList, err := r.FindAll()
	if err != nil {

	}
	assert.NotEmpty(t, userRolesList)
}

func TestStoreUserRole(t *testing.T) {
	db, err := config.GetDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	r := userRoleRepo.NewMysqlUserRoleRepository(db)
	r.Store(&models.UserRole{UserId: 4, RoleId: 3})
}

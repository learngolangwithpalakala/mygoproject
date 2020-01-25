package repository_test

import (
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/stretchr/testify/assert"
	//sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"

	roleRepo "github.com/bxcodec/go-clean-arch/admin/role/repository"
)
/*
func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow(1, "Iman Tumorang", time.Now(), time.Now())

	query := "SELECT id, name, created_at, updated_at FROM author WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	a := repository.NewMysqlAuthorRepository(db)

	anArticle, err := a.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
*/

func TestFindAll(t *testing.T) {
	//db, mock, err := sqlmock.New()
   db, err := config.GetDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	r := roleRepo.NewMysqlRoleRepository(db)
	rolesList ,err := r.FindAll()
	if err != nil {

	}
	assert.NotEmpty(t,rolesList)
}
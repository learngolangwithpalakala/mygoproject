package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
	"github.com/learngolangwithpalakala/mygoproject/admin/user"
	//"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/learngolangwithpalakala/mygoproject/models"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

// NewMysqlUserRepository will create an object that represent the article.Repository interface
func NewMysqlUserRepository(db *gorm.DB) user.Repository {
	return &mysqlUserRepository{
		DB: db,
	}
}

func (m *mysqlUserRepository) FindAll() ([]*models.User, error) {
	db := m.DB
	var users []*models.User
	err := db.Find(&users).Error
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("{}", users)
	return users, nil
}

func (m *mysqlUserRepository) Fetch(ctx context.Context) ([]models.User, error) {
	db := m.DB
	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("{}", users)
	return users, nil
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id int64) (res *models.User, err error) {

	return nil, nil
}

func (m *mysqlUserRepository) GetByEmpNumber(ctx context.Context, empNumber string) (res models.User, err error) {
	var user models.User
	err = m.DB.Where("emp_number = ?", empNumber).Find(&user).Error
	if err != nil {

	}
	return user, nil
}

func (m *mysqlUserRepository) GetByUserName(ctx context.Context, userName string) (res models.User, err error) {
	var user models.User
	err = m.DB.Where("work_email = ?", userName).Find(&user).Error
	if err != nil {
		return user, errors.New("unable to find user")
	}
	return user, nil
}

func (m *mysqlUserRepository) Store(ctx context.Context, u *models.User) (int, error) {
	db := m.DB
	err := db.Create(&u).Error
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (m *mysqlUserRepository) Delete(ctx context.Context, u *models.User) error {
	db := m.DB
	err := db.Unscoped().Delete(&u).Error
	if err != nil {
		return errors.New("user Db deletion failed")
	}
	return nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, u *models.User, id int) error {
	db := m.DB
	var user models.User
	err := db.Model(&user).Update(models.User{ID: id, UserName: u.UserName, EmpNumber: u.EmpNumber,
		AboutUser: u.AboutUser,
		City:      u.City, Password: u.Password, Country: u.Country, FirstName: u.FirstName,
		HomePhoneNumber: u.HomePhoneNumber, InsuranceNumber: u.InsuranceNumber, LastName: u.LastName,
		PassportNumber: u.PassportNumber, PersonalEmail: u.PersonalEmail, PostCode: u.PostCode, Prefix: u.Prefix,
		Suffix: u.Suffix, ProjectEndDate: u.ProjectEndDate, ProjectName: u.ProjectName,
		ProjectStartDate: u.ProjectStartDate, RoleId: u.RoleId, SocialSecurityNumber: u.SocialSecurityNumber,
		HireDate: u.HireDate, EndDate: u.EndDate, Address: u.Address, TaxId: u.TaxId, Skills: u.Skills,
		WorkEmail: u.WorkEmail, WorkPhoneNumber: u.WorkPhoneNumber, Position: u.Position, Gender: u.Gender,
		BirthDay: u.BirthDay, Active: u.Active,
	}).Error
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}

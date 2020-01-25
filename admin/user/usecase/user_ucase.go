package usecase

import (
	"context"
	"errors"
	"github.com/bxcodec/go-clean-arch/admin/role"
	"github.com/bxcodec/go-clean-arch/admin/user"
	userRole "github.com/bxcodec/go-clean-arch/admin/user_role"
	"github.com/bxcodec/go-clean-arch/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

type userUsecase struct {
	userRepo    	user.Repository
	roleRepo        role.Repository
	userRoleRepo    userRole.Repository
	contextTimeout  time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUserUsecase(u user.Repository, r role.Repository,ur userRole.Repository,timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:      u,
		roleRepo:      r,
		userRoleRepo:   ur,
		contextTimeout: timeout,
	}
}

func (a *userUsecase) Store(c context.Context, m *models.User)  error{
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	roleId, ok := validateRole(a,m.RoleId)
	if !ok {
	  return errors.New("Invalid Role")
	}
	userModel , err := a.userRepo.GetByEmpNumber(ctx,m.EmpNumber)
	if userModel.ID > 0 {
         return errors.New("User already exists with this employee number")
	}
	m.RoleId = strconv.Itoa(roleId)

	byteArrPwd := []byte(m.Password)
	encryptedPwd := hashAndSalt(byteArrPwd)
	m.Password = encryptedPwd
	userId, err :=  a.userRepo.Store(ctx, m)
	if err != nil {
      return err
	}
	err =  a.userRoleRepo.Store(&models.UserRole{UserId:userId ,RoleId:roleId })
	if err != nil {
        return  err
	}
	return nil
}

func (a *userUsecase) Update(c context.Context, u *models.User) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	//u.UpdatedAt = time.Now()
	userModel , err := a.userRepo.GetByEmpNumber(ctx,u.EmpNumber)
	if err != nil {
		return errors.New("failed to find user information")
	}
	if userModel.ID > 0 {
		return a.userRepo.Update(ctx, u,userModel.ID)
	}
	 return errors.New("failed to update")
}

func validateRole(a *userUsecase ,roleName string) (int, bool) {
	role ,err := a.roleRepo.FindByRoleName(roleName)
	if err != nil {
        return 0,false
	}
	if role.ID  > 0 {
		return role.ID,true
	}
    return 0,false
}

func (a *userUsecase) Delete(c context.Context, empNumber string) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	existedUser, err := a.userRepo.GetByEmpNumber(ctx, empNumber)
	if err != nil {
		return err
	}
	roleId,e  := strconv.Atoi(existedUser.RoleId)
	if e!= nil {
		return e
	}
	existedUserRole , er := a.userRoleRepo.FindUserRoleByUserIdAndRoleId(existedUser.ID,roleId)
	if er != nil {
		return err
	}
    err = a.userRoleRepo.Delete(&existedUserRole)
	if err != nil {
		return err
	}
	return a.userRepo.Delete(ctx, &existedUser)
}

func (a *userUsecase) GetByEmpNumber(ctx context.Context, empNumber string) (models.User, error){
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	res, err := a.userRepo.GetByEmpNumber(ctx, empNumber)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (a *userUsecase) Fetch(c context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listUsers,  err := a.userRepo.Fetch(ctx)
	if err != nil {
		return listUsers, err
	}
	return listUsers,nil
}

func (a *userUsecase) Login(c context.Context, m models.User)  ( models.User,error){
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	var user models.User
	user , err := a.userRepo.GetByUserName(ctx,m.UserName)
	if err != nil {
		return user, errors.New("User name not found")
	}
   	if comparePasswords(user.Password , []byte(m.Password)) {
		return user, nil
	}
	return  user, errors.New("login failed")
}

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
package models
//import ("github.com/jinzhu/gorm")
type UserRole struct {
	//gorm.Model
	ID        int `json:"id";gorm:"primary_key";"AUTO_INCREMENT"`
	UserId    int `json:"user_id",gorm:"foreignkey:UserId;association_foreignkey:ID"`
	RoleId    int `json:"role_id",gorm:"foreignkey:RoleId;association_foreignkey:ID"`
}

func (u *UserRole) TableName() string {
	// custom table name, this is default
	return "user_role"
}
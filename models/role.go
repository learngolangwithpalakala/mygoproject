package models
//import ("github.com/jinzhu/gorm")
type Role struct {
	//gorm.Model
	ID         int 	`json:"ID";gorm:"primary_key";"AUTO_INCREMENT"`
	Role 	  string 	`json:"role"`
}

func (r *Role) TableName() string {
	// custom table name, this is default
	return "role"
}
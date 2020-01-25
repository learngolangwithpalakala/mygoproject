package models
//import ("github.com/jinzhu/gorm")
type User struct {
	//gorm.Model
	ID        				int     	  `json:"id",gorm:"primary_key";AUTO_INCREMENT"`
	UserName                string        `json:"userName;gorm:"column:user_name"`
	EmpNumber		        string        `json:"empNumber",omitempty;gorm:"column:emp_number"`
	AboutUser	            string        `json:"aboutUser;gorm:"column:about_user"`
	City					string        `json:"city;gorm:"column:city"`
	Password                string        `json:"password;gorm:"column:password"`
	Country	                string        `json:"country;gorm:"column:country"`
	FirstName               string        `json:"firstName;gorm:"column:first_name"`
	HomePhoneNumber	        string        `json:"homePhoneNumber;gorm:"column:home_phone_number"`
	InsuranceNumber         string 		  `json:"insuranceNumber;gorm:"column:insurance_number"`
	LastName			    string        `json:"lastName;gorm:"column:last_name"`
	PassportNumber          string        `json:"passportNumber;gorm:"column:passport_number"`
	PersonalEmail			string        `json:"personalEmail;gorm:"column:personal_email"`
	PostCode                string        `json:"postCode;gorm:"column:post_code"`
	Prefix                  string        `json:"prefix;gorm:"column:prefix"`
	Suffix                  string        `json:"suffix;gorm:"column:suffix"`
	ProjectEndDate          string        `json:"projectEndDate;gorm:"column:project_end_date"`
	ProjectName             string        `json:"projectName;gorm:"column:project_name"`
	ProjectStartDate        string        `json:"projectStartDate;gorm:"column:project_start_date"`
	RoleId                  string        `json:"roleId;gorm:"column:role_id"`
	SocialSecurityNumber    string        `json:"socialSecurityNumber;gorm:"column:social_security_number"`
	HireDate                string        `json:"hireDate;gorm:"column:hire_date"`
	EndDate                 string        `json:"endDate;gorm:"column:end_date"`
	Address                 string        `json:"address;gorm:"column:address"`
	TaxId                   string        `json:"taxId;gorm:"column:tax_id"`
	Skills                  string        `json:"skills;gorm:"column:skills"`
	WorkEmail               string        `json:"workEmail;gorm:"column:work_email"`
	WorkPhoneNumber         string        `json:"workPhoneNumber;gorm:"column:work_phone_number"`
	Position                string        `json:"position;gorm:"column:position"`
	Gender                  string        `json:"gender;gorm:"column:gender"`
	BirthDay                string        `json:"birthDay;gorm:"column:birth_day"`
	Active                  string        `json:"active;gorm:"column:active"`
}

func (u *User) TableName() string {
	// custom table name, this is default
	return "user"
}
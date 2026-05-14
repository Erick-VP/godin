package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model
	Role string
	CompanyID string
	Location string
	Remote bool
	Link string
	Salary int
}

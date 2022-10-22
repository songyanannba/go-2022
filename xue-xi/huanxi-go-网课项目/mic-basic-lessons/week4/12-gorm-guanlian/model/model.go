package model

import "gorm.io/gorm"

type Employer struct {
	gorm.Model
	Name string
	CompanyID int
	Company Company
}

type Company struct {
	gorm.Model
	Name string
}
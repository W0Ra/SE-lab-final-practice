package entity

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	StudentID string `valid:"required~Please insert student id, matches(^[BMD]\\d{7}$)"`
	FirstName string `valid:"required~Please insert your first name, maxstringlength(10)~First name must less than 10 character"`
	LastName  string `valid:"required~Please insert your last name, minstringlength(4)~Last name must be at least 4 character"`
	Email     string `valid:"required~Please insert your email, email~Invalid email address"`
	Phone     string `valid:"required~Please insert your phone number, stringlength(10|10)~Phone number must be 10 character"`
	Url       string `valid:"required~Please insert url, url~Invalid url link"`

	GenderID uint   `valid:"required~Please choose your gender"`
	Gender   Gender `gorm:"foreignKey:GenderID"`
}

type Gender struct {
	gorm.Model
	Name string
}

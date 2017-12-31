package models

import "time"

type User struct {
	ID               uint       `json:"id"`
	CreatedAt        time.Time  `json:"-"`
	UpdatedAt        time.Time  `json:"-"`
	DeletedAt        *time.Time `json:"-"`
	FirstName        string     `form:"firstName" json:"firstName"`
	LastName         string     `form:"lastName" json:"lastName"`
	Cellphone        string     `form:"cellphone" json:"cellphone"`
	ConfirmationCode int        `json:"-"`
	IsConfirmed      bool       `json:"isConfirmed"`
	JwtToken         string     `json:"jwtToken" gorm:"-"` // Ignore this field
	Birthday         time.Time  `json:"birthday"`
}

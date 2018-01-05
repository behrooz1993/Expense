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
	Password         string
	ConfirmationCode int    `json:"-"`
	IsConfirmed      bool   `json:"isConfirmed"`
	JwtToken         string `json:"jwtToken" gorm:"-"` // Ignore this field
}

func (db *DB) AddUser(user *User) error {
	db.database.Create(&user)
	if errs := db.database.GetErrors(); len(errs) > 0 {
		err := errs[0]
		return err
	}
	return nil
}

func (db *DB) SelectByCellphone(user *User) error {
	db.database.Where("cellphone Like ?", user.Cellphone).Find(&user)
	if errs := db.database.GetErrors(); len(errs) > 0 {
		err := errs[0]
		return err
	}
	return nil
}

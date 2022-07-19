package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	IDUser  int64  `json:"id_user" gorm:"primaryKey"`
	Balance int64  `json:"balance"`
	EMail   string `json:"e_mail"`
	Phone   string `json:"phone"`
	Role    []Role `json:"role" gorm:"many2one"`
}

type Role struct {
	Admin    bool `json:"admin"`
	Premium  bool `json:"premium"`
	Ordinary bool `json:"ordinary"`
}

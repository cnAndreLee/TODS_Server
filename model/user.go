package model

import (
	"time"
)

type User struct {
	Account    string
	Class      string
	Key        string
	Schoolname string
	Telephone  int
	Isadmin    bool
	Outdate    *time.Time
}

// type User struct {
// 	gorm.Model
// 	account    string     `gorm:"type:varchar(20);not null"`
// 	class      string     `gorm:"type:varchar(20)"`
// 	key        string     `gorm:"type:varchar(50);not null"`
// 	schoolname string     `gorm:"type:varchar(100)"`
// 	telephone  int        `gorm:"type:integer"`
// 	isadmin    bool       `gorm:"type:boolean;not null"`
// 	outdate    *time.Time `gorm:"type:date"`
// }

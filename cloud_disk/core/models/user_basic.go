package models

import "time"

type UserBasic struct {
	Id        int
	Name      string
	Identity  string
	Password  string
	Email     string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}

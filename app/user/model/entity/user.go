package entity

import (
	"github.com/weblogs/app/user/model/po"
)

type User struct {
	UserID    int64
	FirstName string
	LastName  string
	Email     string
	UserName  string
	Password  string
	UserRole  po.UserRole
	IsDel     bool
	po.Model
}

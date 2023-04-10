package po

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type User struct {
	Model
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID    int64              `json:"user_id" bson:"user_id"`
	UserName  string             `json:"user_name" bson:"user_name"`
	Password  string             `json:"password" bson:"password"`
	UserRole  UserRole           `json:"user_role" bson:"user_role"`
	Email     string             `json:"email" bson:"email"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	IsDel     bool               `json:"is_del" bson:"is_del"`
}

type UserRole int8

const (
	RoleGuest UserRole = 0
	RoleUser  UserRole = 1
	RoleAdmin UserRole = 2
	RoleRoot  UserRole = 3
)

func (u UserRole) String() string {
	m := map[UserRole]string{
		RoleRoot:  "root",
		RoleAdmin: "admin",
		RoleUser:  "user",
		RoleGuest: "guest",
	}
	return m[u]
}

func FromString(s string) (UserRole, error) {
	r := map[string]UserRole{
		"root":  RoleRoot,
		"admin": RoleAdmin,
		"user":  RoleUser,
		"guest": RoleGuest,
	}
	ur, ok := r[s]
	if !ok {
		return UserRole(-1), errors.New(fmt.Sprintf("have no such user role, role=%s", s))
	}
	return ur, nil
}

package database

import (
	"errors"
	"reflect"
)

type User struct {
	Id       int64  `gorm:"primaryKey;column:id;autoIncrement"`
	Name     string `gorm:"uniqueKey;not null;"`
	Password string `gorm:"not null;"`
	email    string `gorm:"not null"`
}

func GetUserById(id int64) (User, error) {
	var user User
	Db.Table("user").Limit(1).Find(&user, id)
	if reflect.DeepEqual(user, User{}) {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func GetUserByName(name string) (User, error) {
	var user User
	Db.Table("user").Limit(1).Find(&user, "name = ?", name)
	if reflect.DeepEqual(user, User{}) {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func GetUserByIdAndPwd(id int64, pwd string) (User, error) {
	var user User
	Db.Table("user").Limit(1).Find(&user, "id = ? AND password = ?", id, pwd)
	if reflect.DeepEqual(user, User{}) {
		return User{}, errors.New("user not found or password is wrong")
	}
	return user, nil
}

func GetUserByNameAndPwd(name string, pwd string) (User, error) {
	var user User
	Db.Table("user").Limit(1).Find(&user, "name = ? AND password = ?", name, pwd)
	if reflect.DeepEqual(user, User{}) {
		return User{}, errors.New("user not found or password is wrong")
	}
	return user, nil
}

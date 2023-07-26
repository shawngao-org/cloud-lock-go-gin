package database

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/util"
	"errors"
	"reflect"
	"strings"
)

type User struct {
	Id       int64  `gorm:"primaryKey;column:id;autoIncrement;not null;"`
	Name     string `gorm:"uniqueKey;column:name;not null;size:16;"`
	Password string `gorm:"not null;column:password;size:120;"`
	email    string `gorm:"uniqueKey;not null;column:email;size:30;"`
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

func GetUserByEmailAndPwd(email string, pwd string) (User, error) {
	var user User
	Db.Table("user").Limit(1).Find(&user, "email = ?", email)
	if reflect.DeepEqual(user, User{}) {
		return User{}, errors.New("invalid email address")
	}
	if strings.ToUpper(config.Conf.Security.Password.Method) == "BCRYPT" {
		e := util.CheckBcrypt(pwd, user.Password)
		if e == nil {
			user.Password = ""
			return user, nil
		} else {
			return User{}, errors.New("invalid password")
		}
	}
	str, err := util.PasswordEncrypt(pwd)
	if err != nil {
		return User{}, err
	}
	if str == user.Password {
		user.Password = ""
		return user, nil
	}
	return User{}, errors.New("invalid password")
}

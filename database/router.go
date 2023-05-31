package database

import (
	"errors"
	"reflect"
)

type Router struct {
	Id         int64  `gorm:"primaryKey;column:id;autoIncrement"`
	Path       string `gorm:"uniqueKey;not null;"`
	Method     string `gorm:"not null;"`
	Permission string `gorm:"not null;"`
}

func CheckRouterPermission(path string, uid int64) (bool, error) {
	var router Router
	Db.Table("router").Limit(1).Find(&router, "path = ?", path)
	if reflect.DeepEqual(router, Router{}) {
		return false, errors.New("not found")
	}
	if router.Permission == "none" {
		return true, nil
	}
	if ExistRouterPermission(uid, router.Permission) {
		return true, nil
	}
	return false, errors.New("forbidden")
}

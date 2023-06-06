package database

import (
	"errors"
	"reflect"
)

type Router struct {
	Id         int64  `gorm:"primaryKey;column:id;autoIncrement;not null;"`
	Path       string `gorm:"uniqueKey;column:path;not null;size:255;"`
	Method     string `gorm:"not null;column:method;size:10;"`
	Permission string `gorm:"not null;column:permission;size:64;"`
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

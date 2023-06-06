package database

import (
	"reflect"
)

type RouterPermission struct {
	Uid        int64  `gorm:"uniqueIndex:uid_permission;column:uid;not null;"`
	Permission string `gorm:"uniqueIndex:uid_permission;column:permission;not null;size:64;"`
}

func ExistRouterPermission(uid int64, p string) bool {
	var rp RouterPermission
	Db.Table("router_permission").Limit(1).Find(&rp, "uid = ? AND permission = ?", uid, p)
	if reflect.DeepEqual(rp, RouterPermission{}) {
		return false
	}
	return true
}

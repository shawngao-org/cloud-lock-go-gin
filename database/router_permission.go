package database

import (
	"reflect"
)

type RouterPermission struct {
	Uid        int64  `gorm:"uniqueIndex:uid_permission"`
	Permission string `gorm:"uniqueIndex:uid_permission"`
}

func ExistRouterPermission(uid int64, p string) bool {
	var rp RouterPermission
	Db.Table("router_permission").Limit(1).Find(&rp, "uid = ? AND permission = ?", uid, p)
	if reflect.DeepEqual(rp, RouterPermission{}) {
		return false
	}
	return true
}

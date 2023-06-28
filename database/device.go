package database

type Device struct {
	Id    string `gorm:"primaryKey;column:id;size:36;not null;comment:Device id;"`
	Name  string `gorm:"not null;column:name;default:未命名设备;size:32;comment:Device name;"`
	Addr  string `gorm:"not null;column:addr;size:100;comment:Device address;"`
	Model string `gorm:"not null;column:model;size:50;comment:Device model;"`
	Sn    string `gorm:"uniqueKey;not null;column:sn;size:36;comment:SN;"`
}

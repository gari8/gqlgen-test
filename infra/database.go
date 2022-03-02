package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Dialector{
		&mysql.Config{
			DriverName: "mysql",
			DSN:        "root:root@tcp(127.0.0.1:3306)/graph",
		},
	}, &gorm.Config{})
}

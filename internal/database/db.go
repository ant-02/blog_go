package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func Exists(db *gorm.DB, model interface{}, query interface{}, args ...interface{}) (bool, error) {
    var count int64
    err := db.Model(model).Where(query, args...).Count(&count).Error
    return count > 0, err
}

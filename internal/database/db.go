package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"zll.blog.com/internal/model"
)

func InitDatabase(dsn, nameDsn, dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	var count int64
	db.Raw("SELECT count(*) FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&count)

	if count == 0 {
		err = db.Exec("CREATE DATABASE " + dbName).Error
		if err != nil {
			return nil, err
		}
	}

	// 3. 关闭临时连接
	sqlDB, _ := db.DB()
	sqlDB.Close()

	// 4. 连接到目标数据库
	return gorm.Open(mysql.Open(nameDsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Article{}, &model.ArticleCategory{}, &model.ArticleTag{}, &model.Category{}, &model.Tag{}, &model.User{}, &model.UserFollow{})
}

func Exists(db *gorm.DB, model interface{}, query interface{}, args ...interface{}) (bool, error) {
	var count int64
	err := db.Model(model).Where(query, args...).Count(&count).Error
	return count > 0, err
}

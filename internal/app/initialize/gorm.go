package initialize

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/varluffy/gindemo/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// InitDB 初始化MySQL连接
func InitDB() (*gorm.DB, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		viper.GetString("MySQL.User"),
		viper.GetString("MySQL.Password"),
		viper.GetString("MySQL.Host"),
		viper.GetInt("MySQL.Port"),
		viper.GetString("MySQL.DBName"),
		viper.GetString("MySQL.Parameters"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, nil, err
	}
	if viper.GetBool("Gorm.Debug") {
		db.Debug()
	}
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("Gorm.MaxLifetime")) * time.Second)
	sqlDB.SetMaxIdleConns(viper.GetInt("Gorm.MaxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("Gorm.MaxOpenConns"))

	cleanFunc := func() {
		err := sqlDB.Close()
		if err != nil {
			logger.Errorf(context.Background(), "Gorm db close error: %s", err.Error())
		}
	}

	return db, cleanFunc, nil
}

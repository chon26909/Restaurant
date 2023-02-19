package lib

import (
	"context"
	"fmt"
	"restaurant/models"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func NewMySqlConnection() *gorm.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Menu{}, &models.Submenu{})
	if err != nil {
		fmt.Println("err auto migrate", err)
	}

	return db
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	stmt, _ := fc()
	fmt.Printf("\nSQL => %v\n", stmt)
}

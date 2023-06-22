package lib

import (
	"context"
	"fmt"
	"restaurant/config"
	"restaurant/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func NewMySqlConnection(config *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Db.Host,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Database,
		// viper.GetString("db.username"),
		// viper.GetString("db.password"),
		// viper.GetString("db.host"),
		// viper.GetString("db.port"),
		// viper.GetString("db.database"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})

	if err != nil {
		panic(err)
	}

	err = db.Migrator().AutoMigrate(
		&models.Menu{},
		&models.Submenu{},
		&models.Buffet{},
		&models.Table{},
		&models.Customer{},
		&models.Order{})

	if err != nil {
		fmt.Println("err auto migrate", err)
	}

	return db
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	stmt, _ := fc()
	fmt.Printf("\nSQL => %v\n", stmt)
}

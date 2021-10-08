package db

import (
	// "echo-demo-project/config"
	// "echo-demo-project/db/seeders"
	"fmt"
	"marketplacego/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)

	fmt.Println(dataSourceName)

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=evermos sslmode=disable password=secret")
	if err != nil {
		panic(err.Error())
	}

	return db
}

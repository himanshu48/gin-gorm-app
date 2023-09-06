package models

import (
	"fmt"
	"gin-gorm/configs"
	"gin-gorm/constants"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbConfig := configs.DbConfig()
	Dbdriver := dbConfig.DbType
	DbUser := dbConfig.Username
	DbPassword := dbConfig.Pswd
	DbPort := dbConfig.Port
	DbHost := dbConfig.Host
	DbName := dbConfig.DbName
	var err error

	if Dbdriver == constants.MySql {
		// DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		DBURL := fmt.Sprintf("%s:''@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbHost, DbPort, DbName)
		DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == constants.Postgres {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	} else if Dbdriver == constants.Sqlite {
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
	} else {
		log.Fatal("unknown db connection:", Dbdriver)
	}

	if configs.AppEnv() == constants.DevelopmentEnv {
		DB = DB.Debug()
	}

	DB.AutoMigrate(&User{}) //database migration
	// DB.AutoMigrate(&User{}, &Post{}) //database migration

}

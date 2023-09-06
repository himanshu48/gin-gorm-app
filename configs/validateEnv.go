package configs

import (
	"gin-gorm/constants"
	"log"
	"strconv"
)

func validateEnvVariables(config appConfig)  {
	// app port check
	appPort, err := strconv.Atoi(config.AppPort);
	if  err != nil || appPort < 1024 || appPort > 49151 {
		log.Fatalf("Invalid Port %s", config.AppPort)
	}

	// db config check
	dbPort, err := strconv.Atoi(config.DbConfig.Port);
	if config.DbConfig.DbType != constants.Sqlite && (err != nil || dbPort < 0 || dbPort > 65535) {
		log.Fatalf("Invalid Databse Port %s", config.DbConfig.Port)
	}

	// jwt config check
	if config.JwtConfig.ExpireIn < 1 || config.JwtConfig.ExpireIn > 24 {
		log.Fatalf("Invalid Jwt expiry time: %d", config.JwtConfig.ExpireIn)
	}
}
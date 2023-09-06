package configs

import (
	"gin-gorm/constants"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	DbType 		constants.DatabaseType
	Host 		string
    Port  		string
	Username 	string
	Pswd 		string
	DbName 		string
}

type jwtConfig struct {
	Key			string
	ExpireIn	int
}

type cookieConfig struct {
	Name		string
	MaxAge		int
	Path 		string
	Domain		string
	Secure		bool
	HttpOnly	bool
}

type appConfig struct {
	AppPort 	string
	AppEnv 		constants.AppEnvironment
	CookieConfig *cookieConfig
	DbConfig	*dbConfig
	JwtConfig	*jwtConfig
	LogLevel	constants.LogLevel
}

var config *appConfig

func getEnv(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        value = fallback
    }
    return value
}

func Init() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	jwtTime, err := strconv.Atoi(getEnv("JWT_EXPIRE_IN_HOUR", "24"));
	if err != nil {
		log.Fatal("Invalid Jwt expiry time.")
	}
	cookieMaxAge, err := strconv.Atoi(getEnv("COOKIE_MaxAge_Hour", "24"));
	if err != nil {
		log.Fatal("Invalid cookie maxage time.")
	}

	config = &appConfig{
		AppPort: getEnv("APP_PORT", "5000"),
		AppEnv: constants.ParseAppEnviroment(getEnv("APP_ENV", "production")),
		CookieConfig: &cookieConfig{
			Name: getEnv("COOKIE_NAME", "Authorization"),
			MaxAge: cookieMaxAge * 60 * 60,
			Path: "/",
			Domain: getEnv("COOKIE_DOMAIN", "localhost"),
			Secure: strings.ToLower(getEnv("COOKIE_SECURE", "true")) != "false",
			HttpOnly: true,
		},
		DbConfig: &dbConfig{
			DbType: constants.ParseDatabaseType(getEnv("DB_TYPE", "sqlite3")),
			Host: getEnv("DB_HOST", "localhost"),
			Port: getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", ""),
			Pswd: getEnv("DB_PASSWORD", ""),
			DbName: getEnv("DB_NAME", "test"),
		},
		JwtConfig: &jwtConfig{
			Key: getEnv("JWT_KEY", "F1vYd!uv0d#u24WP@qC69Qe6uI5LR$Yx"),
			ExpireIn: jwtTime,
		},
		LogLevel: constants.ParseLogLevel(getEnv("LOG_LEVEL", "ERROR")),
	}
	validateEnvVariables(*config)
}

func AppPort() string {
    return config.AppPort
}

func AppEnv() constants.AppEnvironment {
    return config.AppEnv
}

func CookieConfig() *cookieConfig {
	return config.CookieConfig
}

func DbConfig() *dbConfig {
	return config.DbConfig
}

func JwtConfig() *jwtConfig {
	return config.JwtConfig
}

func LogLevel() constants.LogLevel {
	return config.LogLevel
}

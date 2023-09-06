### Description
A sample go app using gin and gorm.


### initialize
go mod init gin-ws

go mod tidy <!-- like npm i -->

### install dependencies
go get -u github.com/gin-gonic/gin

go get gorm.io/gorm
go get gorm.io/driver/postgres
go get gorm.io/driver/sqlite
go get gorm.io/driver/mysql
go get golang.org/x/crypto/bcrypt
go get golang.org/x/time/rate


go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
go get github.com/go-playground/validator/v10
go get github.com/dgrijalva/jwt-go
go get github.com/sirupsen/logrus
go get github.com/gin-gonic/contrib/gzip

### run
go run main.go

### build
go build

### unistall unused package
go mod tidy

### run build
./app-name

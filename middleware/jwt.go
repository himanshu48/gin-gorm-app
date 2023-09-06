package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"gin-gorm/configs"
)

type authCustomClaims struct {
	UserName string `json:"userName"`
	// UserType string `json:"userType"`
	jwt.StandardClaims
}

func GenerateToken(userName string) string {
	claims := &authCustomClaims{
		userName,
		// userType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(configs.JwtConfig().ExpireIn)).Unix(),
			Issuer:    "test",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(configs.JwtConfig().Key))
	if err != nil {
		panic(err)
	}
	return t
}

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}

		// fmt.Printf("%+v", token.Claims.(*authCustomClaims).StandardClaims.ExpiresAt)
		if int64(token.Claims.(*authCustomClaims).StandardClaims.ExpiresAt) < time.Now().Unix() {
			return nil, fmt.Errorf("Expired token")
		}

		return []byte(configs.JwtConfig().Key), nil
	})
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// currentURL := c.Request.URL.String()
		// token := c.GetHeader("Authorization")
		token, err := c.Cookie("Authorization")
		if err != nil || token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token not found"})
			return
		}
		arr := strings.Split(token, " ")
		if len(arr) != 2 || arr[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
		token = arr[1]

		t, err := validateToken(token)
		if claims, ok := t.Claims.(*authCustomClaims); ok && t.Valid {
			c.Request.Header.Set("userName", claims.UserName)
			// fmt.Printf("%v %v", claims.UserName, claims.StandardClaims.ExpiresAt)
			c.Next()
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
	}
}
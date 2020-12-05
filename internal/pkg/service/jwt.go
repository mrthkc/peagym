package service

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	log "github.com/sirupsen/logrus"
)

// SECRET : 
const SECRET = "peagym-20201204"

// GenerateToken :
func GenerateToken(userID int) (string, error){
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(Config.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateToken :
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")
		}
		return []byte(Config.Secret), nil
	})
}

// JWTAuth :
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bearer = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(bearer):]
		token, err := ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Debug(claims)
		} else {
			log.Error(err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
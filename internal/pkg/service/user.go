package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/mrthkc/peagym/internal/pkg/adapter"
	"github.com/mrthkc/peagym/internal/pkg/entity"

	log "github.com/sirupsen/logrus"
)

// Token :
func Token(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": 1,
		"data": "PeaGym",
	})
}

// Register :
func Register(c *gin.Context) {
	entity.DB = adapter.NewClient(Config.DBCred.Host, Config.DBCred.Port, Config.DBCred.DB, Config.DBCred.User, Config.DBCred.Password)
	defer entity.DB.Close()

	user := new(entity.User)
	if err := c.BindJSON(user); err != nil {
		log.Error("Binding user error occured: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
		})
		return
	}

	user.Password = encryptPass([]byte(user.Password))
	data, err := entity.AddUser(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": 1,
		"data":    data,
	})
}

// Login :
func Login(c *gin.Context) {
	entity.DB = adapter.NewClient(Config.DBCred.Host, Config.DBCred.Port, Config.DBCred.DB, Config.DBCred.User, Config.DBCred.Password)
	defer entity.DB.Close()

	user := new(entity.User)
	if err := c.BindJSON(user); err != nil {
		log.Error("Binding user error occured: ", err)
		c.JSON(http.StatusBadRequest, "Binding error")
		return
	}

	checkUser := entity.GetUserByEmail(user.Email)
	if checkUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
		})
		return
	}

	login := comparePasswords(checkUser.Password, []byte(user.Password))
	if login == false {
		c.JSON(http.StatusForbidden, gin.H{
			"success": 0,
		})
		return
	}

	token, err := GenerateToken(checkUser.ID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"success": 0,
		})
		return
	}
	c.JSON(http.StatusForbidden, gin.H{
		"success": 1,
		"data":    gin.H{
			"uid": checkUser.ID,
			"token": token,
		},
	})
	return
}

func encryptPass(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Error("campare pass: ", err)
		return false
	}
	return true
}

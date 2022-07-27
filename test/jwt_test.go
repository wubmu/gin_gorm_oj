package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var myKey = []byte("gin-gorm-oj-key")

//生成token
func TestGenerateToken(t *testing.T) {
	//
	UserClaims := &UserClaims{
		Identity:       "user_1",
		Name:           "Get",
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tokenString)
}

// 解析 token
func TestAnalyseToken(t *testing.T) {
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXJfMSIsIm5hbWUiOiJHZXQifQ.4inO9HZINmKFYO9qEF2SYYPHk0GuuA-qUdwIhUa8USE
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXJfMSIsIm5hbWUiOiJHZXQifQ.4inO9HZINmKFYO9qEF2SYYPHk0GuuA-qUdwIhUa8USE"
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claims.Valid {
		fmt.Println(userClaim)
	}
}

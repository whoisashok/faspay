package generatepwd

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type JwtClaims struct {
	// add as necessary
	jwt.StandardClaims
}

// CreateJwtToken ...
func CreateJwtToken(secret string, apiToken string) (string, error) {
	claims := JwtClaims{
		jwt.StandardClaims{
			Id:        apiToken,
			ExpiresAt: time.Now().Add(2160 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, err
}

func HashAndSalt(plainPwd string) string {

	bytePlainPwd := []byte(plainPwd)
	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(bytePlainPwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println("Passwords to hashAndSalt err: ", err)
		return ""
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPwd string) bool {
	bytePlainPwd := []byte(plainPwd)
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlainPwd)
	if err != nil {
		fmt.Println("ComparePassword err: ", err)
		return false
	}

	return true
}

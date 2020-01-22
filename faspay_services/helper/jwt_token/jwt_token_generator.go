package jwttoken

import (
	"strings"

	cfg "payment_http/config"
	servicePayment "payment_http/payment_services"
	vmPayment "payment_http/payment_services/view_model"

	jwt "github.com/dgrijalva/jwt-go"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
}

type JwtClaims struct {
	CustomClaim vmPayment.ExtUserCustomClaimResponse
	jwt.StandardClaims
}

func CreateJWT(customClaim vmPayment.ExtUserCustomClaimResponse) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	claims := JwtClaims{
		customClaim,
		jwt.StandardClaims{
			Id: customClaim.UserCode,
			// If we want to verify expiry uncomment below ExpiresAt
			//ExpiresAt: time.Now().Add(2160 * time.Hour).Unix(),
			//ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	hmacSecretKey := []byte(config.GetString(`app.secret`))
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecretKey)
	if err != nil {
		//fmt.Println("token.SignedString - Unexpected Error", err)
		return "", err
	}
	//fmt.Println("User Claim tokenString: ", tokenString)
	return tokenString, err
}

func JWTParse(authTokenHeader string) (*JwtClaims, error) {
	//Get token from header string
	//tokenString := strings.Replace(authTokenHeader, "Bearer ", "", -1)
	tokenString := ""
	authHeaderParts := strings.Split(authTokenHeader, " ")
	if len(authHeaderParts) == 2 && strings.ToLower(authHeaderParts[0]) == "bearer" {
		tokenString = authHeaderParts[1]
	}
	if tokenString == "" {
		return nil, servicePayment.ErrUnauthorizedError
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString(`app.secret`)), nil
	})

	if err != nil {
		//fmt.Println("User Claim - Unexpected Error", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		//fmt.Printf("User Claim UserCode %v ExpiresAt %v", claims.CustomClaim.UserCode, claims.StandardClaims.ExpiresAt)
		return claims, nil
	}

	return nil, servicePayment.ErrUnauthorizedError
}

package jwt

import (
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	jwtInfo Info
)

// TokenData information
type TokenData struct {
	ID    string `bson:"id" json:"id" model:"ID"`
	Email string `bson:"email" json:"email" require:"true" model:"Email"`
}

// Info is the details for the Jwt
type Info struct {
	Key       string
	ExpiresAt string
	Issuer    string
}

// Claims - jwt claims information
type Claims struct {
	Data string `json:"data"`
	jwt.StandardClaims
}

// ConfigInfo - config info data
func ConfigInfo(j Info) {
	jwtInfo = j
}

func readConfig() Info {
	return jwtInfo
}

// GetToken - get JWT token using email and userID
func GetToken(email string, ID string) string {
	finalData := TokenData{
		ID,
		email,
	}
	data, _ := json.Marshal(finalData)
	j := readConfig()
	expireToken := time.Now().Add(time.Hour * 72).Unix()
	claims := Claims{
		string(data),
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    j.Issuer,
		},
	}
	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := token.SignedString([]byte(j.Key))
	return signedToken
}

// GetData - get user data using JWT token
func GetData(userToken string) (TokenData, error) {
	j := readConfig()
	var userData TokenData
	token, err := jwt.ParseWithClaims(userToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Key), nil
	})
	if err != nil {
		return userData, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		err := json.Unmarshal([]byte(claims.Data), &userData)
		return userData, err
	}
	return userData, fmt.Errorf("Unexpected claims value")
}

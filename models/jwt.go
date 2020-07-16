package models
import (
	"github.com/dgrijalva/jwt-go"
	"time"
)
var secretKey = "SuperSecretKey"

type JWTPackage struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (j *JWTPackage)Generate(account Account)error{
	var err error
	access := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username":	account.Username,
		"password":	account.Password,
		"exp":		time.Now().Add(time.Minute * 15).Unix(),
	})
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username":	account.Username,
		"password":	account.Password,
		"exp":		time.Now().Add(time.Hour * 24).Unix(),
	})

	j.AccessToken, err = access.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}

	j.RefreshToken, err = refresh.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}

	return nil
}
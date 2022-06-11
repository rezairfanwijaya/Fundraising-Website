package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go/v4"
)

// bikin service
type Service interface {
	GenerateToken(userId int) (string, error)
	ValidasiToken(token string) (*jwt.Token, error)
}

// bikin struct
type jwtToken struct{}

// secret KEY
var KEY string = "fndjfndjbndknbkdjbk"

// bikin newserive agar semua function bisa diakses dari package manapun
func NewServiceAuth() *jwtToken {
	return &jwtToken{}
}

// funtion generate token
func (s *jwtToken) GenerateToken(userId int) (string, error) {
	// bikin payload (claim) yang akan diisi dengan user id
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// tanda tangani token
	signedToken, err := token.SignedString([]byte(KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// function validasi token
func (s *jwtToken) ValidasiToken(token string) (*jwt.Token, error) {
	// parsing token untuk di cek apakah valid atau tidak
	myToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// cek method token
		// hmac adalah hs256
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		// mengembalikan secret key
		return []byte(KEY), nil
	})

	if err != nil {
		return myToken, err
	}

	return myToken, nil
}

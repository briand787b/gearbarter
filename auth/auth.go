package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "auth/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "auth/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

// remove this function if code ever goes
// to production.
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateJWT generates a json web token (JWT)
func GenerateJWT() (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = jwt.StandardClaims{
		Subject:   "bar",
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
	}

	tokenString, err := t.SignedString(signKey)
	// fmt.Println(tokenString) // dev only
	return tokenString, err
}

// ValidateToken returns a bool that indicates if
// the token from the passed request is valid
func validateToken(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	fmt.Println("auth header: ", authHeader)

	if authHeader != "" {
		return true
	}

	return false
}

func Authmiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("in auth middleware")
	if err := r.ParseForm(); err != nil {
		http.Error(w, "error parsing form", 500)
	}

	if r.Form.Get("password") != "123" {
		w.WriteHeader(401)
		return
	}

	next(w, r)
}

// GetJWTMiddleware returns a JWTMiddleware
// MAKE THIS MORE DESCRIPTIVE
func GetJWTMiddleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})
}

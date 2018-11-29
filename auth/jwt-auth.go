package auth

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/salambayev/x-boat-project/server"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//  Generate RSA signing files via shell (adjust as needed):
//
//  $ openssl genrsa -out app.rsa 1024
//  $ openssl rsa -in app.rsa -pubout > app.rsa.pub
//

const (
	// For simplicity these files are in the same folder as the app binary.
	// You shouldn't do this in production.
	privKeyPath = "app.rsa"
	pubKeyPath  = "app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

type UserCredentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userCredentials UserCredentials
	err := json.NewDecoder(r.Body).Decode(&userCredentials)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Error in request: {}", w)
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(12)).Unix()
	claims["iat"] = time.Now().Unix()

	token.Claims = claims
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error while signing the token: {}", err.Error())
		return
	}
	response := Token{tokenString}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(jsonResponse)
}

func ValidateToken(w http.ResponseWriter, r *http.Request, handler server.HttpHandler) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	if err == nil {
		if token.Valid {
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("Token is not valid: {}", err.Error())
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("Unauthorized access to this resource: {}", err.Error())
	}
}
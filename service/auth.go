package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BigBrother3/server/database/database"
	jwt "github.com/dgrijalva/jwt-go"
)

const secret = "bigbrother"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

//Log in and register

func registerHandler(w http.ResponseWriter, req *http.Request) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error in register request format")
		w.Write([]byte("register request format is wrong."))
		return
	}

	if database.CheckKeyExist([]byte("users"), []byte(user.Username)) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error in register info")
		w.Write([]byte("Existent username."))
		return
	}

	database.Update([]byte("users"), []byte(user.Username), []byte(user.Password))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create a account"))
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error in login request format")
		w.Write([]byte("login request format is wrong."))
		return
	}

	if !database.CheckKeyExist([]byte("users"), []byte(user.Username)) || user.Password != database.GetValue([]byte("users"), []byte(user.Username)) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error in login info")
		w.Write([]byte("Inexistent user or wrong password."))
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in signing the token")
		log.Fatal(err)
	}

	response := Token{tokenStr}
	JsonResponse(response, w)
}

func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

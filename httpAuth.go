package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	mysql "./mysql"

	gorillactx "github.com/gorilla/context"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("Secret_Key_Shop")

type Claims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Welcome")
	// 	credets := &mysql.User{}
	// 	err := json.NewDecoder(r.Body).Decode(&credets)
	// 	log.Printf("%+v", credets)
	// 	if err != nil {
	// 		http.Error(w, "Wrong JSON", http.StatusBadRequest)
	// 		return
	// 	}
	// 	json, _ := json.Marshal(map[string]string{"token": "userToken123"})
	// 	fmt.Fprintf(w, string(json))
}

func AuthCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt-token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenString := c.Value

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		gorillactx.Set(r, "user_id", claims.UserID)
		next(w, r)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	log.Println("SignIn")
	credets := &mysql.User{}
	err := json.NewDecoder(r.Body).Decode(&credets)
	log.Printf("%+v", credets)
	if err != nil {
		http.Error(w, "Wrong JSON", http.StatusBadRequest)
		return
	}
	//get password hash by login in DB
	userID, password, err := db.GetUserByLogin([]interface{}{credets.Login})
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	//check if the password is not empty
	if password == "" {
		http.Error(w, "Wrong Login", http.StatusUnauthorized)
		return
	}
	//check recieved password with has from DB
	if !CheckPasswordHash(credets.Password, password) {
		http.Error(w, "Wrong Password", http.StatusUnauthorized)
		return
	}
	//if the password is valid then generate JWT-Token
	token, expTime, err := GenerateJWTToken(userID)
	if err != nil {
		http.Error(w, "Token error", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, "Can't generate auth token", http.StatusInternalServerError)
		return
	}
	//Set JWT token in Cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt-token",
		Value:   token,
		Expires: expTime,
	})

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWTToken(user_id string) (string, time.Time, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		UserID: user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	tokenSigned, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}
	return tokenSigned, expirationTime, nil
}

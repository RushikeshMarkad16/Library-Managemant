package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/RushikeshMarkad16/Library-Managemant/api"
	"github.com/RushikeshMarkad16/Library-Managemant/book"
	"github.com/RushikeshMarkad16/Library-Managemant/transaction"
	"github.com/RushikeshMarkad16/Library-Managemant/user"
	"github.com/golang-jwt/jwt"

	"github.com/gorilla/mux"
	// "github.com/go-gorm/gorm"
	// "github.com/golang-jwt/jwt"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	Role        string `json:"role"`
	TokenString string `json:"token"`
}

const (
	secretkey = "jsd549$^&"
)

func GenerateJWT(email, role, id string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			err := errors.New("no Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			err := errors.New("your Token has been expired")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {

				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			} else if claims["role"] == "superadmin" {

				r.Header.Set("Role", "superadmin")
				handler.ServeHTTP(w, r)
				return
			}
		}
		reserr := errors.New("not Authorized")
		json.NewEncoder(w).Encode(reserr)
	}
}

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//User
	router.HandleFunc("/login", user.Login()).Methods(http.MethodPost)
	router.HandleFunc("/users", user.Create(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", user.List(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.FindByID(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.DeleteByID(dep.UserService)).Methods(http.MethodDelete)
	router.HandleFunc("/users", user.Update(dep.UserService)).Methods(http.MethodPut)

	//Book
	router.HandleFunc("/books", book.Create(dep.BookService)).Methods(http.MethodPost)
	router.HandleFunc("/books", book.List(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", book.FindByID(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", book.DeleteByID(dep.BookService)).Methods(http.MethodDelete)
	router.HandleFunc("/books", book.Update(dep.BookService)).Methods(http.MethodPut)

	//Transaction
	router.HandleFunc("/book/issue", transaction.Create(dep.TransactionService)).Methods(http.MethodPost)
	router.HandleFunc("/book/return", transaction.Update(dep.TransactionService)).Methods(http.MethodPut)
	router.HandleFunc("/userbook/transaction", transaction.List(dep.TransactionService)).Methods(http.MethodGet)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}

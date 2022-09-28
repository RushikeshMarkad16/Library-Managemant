package server

import (
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
)

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
type TokenData struct {
	Id    string
	Email string
	Role  string
}

const (
	SUPERADMIN = iota
	ADMIN
	USER
)

var RoleMap = map[string]int{"superadmin": SUPERADMIN, "admin": ADMIN, "user": USER}

var secretkey = []byte("jsd549$^&")

func Authorize(handler http.HandlerFunc, role int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		isValid, tokenData, err := ValidateToken(token)
		fmt.Println(isValid)
		if err != nil {
			fmt.Println("error")
		}

		fmt.Println("Token Data : ", tokenData)

		if !isValid {
			//Send error response to api
			api.Error(w, http.StatusBadRequest, api.Response{Message: "Token is not valid"})
			return
		}

		tokenRole := tokenData.Role
		if RoleMap[tokenRole] > role {
			api.Error(w, http.StatusBadRequest, api.Response{Message: "You don't have the access"})
			return
		}

		handler.ServeHTTP(w, r)
		return
	}
}

func ValidateToken(tokenString string) (isValid bool, tokenData TokenData, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	isValid = true

	tokenData = TokenData{
		Id:    claims.Id,
		Email: claims.Email,
		Role:  claims.Role,
	}
	return
}

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//User
	router.HandleFunc("/login", user.Login(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", Authorize(user.Create(dep.UserService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/users", Authorize(user.List(dep.UserService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", Authorize(user.FindByID(dep.UserService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/users/{filterData}", Authorize(user.FilterByData(dep.UserService), ADMIN)).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", Authorize(user.DeleteByID(dep.UserService), ADMIN)).Methods(http.MethodDelete)
	router.HandleFunc("/users", Authorize(user.Update(dep.UserService), USER)).Methods(http.MethodPut)
	router.HandleFunc("/user/password/reset", Authorize(user.UpdatePassword(dep.UserService), USER)).Methods(http.MethodPut)

	//Book
	router.HandleFunc("/books", Authorize(book.Create(dep.BookService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/books", Authorize(book.List(dep.BookService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", Authorize(book.FindByID(dep.BookService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", Authorize(book.DeleteByID(dep.BookService), ADMIN)).Methods(http.MethodDelete)
	router.HandleFunc("/books", Authorize(book.Update(dep.BookService), ADMIN)).Methods(http.MethodPut)

	//Transaction
	router.HandleFunc("/book/issue", Authorize(transaction.Create(dep.TransactionService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/book/return", Authorize(transaction.Update(dep.TransactionService), ADMIN)).Methods(http.MethodPut)
	router.HandleFunc("/userbook/transaction", Authorize(transaction.List(dep.TransactionService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/bookstatus", Authorize(transaction.GetBookStatus(dep.TransactionService), USER)).Methods(http.MethodGet)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}

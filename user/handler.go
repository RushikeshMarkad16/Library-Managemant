package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RushikeshMarkad16/Library-Managemant/api"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type JWTClaim struct {
	UserID string `json:user_id`
	Email  string `json:email`
	Role   string `json:role`
	jwt.StandardClaims
}

var jwtKey = []byte("bd46*@^g")

func Login() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var j JWTClaim
		err := json.NewDecoder(req.Body).Decode(&j)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		jwtString, err := GenerateJWT(j.UserID, j.Email, j.Role)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, jwtString)

	})
}

func GenerateJWT(UserID string, Email string, Role string) (tokenString string, err error) {

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:  Email,
		UserID: UserID,
		Role:   Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func Authorize(handler http.HandlerFunc, role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//TODO:
		//1. get token from reuqest header
		//2. decode token
		//3. check if user exist from token.ID
		//4. if role is allowed
		//5. call handler

		tokenString := r.Header.Get("Authorization")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		for key, val := range claims {
			fmt.Printf("Key: %v, value: %v\n", key, val)
			fmt.Println(key, val)
		}

		fmt.Println(token, err)

	}
}

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var createUserRequest User
		err := json.NewDecoder(req.Body).Decode(&createUserRequest)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.create(req.Context(), createUserRequest)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
	})
}

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resp, err := service.list(req.Context())
		if err == errNoUsers {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func FindByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		resp, err := service.findByID(req.Context(), vars["id"])

		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func DeleteByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		err := service.deleteByID(req.Context(), vars["id"])
		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Deleted Successfully"})
	})
}

func Update(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c User
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.update(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
	})
}

func isBadRequest(err error) bool {
	return err == errEmptyFirstName || err == errEmptyID
}

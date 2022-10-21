package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RushikeshMarkad16/Library-Managemant/api"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var j Authentication
		err := json.NewDecoder(req.Body).Decode(&j)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		jwtString, err1 := service.GenerateJWT(req.Context(), j.Email, j.Password)
		if err1 != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err1.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, jwtString)

	})
}

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var createUserRequest User
		err := json.NewDecoder(req.Body).Decode(&createUserRequest)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.Create(req.Context(), createUserRequest)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Added user Successfully"})
	})
}

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resp, err := service.List(req.Context())
		if err == errNoUsers {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		var temp []UserToDisplay
		for _, j := range resp.Users {
			var temp1 UserToDisplay
			temp1.ID = j.ID
			temp1.FirstName = j.FirstName
			temp1.Last_name = j.Last_name
			temp1.Gender = j.Gender
			temp1.Address = j.Address
			temp1.Email = j.Email
			temp1.Mob_no = j.Mob_no
			temp1.Role = j.Role
			temp = append(temp, temp1)
		}

		api.Success(rw, http.StatusOK, temp)
	})
}

func FindByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		resp, err := service.FindByID(req.Context(), vars["id"])

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

func FilterByData(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		resp, err := service.FindByID(req.Context(), vars["id"])

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

		err := service.DeleteByID(req.Context(), vars["id"])
		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "User Deleted Successfully"})
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

		err = service.Update(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "User details Updated Successfully"})
	})
}

var v User

//var flag = 0

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}

func UpdatePassword(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c ChangePassword
		resp, err := service.List(req.Context())
		var flag = 0
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		err = json.NewDecoder(req.Body).Decode(&c)

		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		for _, v = range resp.Users {
			if v.ID == c.ID {
				if CheckPasswordHash(c.Password, v.Password) {
					flag = 1
					err = service.UpdatePassword(req.Context(), c)
					if isBadRequest(err) {
						api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
						return
					}

					if err != nil {
						api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
						return
					}

					api.Success(rw, http.StatusOK, api.Response{Message: "Password Updated Successfully"})
					return
				}
			} else {
				flag = 0
			}
		}
		if flag == 0 {
			api.Success(rw, http.StatusOK, api.Response{Message: "Invalid ID or Password"})
		}

	})
}

func isBadRequest(err error) bool {
	return err == errInvalidFirstName || err == errEmptyID || err == errInvalidLastName
}

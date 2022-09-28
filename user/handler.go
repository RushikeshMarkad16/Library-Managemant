package user

import (
	"encoding/json"
	"net/http"

	"github.com/RushikeshMarkad16/Library-Managemant/api"
	"github.com/gorilla/mux"
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

		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
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

		api.Success(rw, http.StatusOK, resp)
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

		err = service.Update(req.Context(), c)
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

var v User
var flag = 0

func UpdatePassword(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c ChangePassword
		resp, err := service.List(req.Context())

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
			if v.ID == c.ID && v.Password == c.Password {
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

				api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
			}
		}
		if flag != 1 {
			api.Success(rw, http.StatusOK, api.Response{Message: "Wrong ID or Password"})
		}

	})
}

func isBadRequest(err error) bool {
	return err == errInvalidFirstName || err == errEmptyID || err == errInvalidLastName
}

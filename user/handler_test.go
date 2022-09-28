package user_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RushikeshMarkad16/Library-Managemant/user"
	usermock "github.com/RushikeshMarkad16/Library-Managemant/user/mocks"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	fmt.Println("Expected Code :", expected, "Actual Code : ", actual)
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}

}

func makeHTTPCall(handler http.HandlerFunc, method, path, body string) (rr *httptest.ResponseRecorder) {
	request := []byte(body)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(request))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return
}

// Create:
func TestSuccessfullCreate(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Success"))

	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
		"first_name": "Yash",
		"last_name": "Shinde",
		"gender": "Male",
		"address": "Chinchwad",
		"email": "yash.shinde@gmail.com",
		"password": "^HfO85)#",
		"mob_no": "9763528946",
		"role":"user"
	}`)
	fmt.Println("Account Success")
	checkResponseCode(t, http.StatusCreated, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInvalidRequestBody(t *testing.T) {
	cs := &usermock.Service{}
	// cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("HEllo"))
	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
		"first_name": "Yash",
		"last_name": "Shinde",
		"gender": "Male",
		"address": "Chinchwad",
		"email": "yash.shinde@gmail.com",
		"password": "^HfO85)#",
		"mob_no": "9763528946",
		"role":"user"
	}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenEmptyName(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Empty Name"))

	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
        "id":"67",
        "first_name": "Rushikesh",
        "last_name": "",
        "gender": "Male",
        "age": 22,
        "address": "Pune",
        "email": "markaad@gmail.com",
        "password": "ertikol@123",
        "mob_no": "9623614171",
        "role":"user"
    }`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

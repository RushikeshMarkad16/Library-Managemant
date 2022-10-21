package book_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RushikeshMarkad16/Library-Managemant/book"
	bookmock "github.com/RushikeshMarkad16/Library-Managemant/book/mocks"
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
	cs := &bookmock.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Success "))

	rr := makeHTTPCall(book.Create(cs), http.MethodPost, "/books", `{
		"name": "Golang",
		"author":"Harry Inglis",
		"price":180,
		"total_copies":85,
		"status": "available",
		"available_copies":85
    }`)
	fmt.Println("Book Success")
	checkResponseCode(t, http.StatusCreated, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInvalidRequestBody(t *testing.T) {
	cs := &bookmock.Service{}
	// cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("HEllo"))
	rr := makeHTTPCall(book.Create(cs), http.MethodPost, "/books", `{
		"name": "Eloquent JavaScript,
		"author": "Benn Goues",
		"price":320,
		"totalcopies":100,
		"status": "Available",
		"availablecopies":100
    }`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenEmptyName(t *testing.T) {
	cs := &bookmock.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Empty Name"))

	rr := makeHTTPCall(book.Create(cs), http.MethodPost, "/books", `{
		"name": "",
		"author": "",
		"price":320,
		"totalcopies":100,
		"status": "Available",
		"availablecopies":100
    }`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

// List :
func TestSuccessfullList(t *testing.T) {
	cs := &bookmock.Service{}
	var resp book.ListResponse
	cs.On("List", mock.Anything).Return(resp, nil)

	rr := makeHTTPCall(book.List(cs), http.MethodGet, "/books", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

// func TestListWhenNoBooks(t *testing.T) {
// 	cs := &bookmock.Service{}
// 	var resp book.ListResponse
// 	cs.On("List", mock.Anything).Return(resp, nil)

// 	rr := makeHTTPCall(book.List(cs), http.MethodGet, "/books", "")

// 	checkResponseCode(t, http.StatusNotFound, rr.Code)
// 	cs.AssertExpectations(t)
// }

func TestListInternalError(t *testing.T) {
	cs := &bookmock.Service{}
	var resp book.ListResponse
	cs.On("List", mock.Anything).Return(resp, errors.New("Internal Error"))

	rr := makeHTTPCall(book.List(cs), http.MethodGet, "/books", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

//FindById
//not bad reqe
//not find err
func TestSuccessfullFindByID(t *testing.T) {
	cs := &bookmock.Service{}
	var lr book.FindByIDResponse
	cs.On("FindByID", mock.Anything, mock.Anything).Return(lr, nil)

	rr := makeHTTPCall(book.FindByID(cs), http.MethodGet, "/books/39d3e1ce-d66a-49fa-8455-4966357a5e66", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

//DeleteByID
func TestSuccessfullDeleteByID(t *testing.T) {
	cs := &bookmock.Service{}

	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(book.DeleteByID(cs), http.MethodDelete, "/books/39d3e1ce-d66a-49fa-8455-4966357a5e66", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestDeleteByIDWhenInternalError(t *testing.T) {
	cs := &bookmock.Service{}
	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(book.DeleteByID(cs), http.MethodDelete, "/books/39d3e1ce-d66a-49fa-8455-4966357a5e66", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

func TestSuccessfullUpdate(t *testing.T) {
	cs := &bookmock.Service{}
	cs.On("Update", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(book.Update(cs), http.MethodPut, "/books", `{"id":"39d3e1ce-d66a-49fa-8455-4966357a5e66", "Name":"Eloquent JavaScript"}`)

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenInvalidRequestBody(t *testing.T) {
	cs := &bookmock.Service{}

	rr := makeHTTPCall(book.Update(cs), http.MethodPut, "/books", `{"id":"0c3f6455-e904-44b8-88e7-af191828ee6", "Name":"Empty Book",}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenInternalError(t *testing.T) {
	cs := &bookmock.Service{}
	cs.On("Update", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(book.Update(cs), http.MethodPut, "/books", `{"id":"0c3f6455-e904-44b8-88e7-af191828ee64"}`)

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

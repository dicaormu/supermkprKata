package super

import (
	"testing"
	"os"
	"bytes"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestBuyProductWithError(t *testing.T) {
	payload := []byte(`{"name":"totoprod","price":11.22}`)

	req, err := http.NewRequest("POST", "/stock/totoprod", bytes.NewBuffer(payload))
	response := executeRequest(req)
	assert.Equal(t, err, nil)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
	assert.Contains(t, response.Body.String(), "Error finding a product")
}

func TestBuyProduct(t *testing.T) {
	payload := []byte(`{"Amt":10,"Prod":{"Price":0,"Name":"","Discount":0}}`)
	a.Stock = initSupermarket()
	req, err := http.NewRequest("POST", "/stock/eau", bytes.NewBuffer(payload))
	response := executeRequest(req)
	assert.Equal(t, err, nil)
	checkResponseCode(t, http.StatusCreated, response.Code)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

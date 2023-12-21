package tests_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GrandOichii/messager-app/back/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func GetRouter() *gin.Engine {
	return gin.Default()
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func toData(o interface{}) io.Reader {
	j, _ := json.Marshal(o)
	return bytes.NewBuffer(j)
}

func req(r *gin.Engine, t *testing.T, request string, path string, data interface{}) (*httptest.ResponseRecorder, []byte) {
	var reqData io.Reader = nil
	if data != nil {
		reqData = toData(data)
	}
	req, err := http.NewRequest(request, path, reqData)
	checkErr(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	result, err := io.ReadAll(w.Body)
	checkErr(t, err)
	return w, result
}

func Test_GetUsers(t *testing.T) {
	r := router.CreateRouter()

	w, _ := req(r, t, "GET", "/api/users", nil)

	assert.Equal(t, w.Code, http.StatusOK)
}

func Test_Register(t *testing.T) {
	// r := router.CreateRouter()

}

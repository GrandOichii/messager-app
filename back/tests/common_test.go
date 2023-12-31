package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/router"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gin-gonic/gin"
)

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func toData(o interface{}) io.Reader {
	j, _ := json.Marshal(o)
	return bytes.NewBuffer(j)
}

func req(r *gin.Engine, t *testing.T, request string, path string, data interface{}, token string) (*httptest.ResponseRecorder, []byte) {
	var reqData io.Reader = nil
	if data != nil {
		reqData = toData(data)
	}
	req, err := http.NewRequest(request, path, reqData)
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	checkErr(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	result, err := io.ReadAll(w.Body)
	checkErr(t, err)
	return w, result
}

func createUser(r *gin.Engine, t *testing.T, handle string, email string, password string) {
	req(r, t, "POST", "/api/users/register", models.CreateUser{
		Email:    email,
		Password: password,
		Handle:   handle,
	}, "")
}

func loginAs(r *gin.Engine, t *testing.T, handle string, email string, password string) string {
	createUser(r, t, handle, email, password)

	_, data := req(r, t, "POST", "/api/users/login", models.LoginUser{
		Email:    email,
		Password: password,
	}, "")

	var res struct {
		Token string `json:"token"`
	}
	json.Unmarshal(data, &res)

	return res.Token
}

func createRouter() *router.Router {
	result := router.CreateRouter()

	result.Services.UserServicer = services.NewUserService()
	result.Services.ChatServicer = services.NewChatService(result.Services)

	return result
}

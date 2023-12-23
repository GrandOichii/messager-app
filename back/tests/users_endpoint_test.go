package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/router"
	"github.com/stretchr/testify/assert"
)

// func Test_GetUsers(t *testing.T) {
// 	r := router.CreateRouter()

// 	w, _ := req(r, t, "GET", "/api/users", nil)

// 	assert.Equal(t, http.StatusOK, w.Code)
// }

func Test_Register(t *testing.T) {
	r := router.CreateRouter()

	handle := "coolhandle"
	w, data := req(r, t, "POST", "/api/users/register", models.CreateUser{
		Email:    "mymail@mail.com",
		Password: "pass",
		Handle:   handle,
	})

	var newUser *models.GetUser
	err := json.Unmarshal(data, &newUser)
	checkErr(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, handle, newUser.Handle)
}

func Test_RegisterFail(t *testing.T) {

	// TODO add more
	testCases := []struct {
		desc     string
		email    string
		password string
		handle   string
	}{
		{
			desc:     "Empty email",
			email:    "",
			password: "1234",
			handle:   "handle",
		},
		{
			desc:     "Empty password",
			password: "",
			email:    "mymail@mail.com",
			handle:   "handle",
		},
		{
			desc:     "Empty handle",
			handle:   "",
			email:    "mymail@mail.com",
			password: "1234",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := router.CreateRouter()
			w, _ := req(r, t, "POST", "/api/users/register", tC)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}

func Test_Login(t *testing.T) {
	r := router.CreateRouter()
	req(r, t, "POST", "/api/users/register", models.CreateUser{
		Email:    "mymail@mail.com",
		Password: "1234",
		Handle:   "coolhandle",
	})

	w, _ := req(r, t, "POST", "/api/users/login", models.PostUser{
		Email:    "mymail@mail.com",
		Password: "1234",
	})

	assert.Equal(t, http.StatusOK, w.Code)

}

func Test_LoginFailed(t *testing.T) {
	r := router.CreateRouter()
	req(r, t, "POST", "/api/users/register", models.CreateUser{
		Email:    "mymail@mail.com",
		Password: "1234",
		Handle:   "coolhandle",
	})

	testCases := []struct {
		desc     string
		email    string
		password string
	}{
		{
			desc:     "Empty email",
			email:    "",
			password: "1234",
		},
		{
			desc:     "Empty password",
			email:    "mymail@mail.com",
			password: "",
		},
		{
			desc:     "Invalid email",
			email:    "wrongemail@mail.com",
			password: "1234",
		},
		{
			desc:     "Invalid password",
			email:    "mymail@mail.com",
			password: "wrong password",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w, _ := req(r, t, "POST", "/api/users/login", tC)
			assert.Equal(t, http.StatusUnauthorized, w.Code)
		})
	}
}

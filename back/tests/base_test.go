package tests_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/gin-gonic/gin"
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

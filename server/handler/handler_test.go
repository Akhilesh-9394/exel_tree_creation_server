package handler_test

import (
	"bytes"
	"fmt"
	"gin_server/server/handler"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetStudentsHandler(t *testing.T) {
	r := gin.New()
	r.GET("/getstudents", handler.GetStudents)

	req, err := http.NewRequest("GET", "/getstudents", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestStudentHandler(t *testing.T) {
	r := gin.New()
	r.POST("/upload", handler.StudentHandler)
	file, _ := os.Open("./students.xlsx")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	fmt.Println()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

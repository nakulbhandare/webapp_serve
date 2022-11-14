package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/test-web/auth"
	"github.com/test-web/models"
	"github.com/test-web/utils"
	"github.kyndryl.net/MCMP-Development/go-utils/pkg/httpx"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("request", r.URL)
	httpx.WriteJSON(w, "test")
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	log.Println("request", r.URL)
	enableCors(&w)
	var students []interface{}
	student := models.Student{
		ID:   "test",
		Name: "Nakul",
		Age:  23,
	}
	for i := 0; i <= 10; i++ {
		students = append(students, student)

	}

	httpx.WriteJSON(w, students)
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	fmt.Println("headers", r.Header)
	enableCors(&w)

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	apikey := r.URL.Query().Get("Access")
	if apikey != "" {
		if apikey != utils.APIKEY {
			return
		} else {
			token, err := auth.GenerateJWT(username, password)
			fmt.Println("token", token, "error", err)
			if err != nil {
				return
			}
			tokenResponse := models.TokenResponse{
				Token: token,
			}
			httpx.WriteJSON(w, tokenResponse)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, "login successFull")
}

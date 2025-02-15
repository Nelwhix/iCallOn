package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type baseResponse struct {
	Message string `json:"message"`
}

type okResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserAttributes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes UserAttributes `json:"attributes"`
}

type UserWithToken struct {
	ID         string                  `json:"id"`
	Type       string                  `json:"type"`
	Attributes UserAttributesWithToken `json:"attributes"`
}

type Game struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes GameAttributes `json:"attributes"`
}

type GameAttributes struct {
	UserID      string    `json:"userID"`
	RoundLength int       `json:"roundLength"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UserAttributesWithToken struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type GenerateOtp struct {
	OtpUrl string `json:"otpUrl"`
}

func NewInternalServerError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	response := baseResponse{
		Message: message,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}
}

func NewUnauthorized(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnauthorized)
	response := baseResponse{
		Message: message,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}
}

func NewUnprocessableEntity(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	response := baseResponse{
		Message: message,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}
}

func NewBadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	response := baseResponse{
		Message: message,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}
}

func NewNotFoundResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	response := baseResponse{
		Message: message,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}
}

func NewOKResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	response := okResponse{
		Message: message,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}
}

func NewOKResponseWithData(w http.ResponseWriter, message string, data interface{}) {
	w.WriteHeader(http.StatusOK)
	response := okResponse{
		Message: message,
		Data:    data,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func NewCreatedResponseWithData(w http.ResponseWriter, message string, data interface{}) {
	w.WriteHeader(http.StatusCreated)
	response := okResponse{
		Message: message,
		Data:    data,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func NewRedirect(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusSeeOther)
	response := baseResponse{
		Message: message,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Fatal(err)
	}
}

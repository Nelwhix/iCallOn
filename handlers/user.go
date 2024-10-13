package handlers

import (
	"encoding/json"
	"github.com/Nelwhix/iCallOn/pkg"
	"github.com/Nelwhix/iCallOn/pkg/models"
	"github.com/Nelwhix/iCallOn/pkg/requests"
	"github.com/Nelwhix/iCallOn/pkg/responses"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	request, err := pkg.ParseRequestBody[requests.SignUp](w, r)
	if err != nil {
		responses.NewUnprocessableEntity(w, err.Error())

		return
	}

	err = h.Validator.Struct(request)
	if err != nil {
		responses.NewUnprocessableEntity(w, err.Error())
		return
	}

	err = pkg.StrictPasswordValidation(request.Password)
	if err != nil {
		responses.NewUnprocessableEntity(w, err.Error())
		return
	}

	_, err = h.Model.GetUserByEmail(r.Context(), request.Email)
	if err == nil {
		responses.NewBadRequest(w, "Email already taken")
		return
	}

	user, err := h.Model.InsertIntoUsers(r.Context(), request)
	if err != nil {
		responses.NewInternalServerErrorResponse(w, err.Error())
		return
	}

	response := responses.User{
		ID:   user.ID,
		Type: "user",
		Attributes: responses.UserAttributes{
			Username: user.Username,
			Email:    user.Email,
		},
	}

	responses.NewCreatedResponseWithData(w, "User created successfully.", response)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.NewUnprocessableEntity(w, err.Error())
		return
	}
	defer r.Body.Close()

	var request requests.Login
	err = json.Unmarshal(body, &request)
	if err != nil {
		responses.NewUnprocessableEntity(w, err.Error())
		return
	}

	user, err := h.Model.GetUserByEmail(r.Context(), request.Email)
	if err != nil {
		responses.NewBadRequest(w, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		responses.NewBadRequest(w, "Email or Password is incorrect")
		return
	}

	token, err := pkg.CreateToken(h.Model, user.ID)
	if err != nil {
		responses.NewInternalServerErrorResponse(w, err.Error())
		return
	}

	response := responses.UserWithToken{
		ID:   user.ID,
		Type: "user",
		Attributes: responses.UserAttributesWithToken{
			Username: user.Username,
			Email:    user.Email,
			Token:    token,
		},
	}

	responses.NewOKResponseWithData(w, "Login success.", response)
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		responses.NewInternalServerErrorResponse(w, "User not found")
		return
	}

	response := responses.User{
		ID:   user.ID,
		Type: "user",
		Attributes: responses.UserAttributes{
			Username: user.Username,
			Email:    user.Email,
		},
	}

	responses.NewOKResponseWithData(w, "Get user.", response)
}

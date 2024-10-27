package handlers

import (
	"github.com/Nelwhix/iCallOn/pkg"
	"github.com/Nelwhix/iCallOn/pkg/models"
	"github.com/Nelwhix/iCallOn/pkg/requests"
	"github.com/Nelwhix/iCallOn/pkg/responses"
	"net/http"
)

func (h *Handler) CreateNewGame(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		responses.NewInternalServerError(w, "User not found")
		return
	}

	request, err := pkg.ParseRequestBody[requests.NewGame](r)
	if err != nil {
		responses.NewUnprocessableEntity(w, err.Error())

		return
	}

	err = h.Validator.Struct(request)
	if err != nil {
		responses.NewUnprocessableEntity(w, err.Error())
		return
	}

	if request.RoundLength == 0 {
		request.RoundLength = 120 // 2 minutes
	}

	request.UserID = user.ID

	cGame, err := h.Model.InsertIntoGames(r.Context(), request)
	if err != nil {
		responses.NewInternalServerError(w, err.Error())

		return
	}

	response := responses.Game{
		ID:   cGame.ID,
		Type: "game",
		Attributes: responses.GameAttributes{
			UserID:      cGame.UserID,
			RoundLength: cGame.RoundLength,
			CreatedAt:   cGame.CreatedAt,
		},
	}

	responses.NewCreatedResponseWithData(w, "Game created successfully!", response)
}

package authhandler

import (
	"encoding/json"
	"net/http"

	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
	"github.com/AdityaTaggar05/annora-auth/pkg/response"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Error decoding request body", nil)
		return
	}

	tokens, err := h.Service.Login(r.Context(), req.Email, req.Password)

	if err != nil {
		switch err {
			case authservice.ErrInvalidEmailFormat, authservice.ErrInvalidPasswordFormat:
				response.BadRequest(w, "Invalid data format", nil)
			case authservice.ErrUserNotFound, authservice.ErrIncorrectPassword:
				response.Unauthorized(w, "Invalid Credentials")
			case authservice.ErrEmailNotVerified:
				response.Forbidden(w, "Email not verified. Verify it from the link sent to your email!")
			default:
				response.InternalServerError(w, err.Error())
		}
		return
	}

	response.JSON(w, http.StatusOK, tokens);
}

package authhandler

import (
	"net/http"

	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
	"github.com/AdityaTaggar05/annora-auth/pkg/response"
)

func (h *Handler) HandleVerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if token == "" {
		response.BadRequest(w, "Token is required", nil)
		return
	}

	err := h.Service.VerifyEmail(r.Context(), token)

	if err != nil {
		switch err {
			case authservice.ErrInvalidToken:
				response.BadRequest(w, "Invalid verification token", nil)
			default:
				response.InternalServerError(w, err.Error())
		}
		return
	}

	response.Success(w, nil, "Email verified successfully!")
}

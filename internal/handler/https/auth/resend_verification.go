package authhandler

import (
	"encoding/json"
	"net/http"

	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
	"github.com/AdityaTaggar05/annora-auth/pkg/response"
)

func (h *Handler) HandleResendVerification(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.ResendVerification(r.Context(), req.Email); err != nil {
		switch err {
			case authservice.ErrInvalidEmailFormat:
				response.BadRequest(w, "Invalid email format", nil)
			case authservice.ErrUserNotFound:
				response.NotFound(w, "User not found")
			case authservice.ErrTooManyRequests:
				response.Error(w, http.StatusTooManyRequests, "TOO_MANY_REQUESTS", "Too many requests. Try again later!", nil)
			default:
				response.InternalServerError(w, err.Error())
		}
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

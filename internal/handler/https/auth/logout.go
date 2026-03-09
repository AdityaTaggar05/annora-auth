package authhandler

import (
	"encoding/json"
	"net/http"

	tokenservice "github.com/AdityaTaggar05/annora-auth/internal/service/token"
	"github.com/AdityaTaggar05/annora-auth/pkg/response"
)

type logoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	var req logoutRequest

	json.NewDecoder(r.Body).Decode(&req)

	if err := h.Service.Logout(r.Context(), req.RefreshToken); err != nil {
		switch err {
		case tokenservice.ErrInvalidRefreshTokenFormat:
			response.BadRequest(w, "Invalid refresh token format", nil)
		default:
			response.InternalServerError(w, err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

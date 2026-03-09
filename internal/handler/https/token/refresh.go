package tokenhandler

import (
	"encoding/json"
	"net/http"

	tokenservice "github.com/AdityaTaggar05/annora-auth/internal/service/token"
	"github.com/AdityaTaggar05/annora-auth/pkg/response"
)

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	var req refreshRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Error decoding request body", nil)
		return
	}

	tokens, err := h.Service.Refresh(r.Context(), req.RefreshToken)
	if err != nil {
		switch err {
			case tokenservice.ErrInvalidRefreshTokenFormat, tokenservice.ErrInvalidRefreshToken:
				response.BadRequest(w, "Invalid refresh token", nil)
			default:
				response.InternalServerError(w, err.Error())
		}
		return
	}
	
	response.JSON(w, http.StatusOK, tokens)
}

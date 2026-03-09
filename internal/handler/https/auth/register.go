package authhandler

import (
	"encoding/json"
	"net/http"

	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
	"github.com/AdityaTaggar05/annora-auth/pkg/response"
)

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Error decoding request body", nil)
		return
	}

	if err := h.Service.Register(r.Context(), req.Email, req.Password); err != nil {
		switch err {
			case authservice.ErrInvalidEmailFormat, authservice.ErrInvalidPasswordFormat:
				response.BadRequest(w, "Invalid data format", nil)
			case authservice.ErrUserAlreadyExists:
				response.Error(w, http.StatusConflict, "STATUS_CONFLICT", "User already exists", nil)
			default:
				response.InternalServerError(w, err.Error())
		}
		return
	}
	
	response.Created(w, nil, "user registered successfully!")
}

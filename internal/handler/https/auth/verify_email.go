package authhandler

import (
	"net/http"

	"github.com/redis/go-redis/v9"
)

func (h *Handler) HandleVerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if token == "" {
		http.Error(w, "token is required", http.StatusBadRequest)
		return
	}

	err := h.Service.VerifyEmail(r.Context(), token)

	if err != nil {
		switch err {
		case redis.Nil:
			http.Error(w, "invalid or expired token", http.StatusBadRequest)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte("email verified successfully"))
	w.WriteHeader(http.StatusOK)
}

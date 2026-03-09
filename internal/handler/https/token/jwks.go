package tokenhandler

import (
	"net/http"

	"github.com/AdityaTaggar05/annora-auth/pkg/response"
)

func (h *Handler) HandleJWKS(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, h.Service.JWKS())
}

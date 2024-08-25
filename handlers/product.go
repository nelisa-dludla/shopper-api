package handlers

import (
	"net/http"

	"github.com/nelisa-dludla/shopper-api/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
  utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}

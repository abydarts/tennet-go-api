package controllers

import (
	"net/http"

	"github.com/abydarts/tennet-go-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Tennet Test API")

}
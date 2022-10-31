package controllers

import "github.com/abydarts/tennet-go-api/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/wallets", middlewares.SetMiddlewareJSON(s.CreateWallet)).Methods("POST")
	s.Router.HandleFunc("/wallets", middlewares.SetMiddlewareJSON(s.GetWallets)).Methods("GET")
	s.Router.HandleFunc("/wallets/{id}", middlewares.SetMiddlewareJSON(s.GetWallet)).Methods("GET")
	s.Router.HandleFunc("/wallets/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateWallet))).Methods("PUT")
	s.Router.HandleFunc("/wallets/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteWallet)).Methods("DELETE")
}
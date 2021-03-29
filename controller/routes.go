package controller

import (
	"github.com/gorilla/mux"
	"github.com/saurabh-arch/send-email/middleware"
)

/*
routes.go - This file stores all the routes/API endpoints for the application
*/

// GetRouter returns an instance of gorilla mux router
func GetRouter() *mux.Router {
	// initialize router
	r := mux.NewRouter()

	// Attach routes/API endpoints
	r.HandleFunc("/ping", ping)
	r.HandleFunc("/sendemail", middleware.BasicAuthMiddleware(middleware.Logger(SendEmail))).Methods("POST")

	// return router instance
	return r
}

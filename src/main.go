package main

import (
	"net/http"

	. "./business"

	"github.com/gorilla/mux"
)

func main() {
	apiRouter := mux.NewRouter().
		StrictSlash(false)

	initBusinessRouter(apiRouter)

	http.ListenAndServe(":8000", apiRouter)
}

func initBusinessRouter(parent *mux.Router) {
	// All subrouted requests will be suffixes of the URL pattern /api/business
	businessRouter := parent.PathPrefix("/api/business").
		Subrouter()

	// e.g. /api/business/login/
	businessRouter.HandleFunc("/login/", BusinessLoginHandler).
		Methods("POST")
	businessRouter.HandleFunc("/register/", BusinessRegistrationHandler).
		Methods("POST")
	businessRouter.HandleFunc("/restaurant/{restaurantID}/menu/{menuItemID}", MenuGetHandler)
	businessRouter.HandleFunc("/restaurant/{restaurantID}/menu/{menuItemID}/update/", MenuUpdateHandler).
		Methods("POST")
	businessRouter.HandleFunc("/restaurant/{restaurantID}", RestaurantGetHandler)
	businessRouter.HandleFunc("/restaurant/{restaurantID}/update/", RestaurantUpdateHandler).
		Methods("POST")
	businessRouter.HandleFunc("/discount/create/{restaurantID}/{menuItemID}", DiscountCreateHandler).
		Methods("POST")
}

func initClientRotuer(parent *mux.Router) {
	// clientRouter := parent.PathPrefix("/api/client").
	// Subrouter()
}

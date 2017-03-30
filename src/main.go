package main

import (
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"./business"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize connection to DB
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/naexpire?parseTime=true&charset=utf8")
	panicOnErr(err)
	err = db.Ping()
	panicOnErr(err)

	// Initialize request routing
	apiRouter := mux.NewRouter().
		StrictSlash(false)
	initBusinessRouter(apiRouter, db)
	initClientRotuer(apiRouter)

	// Listen for incoming connections on port 8000
	http.ListenAndServe(":8000", apiRouter)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initBusinessRouter(parent *mux.Router, db *sql.DB) {
	// All subrouted requests will be suffixes of the URL pattern /api/business
	businessRouter := parent.PathPrefix("/api/business").
		Subrouter()

	// e.g. /api/business/login/
	businessRouter.HandleFunc("/login/", business.BusinessLoginHandler).
		Methods("POST")
	businessRouter.Handle("/register/", business.RegistrationHandler{DB: db}).
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
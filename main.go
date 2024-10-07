// package main

// import (
// 	"jwt-auth/handler"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/login", handler.Login)
// 	http.HandleFunc("/signup", handler.Signup)
// 	http.HandleFunc("/home", handler.Home)       //Protected
// 	http.HandleFunc("/refresh", handler.Refresh) //Protected
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import (
	"log"
	"net/http"

	"jwt-auth/db"
	"jwt-auth/handler" // Replace with your actual module name
	"jwt-auth/utils"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	pgdb := db.NewMydb()
	// defer pgdb

	// Define your route handlers
	mux.HandleFunc("/login", handler.Login(pgdb))
	mux.HandleFunc("/signup", handler.Signup(pgdb))
	mux.HandleFunc("/home", utils.AuthMiddleware((handler.Home)))
	// Protected
	mux.HandleFunc("/refresh", handler.Refresh) // Protected
	mux.HandleFunc("/check-auth", utils.AuthMiddleware(handler.CheckAuth))
	mux.HandleFunc("/getuser/{id}", utils.AuthMiddleware(handler.GetUserById(pgdb)))
	mux.HandleFunc("/getuser/{id}/workoutplan", utils.AuthMiddleware(handler.GetUserWorkOutPlan(pgdb)))
	mux.HandleFunc("/getuser/{id}/cardioplan", utils.AuthMiddleware(handler.GetUserCardioPlan(pgdb)))

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"}, // Replace with your Flutter app's URL
		AllowedMethods:      []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:      []string{"*"},
		ExposedHeaders:      []string{"*"},
		AllowPrivateNetwork: true,
		AllowCredentials:    true,
		// OptionsPassthrough: true,
		Debug: true,
	})

	// Wrap your router with the CORS handler
	handler := c.Handler(mux)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", handler))
}

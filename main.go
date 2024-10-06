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

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	_ = db.NewMySQLdb()

	// Define your route handlers
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/signup", handler.Signup)
	mux.HandleFunc("/home", handler.Home)       // Protected
	mux.HandleFunc("/refresh", handler.Refresh) // Protected
	mux.HandleFunc("/check-auth", handler.CheckAuth)

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080/*", "http://localhost:8081/*", "http://localhost:8080/", "http://localhost:8081"}, // Replace with your Flutter app's URL
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Set-Cookie", "Content-Length", "Date"},
		AllowCredentials: true,
		// OptionsPassthrough: true,
		Debug: true,
	})

	// Wrap your router with the CORS handler
	handler := c.Handler(mux)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", handler))
}

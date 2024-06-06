package main

import (
	"PurchaseOrderSystem/routes"
	"PurchaseOrderSystem/services"
	"PurchaseOrderSystem/utils"
	pages "PurchaseOrderSystem/views/pages"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	utils.SetTokenAuth(os.Getenv("JWT_SECRET_KEY"))
	err := services.InitDB()
	if err != nil {
		log.Fatalf("database initilisation error: %v", err)
	}
	testData := false
	// Insert test data into the database
	if testData {
		services.TestData()
	}
}

// TODO: Need to find out why the LoggedInMiddleware is not being attached
func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// creating a fileserver
	dir := http.Dir("./static")
	fs := http.FileServer(dir)

	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	router.Mount("/", routes.UserRouter())
	router.Mount("/admin", routes.AdminRouter())

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		pages.Error404().Render(context.Background(), w)
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Jesus is watching, don't do something stupid"))
		return
	})

	http.ListenAndServe(os.Getenv("PORT"), router)
}

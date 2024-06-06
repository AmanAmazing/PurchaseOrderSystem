package routes

import (
	"PurchaseOrderSystem/middlewares"
	"PurchaseOrderSystem/services"
	"PurchaseOrderSystem/utils"
	pages "PurchaseOrderSystem/views/pages"
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(middlewares.LoggedInRedirector)

		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			pages.GetLogin().Render(context.Background(), w)
		})
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// getting the jwt token
		jwtToken, user_role, err := services.PostLogin(username, password)
		if err != nil {
			if err.Error() == "Invalid username or password" {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			} else {
				// TODO: Send Errors back to the form
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    jwtToken,
			HttpOnly: true,
			// secure:   true, // set to true if using https
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(7 * 24 * time.Hour),
		})

		// redirecting
		switch user_role {
		case "admin":
			http.Redirect(w, r, "/admin/home", http.StatusFound)
		case "manager":
			http.Redirect(w, r, "/home", http.StatusFound)
		case "user":
			http.Redirect(w, r, "/home", http.StatusFound)
		default:
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})

	r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    "",
			HttpOnly: true,
			// secure:   true, // set to true if using https
			SameSite: http.SameSiteStrictMode,
			MaxAge:   -1,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)

	})

	// for authenticated routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the public router"))
	})
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the home page"))
	})

	return r
}

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the admin router"))
	})

	return r
}

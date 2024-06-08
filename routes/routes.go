package routes

import (
	"PurchaseOrderSystem/middlewares"
	"PurchaseOrderSystem/services"
	"PurchaseOrderSystem/utils"
	"PurchaseOrderSystem/views/components"
	pages "PurchaseOrderSystem/views/pages"
	"context"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()

	// For routes that logged in users must not use
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(middlewares.LoggedInRedirector)

		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			pages.GetLogin().Render(context.Background(), w)
		})
	})

	// for routes that the user needs to be logged in for
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(middlewares.UnloggedInRedirector)

		r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("HX-Request") != "" {
				// TODO: Need to return home component
				// Also need to return <title> Home </titl> so that it changes the title correctly
				time.Sleep(time.Second * 2)
				w.Header().Set("HX-Push", "/home")
				w.Write([]byte("<title> Home</title><h1>SOmethign</h1>"))
				return
			}
			// don't need currentPage anymore
			data := map[string]interface{}{
				"title": "home",
			}
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			pages.GetHome(data).Render(context.Background(), w)
		})

		r.Get("/form", func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("HX-Request") != "" {
				time.Sleep(time.Second * 2)
				w.Header().Set("HX-Push", "/form")
				components.PurchaseOrderForm().Render(context.Background(), w)
				return
			}
			data := map[string]interface{}{
				"currentPage": "form",
			}
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			pages.GetForm(data).Render(context.Background(), w)
		})

		r.Get("/form/get/item", func(w http.ResponseWriter, r *http.Request) {
			item_count_string := r.URL.Query().Get("item_count")
			item_count, err := strconv.Atoi(item_count_string)
			item_count += 1
			if err != nil {
				// TODO: Need to return error
				return
			}
			w.Header().Add("HX-Trigger", "updateItemCountEvet")
			w.Header().Add("HX-Trigger-After-Swap", "updateItemCountEvent")
			components.PurchaseOrderFormItem(item_count).Render(context.Background(), w)
		})
		r.Get("/form/get/suppliers", func(w http.ResponseWriter, r *http.Request) {
			suppliers, err := services.GetSuppliers()
			if err != nil {
				println("error occurred fetching suppliers")
				return
			}
			tmpl := template.Must(template.New("suppliers").Parse(`
			{{range .}}
				<option value="{{.ID}}">{{.Name}}</option>
			{{end}}
			`))
			tmpl.Execute(w, suppliers)
		})
		r.Get("/form/get/nominals", func(w http.ResponseWriter, r *http.Request) {
			nominals, err := services.GetNominals()
			if err != nil {
				println("error occurred fetching nominals")
			}
			tmpl := template.Must(template.New("nominals").Parse(`
			{{range .}}
				<option value="{{.ID}}">{{.Name}}</option>
			{{end}}
		`))
			tmpl.Execute(w, nominals)

		})
		r.Get("/form/get/products", func(w http.ResponseWriter, r *http.Request) {
			products, err := services.GetProducts()
			if err != nil {
				println("error occurred fetching products")
			}
			tmpl := template.Must(template.New("products").Parse(`
			{{range .}}
				<option value="{{.ID}}">{{.Name}}</option>
			{{end}}
		`))
			tmpl.Execute(w, products)

		})

		r.Get("/repo", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Partial repo"))
		})

		r.Get("/notifications", func(w http.ResponseWriter, r *http.Request) {
			// TODO: Full notifications page if the request is not a htmx request.
			// need to decide whether notifications should be a full page or a popout
		})

	})

	// public routes
	r.Group(func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Welcome to the public router"))
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

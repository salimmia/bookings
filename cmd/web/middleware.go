package main

import (
	"context"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/salimmia/bookings/internal/helpers"
	"github.com/salimmia/bookings/internal/models"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r){
			session.Put(r.Context(), "error", "Log in first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetUserMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Create or fetch the User object based on your authentication logic
        // Here, we assume you have a function named "getUserFromSession" that retrieves the user from the session or authentication mechanism
        user_information := getUserFromSession(r)

        // Set the User object in the request context
        ctx := context.WithValue(r.Context(), "user_information", user_information)
        r = r.WithContext(ctx)

		// log.Println(ctx.Value("user_information"))

        // Call the next handler in the chain
        next.ServeHTTP(w, r)
    })
}

func getUserFromSession(r *http.Request) models.User{
	var user models.User

	exists := app.Session.Exists(r.Context(), "user_information")

	if !exists{
		return user
	}

	// fmt.Printf("exists: %v\n", exists)

	user = app.Session.Get(r.Context(), "user_information").(models.User)

	// log.Println(user.FirstName)

	return user
}
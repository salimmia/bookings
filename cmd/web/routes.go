package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/salimmia/bookings/internal/config"
	"github.com/salimmia/bookings/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)
	mux.Get("/user/forgot-password", handlers.Repo.ForgotPassword)
	mux.Post("/user/change-password", handlers.Repo.ChangePassword)
	mux.Get("/user/reset-password", handlers.Repo.ResetPassword)
	mux.Post("/user/reset-password", handlers.Repo.PostResetPassword)
	mux.Get("/user/signup", handlers.Repo.ShowSignup)
	mux.Post("/user/signup", handlers.Repo.PostShowSignup)
	mux.Get("/user/logout", handlers.Repo.LogOut)
	mux.Get("/user/profile", handlers.Repo.UserProfile)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		// for authentication
		mux.Use(Auth)
		
		mux.Get("/dashboard", handlers.Repo.AdminDashboard)

		mux.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		mux.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		mux.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
		mux.Post("/reservations-calendar", handlers.Repo.AdminPostReservationsCalendar)
		mux.Get("/process-reservation/{src}/{id}/do", handlers.Repo.AdminProcessReservation)
		mux.Get("/delete-reservation/{src}/{id}/do", handlers.Repo.AdminDeleteReservation)

		mux.Get("/reservations/{src}/{id}/show", handlers.Repo.AdminShowReservation)
		mux.Post("/reservations/{src}/{id}", handlers.Repo.AdminPostShowReservation)

	})

	return mux
}

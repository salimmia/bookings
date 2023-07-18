package helpers

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/salimmia/bookings/internal/config"
	"github.com/salimmia/bookings/internal/models"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAuthenticated(r *http.Request) bool{
	exists := app.Session.Exists(r.Context(), "user_id")

	return exists
}

func IsAdmin(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "is_admin")
	// log.Println(exists)
	return exists
}

func UserInformation(r *http.Request) map[string]interface{} {
	user := app.Session.Get(r.Context(), "user_information").(models.User)

	user_information := make(map[string]interface{})

	user_information["id"] = user.ID
	user_information["first_name"] = user.FirstName
	user_information["last_name"] = user.LastName
	user_information["phone"] = user.Phone
	user_information["email"] = user.Email
	user_information["password"] = user.Password
	user_information["created_at"] = user.CreatedAt
	user_information["updated_at"] = user.UpdatedAt

	log.Println(user_information)

	return user_information
}
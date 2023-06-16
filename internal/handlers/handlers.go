package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/salimmia/bookings/internal/config"
	"github.com/salimmia/bookings/internal/driver"
	"github.com/salimmia/bookings/internal/forms"
	"github.com/salimmia/bookings/internal/helpers"
	"github.com/salimmia/bookings/internal/models"
	"github.com/salimmia/bookings/internal/render"
	"github.com/salimmia/bookings/internal/repository"
	"github.com/salimmia/bookings/internal/repository/dbrepo"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewMysqlRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        helpers.ServerError(w, err)
        return
    }

    sd := r.Form.Get("start_date")
    ed := r.Form.Get("end_date")

    // 2020-01-01 01/02 03:04:05PM '06 -0700
    layout := "2006-01-02"

    startDate, _:= time.Parse(layout, sd)
    // if err != nil{
    //     helpers.ServerError(w, err)
    // }

    endDate, _:= time.Parse(layout, ed)
    // if err != nil{
    //     helpers.ServerError(w, err)
    // }

    roomID, err:= strconv.Atoi(r.Form.Get("room_id"))
    if err != nil{
        helpers.ServerError(w, err)
    }

    reservation := models.Reservation{
        FirstName: r.Form.Get("first_name"),
        LastName:  r.Form.Get("last_name"),
        Email:     r.Form.Get("email"),
        Phone:     r.Form.Get("phone"),
        StartDate: startDate,
        EndDate: endDate,
        RoomID: roomID,
    }

    form := forms.New(r.PostForm)

    form.Required("first_name", "last_name", "email")
    form.MinLength("first_name", 3)
    // form.IsEmail("email")

    if !form.Valid() {
        data := make(map[string]interface{})
        data["reservation"] = reservation
        render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
            Form: form,
            Data: data,
        })
        return
    }

    NewReservationID, err := m.DB.InsertReservation(reservation)

    if err != nil{
        helpers.ServerError(w, err)
		return
    }
	log.Println(NewReservationID)

	restriction := models.RoomRestriction{
		StartDate: startDate,
		EndDate: endDate,
		RoomID: roomID,
		ReservationID: NewReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)

    m.App.Session.Put(r.Context(), "reservation", reservation)
    http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles post
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	layout := "2006-01-02"

    startDate, _:= time.Parse(layout, start)
    // if err != nil{
    //     helpers.ServerError(w, err)
    // }

    endDate, _:= time.Parse(layout, end)
    // if err != nil{
    //     helpers.ServerError(w, err)
    // }

	rooms, err := m.DB.SearchAvailabilityByDatesByAllRooms(startDate, endDate)

	if err != nil{
		helpers.ServerError(w, err)
		return
	}
	// for _, i:= range rooms{
	// 	m.App.InfoLog.Println("Room: ", i.ID, i.RoomName)
	// }

	if len(rooms) == 0{
		// no Available
		// m.App.InfoLog.Println("No Available!")
		m.App.Session.Put(r.Context(), "error", "No Availability")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: startDate,
		EndDate: endDate,
	}
	m.App.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// ReservationSummary displays the res summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Can't get reservation from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request){
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil{
		helpers.ServerError(w, err)
		return
	}
	
	// m.App.Session.Get(r.Context(), "reservation")

	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID

	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}

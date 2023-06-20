package repository

import (
	"time"

	"github.com/salimmia/bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
	
	GetUserByID(id int) (models.User, error)
	UpdatedUser(user models.User) error
	Authenticate(email, testPassword string) (int, string, error)
	
	AllReservation() ([]models.Reservation, error)
}

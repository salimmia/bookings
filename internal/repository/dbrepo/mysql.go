package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/salimmia/bookings/internal/models"
)

func (m *MysqlDBRepo) AllUsers() bool {
	return true
}

// InsertReservation insert a reservation into the database
func (m *MysqlDBRepo) InsertReservation(res models.Reservation) (int, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	
	stmt := `
		INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)

	if err != nil{
		return 0, err
	}

	var lastInsertID int
	err = m.DB.QueryRowContext(ctx, "SELECT LAST_INSERT_ID()").Scan(&lastInsertID)
	if err != nil {
		log.Fatal(err)
	}


	return lastInsertID, nil
}

// InsertRoomRestriction insert a restriction into the database
func (m *MysqlDBRepo) InsertRoomRestriction(r models.RoomRestriction) error{
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	stmt := `
		INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		time.Now(),
		time.Now(),
		r.RestrictionID,
	)
	if err != nil{
		return err
	}

	return nil
}

func (m *MysqlDBRepo) SearchAvailabilityByDates(start, end time.Time, roomId int) (bool, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var numRows int
	query:= `
	SELECT COUNT(id)
		FROM room_restrictions
		WHERE room_id = ?
			AND ? < end_date and ? > start_date;
	`
	row := m.DB.QueryRowContext(ctx, query,
		roomId,
		start,
		end,
	)
	err := row.Scan(&numRows)

	if err != nil{
		return false, err
	}
	if numRows == 0{
		return true, nil
	}
	
	return false, nil
}
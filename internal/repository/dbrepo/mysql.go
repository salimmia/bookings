package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/salimmia/bookings/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *mysqlDBRepo) AllUsers() bool {
	return true
}

// InsertReservation insert a reservation into the database
func (m *mysqlDBRepo) InsertReservation(res models.Reservation) (int, error){
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
func (m *mysqlDBRepo) InsertRoomRestriction(r models.RoomRestriction) error{
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

func (m *mysqlDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomId int) (bool, error){
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

// SearchAvailabilityByDatesByAllRooms returns a slice of available rooms for any given date range
func (m *mysqlDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
		SELECT r.id, r.room_name
		FROM rooms AS r
		WHERE r.id NOT IN(
			SELECT rr.room_id
			FROM room_restrictions AS rr
			WHERE ? < rr.end_date and ? > rr.start_date
		)
	`
	rows, err:= m.DB.QueryContext(ctx, query,
		start,
		end,
	)
	if err != nil{
		return rooms, err
	}

	for rows.Next(){
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil{
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil{
		return rooms, err
	}

	return rooms, nil
}

func (m *mysqlDBRepo) GetRoomByID(id int) (models.Room, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var room models.Room

	query := `
		SELECT id, room_name
		FROM rooms
		WHERE id = ?
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&room.ID,
		&room.RoomName,
		// &room.CreatedAt, error for formating time.Time
	)
	if err != nil{
		return room, err
	}

	return room, nil
}

//GetUserByID returns a user details by ID
func (m *mysqlDBRepo) GetUserByID(id int) (models.User, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var user models.User

	query := `
	SELECT id, first_name, last_name, email, password, access_level, created_at, updated_at
	FROM users
	WHERE id = ?
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.AccessLevel,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil{
		return user, err
	}

	return user, nil
}

//UpdatedUser Update user details
func (m *mysqlDBRepo) UpdateUser(user models.User) error{
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
	UPDATE users SET first_name = ?, last_name = ?, email = ?, access_level = ?, updated_at = ?
	`

	_, err := m.DB.ExecContext(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.AccessLevel,
		time.Now(),
	)

	if err != nil{
		return err
	}

	return nil
}

// Authenticate authenticate a user
func (m *mysqlDBRepo) Authenticate(email, testPassword string) (int, string, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		select id, password from users where email = ?
	`

	var id int
	var hashedPassword string

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&id,
		&hashedPassword,
	)
	if err != nil{
		return id, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword{
		return 0, "", errors.New("incorrect password")
	}else if err != nil{
		return 0, "", err
	}

	return id, hashedPassword, nil
}

// AllReservation returns a slice of all reservations
func (m *mysqlDBRepo) AllReservation() ([]models.Reservation, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `
		SELECT r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, r.processed,
		rm.id, rm.room_name
		FROM reservations r 
		LEFT JOIN rooms rm
		ON(r.room_id = rm.id)
		ORDER BY r.start_date ASC
	`
	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil{
		return reservations, err
	}
	defer rows.Close()

	for rows.Next(){
		var reservation models.Reservation
		var start_date, end_date, created_at, updated_at []uint8

		var start_date_value, end_date_value, created_at_value, updated_at_value *time.Time

		err := rows.Scan(
			&reservation.ID,
			&reservation.FirstName,
			&reservation.LastName,
			&reservation.Email,
			&reservation.Phone,
			&start_date,
			&end_date,
			&reservation.RoomID,
			&created_at,
			&updated_at,
			&reservation.Processed,
			&reservation.Room.ID,
			&reservation.Room.RoomName,
		)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}

		start_dateStr := string(start_date)
		parsedTimeStart, err := time.Parse("2006-01-02", start_dateStr)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}
		start_date_value = &parsedTimeStart

		end_dateStr := string(end_date)
		parsedTimeeEnd, err := time.Parse("2006-01-02", end_dateStr)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}
		end_date_value = &parsedTimeeEnd

		created_atStr := string(created_at)
		parsedTimeeCreated, err := time.Parse("2006-01-02 15:04:05", created_atStr)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}
		created_at_value = &parsedTimeeCreated

		updated_atStr := string(updated_at)
		parsedTimeeUpdated, err := time.Parse("2006-01-02 15:04:05", updated_atStr)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}
		updated_at_value = &parsedTimeeUpdated

		reservation.StartDate = *start_date_value
		reservation.EndDate = *end_date_value
		reservation.CreatedAt = *created_at_value
		reservation.UpdatedAt = *updated_at_value

		// log.Println(reservation.StartDate)

		reservations = append(reservations, reservation)
	}

	if err = rows.Err(); err != nil{
		return reservations, err
	}

	return reservations, nil
}

// NewReservation returns a slice of new reservations
func (m *mysqlDBRepo) NewReservation() ([]models.Reservation, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `
		SELECT r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, r.processed, 
		rm.id, rm.room_name
		FROM reservations r 
		LEFT JOIN rooms rm
		ON(r.room_id = rm.id)
		where r.processed = 0
		ORDER BY r.start_date ASC
	`
	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil{
		return reservations, err
	}
	defer rows.Close()

	for rows.Next(){
		var reservation models.Reservation
		var start_date, end_date, created_at, updated_at []uint8

		var start_date_value, end_date_value, created_at_value, updated_at_value *time.Time

		err := rows.Scan(
			&reservation.ID,
			&reservation.FirstName,
			&reservation.LastName,
			&reservation.Email,
			&reservation.Phone,
			&start_date,
			&end_date,
			&reservation.RoomID,
			&created_at,
			&updated_at,
			&reservation.Processed,
			&reservation.Room.ID,
			&reservation.Room.RoomName,
		)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}

		start_dateStr := string(start_date)
		parsedTimeStart, err := time.Parse("2006-01-02", start_dateStr)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}
		start_date_value = &parsedTimeStart

		end_dateStr := string(end_date)
		parsedTimeeEnd, err := time.Parse("2006-01-02", end_dateStr)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}
		end_date_value = &parsedTimeeEnd

		created_atStr := string(created_at)
		parsedTimeeCreated, err := time.Parse("2006-01-02 15:04:05", created_atStr)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}
		created_at_value = &parsedTimeeCreated

		updated_atStr := string(updated_at)
		parsedTimeeUpdated, err := time.Parse("2006-01-02 15:04:05", updated_atStr)
		if err != nil{
			return reservations, err
			// log.Println(err)
		}
		updated_at_value = &parsedTimeeUpdated

		reservation.StartDate = *start_date_value
		reservation.EndDate = *end_date_value
		reservation.CreatedAt = *created_at_value
		reservation.UpdatedAt = *updated_at_value

		// log.Println(reservation.StartDate)

		reservations = append(reservations, reservation)
	}

	if err = rows.Err(); err != nil{
		return reservations, err
	}

	return reservations, nil
}

// AllReservation returns a slice of all reservations
func (m *mysqlDBRepo) GetReservationByID(id int) (models.Reservation, error){
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	var reservation models.Reservation

	query := `
		SELECT r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, r.processed,
		rm.id, rm.room_name
		FROM reservations r 
		LEFT JOIN rooms rm
		ON(r.room_id = rm.id)
		where r.id = ?
		ORDER BY r.start_date ASC
	`
	rows := m.DB.QueryRowContext(ctx, query, id)

	var start_date, end_date, created_at, updated_at []uint8

	var start_date_value, end_date_value, created_at_value, updated_at_value *time.Time

	err := rows.Scan(
		&reservation.ID,
		&reservation.FirstName,
		&reservation.LastName,
		&reservation.Email,
		&reservation.Phone,
		&start_date,
		&end_date,
		&reservation.RoomID,
		&created_at,
		&updated_at,
		&reservation.Processed,
		&reservation.Room.ID,
		&reservation.Room.RoomName,
	)
	if err != nil{
		return reservation, err
		// log.Println(err)
	}

	start_dateStr := string(start_date)
	parsedTimeStart, err := time.Parse("2006-01-02", start_dateStr)
	if err != nil{
		return reservation, err
		// log.Println(err)
	}
	start_date_value = &parsedTimeStart

	end_dateStr := string(end_date)
	parsedTimeeEnd, err := time.Parse("2006-01-02", end_dateStr)
	if err != nil{
		return reservation, err
		// log.Println(err)
	}
	end_date_value = &parsedTimeeEnd

	created_atStr := string(created_at)
	parsedTimeeCreated, err := time.Parse("2006-01-02 15:04:05", created_atStr)
	if err != nil{
		return reservation, err
		// log.Println(err)
	}
	created_at_value = &parsedTimeeCreated

	updated_atStr := string(updated_at)
	parsedTimeeUpdated, err := time.Parse("2006-01-02 15:04:05", updated_atStr)
	if err != nil{
		return reservation, err
		// log.Println(err)
	}
	updated_at_value = &parsedTimeeUpdated

	reservation.StartDate = *start_date_value
	reservation.EndDate = *end_date_value
	reservation.CreatedAt = *created_at_value
	reservation.UpdatedAt = *updated_at_value

	// log.Println(reservation.StartDate)

	if err = rows.Err(); err != nil{
		return reservation, err
	}

	return reservation, nil
}

//UpdatedReservation Update Reservation details
func (m *mysqlDBRepo) UpdateReservation(reservation models.Reservation) error{
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		UPDATE reservations SET first_name = ?, last_name = ?, email = ?, phone = ?, updated_at = ?
	`

	_, err := m.DB.ExecContext(ctx, query,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.Phone,
		time.Now(),
	)

	if err != nil{
		return err
	}

	return nil
}

// DeleteReservation deletes one reservation by id
func (m *mysqlDBRepo) DeleteReservation(id int) error{
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		delete from reservations where id = ?
	`

	_, err := m.DB.ExecContext(ctx, query, id)

	if err != nil{
		return err
	}

	return nil
}

// UpdateProcessedForReservation update processed in reservations table by id
func (m *mysqlDBRepo) UpdateProcessedForReservation(id, processed int) error{
	ctx, cancel:= context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	query := `
		UPDATE reservations SET processed = ? where id = ?
	`
	_, err := m.DB.ExecContext(ctx, query, processed, id)

	if err != nil{
		return err
	}

	return nil
}
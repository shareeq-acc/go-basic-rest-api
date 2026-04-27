package models

import (
	"api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {

	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	res, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	e.ID = id
	return err
}

func (e *Event) Update() error {
	query := `
	UPDATE events SET name = ?, description = ? , location = ?, dateTime = ? , user_id =?
	WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	res, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events where id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Register(userId int64) error {
	query := `
	INSERT INTO Registrations(event_id, user_id)
	VALUES(?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	res, err := statement.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM Registrations
	WHERE user_id = ? and event_id = ? 
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	res, err := statement.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	return err
}

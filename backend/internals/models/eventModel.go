package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	v "github.com/go-ozzo/ozzo-validation"
)

type Location struct {
	Lat  float64 `json:"lat"  db:"lat"`
	Lang float64 `json:"lang" db:"lang"`
}
type Event struct {
	ID               int       `json:"id"                db:"id"`
	Name             string    `json:"name"              db:"name"`
	Date             time.Time `json:"date"              db:"date"`
	Address          string    `json:"address"           db:"address"`
	Category         string    `json:"category"          db:"category"`
	TicketPrice      int       `json:"ticket_price"      db:"price"`
	Description      string    `json:"description"       db:"description"`
	Images           Images    `json:"images"            db:"images"`
	NumberOfTickets  int       `json:"tickets"           db:"number_of_tickets"`
	AvialableTickets int       `json:"avialable_tickets" db:"available_tickets"`
	Organizer        string    `json:"organizer"         db:"organizer"`
	Location         `json:"location"`
}

type Images []string

func (c Images) Value() (driver.Value, error) {
	return json.Marshal(c) // return json marshalled value
}

func (c *Images) Scan(v interface{}) error {
	switch tv := v.(type) {
	// case []byte:
	// return json.Unmarshal(tv, &c) // unmarshal
	case []uint8:
		return json.Unmarshal([]byte(tv), &c) // can't remember the specifics, but this may be needed
	}
	return errors.New("unsupported type")
}

func (event Event) Validate() error {
	return v.ValidateStruct(
		event,
		v.Field(&event.Name, v.Required.Error("The name is requires")),
		v.Field(&event.Date, v.Date(time.ANSIC)),
	)
}

type AddEventBody struct {
	Name             string    `json:"name"              form:"name"`
	Description      string    `json:"description"       form:"description"`
	Address          string    `json:"address"           form:"address"`
	Date             time.Time `json:"date"              form:"date"`
	NumberOfTickets  int       `json:"tickets"           form:"tickets"`
	TicketPrice      int       `json:"ticket_price"      form:"ticket_price"`
	Category         string    `json:"category"          form:"category"`
	Organizer        string    `json:"organizer"         form:"organizer"`
	Images           []string  `json:"images"            form:"images"`
	Location         Location  `json:"location"          form:"location"`
	AvialableTickets int       `json:"available_tickets" form:"avialable_tickets"`
}
type UpdateEventBody struct {
	AddEventBody
	DeletedImages []string `json:"deleted_images" form:"deleted_images[]"`
}

func (e AddEventBody) Validate() error {
	return v.ValidateStruct(
		&e,
		v.Field(&e.Name, v.Required.Error("name is required")),
		v.Field(
			&e.TicketPrice,
			v.Required,
			v.Min(0),
		),
		v.Field(&e.Description, v.Required),
	)
}

func (e UpdateEventBody) Validate() error {
	return v.ValidateStruct(
		&e,
	)
}

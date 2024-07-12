package queries

import (
	"fmt"
	"mime/multipart"

	"github.com/ra00d/event_management/internals/configs"
	"github.com/ra00d/event_management/internals/constants/errors"
	"github.com/ra00d/event_management/internals/models"
)

func GetAllEvents() []models.Event {
	events := []models.Event{}
	db := configs.AppDB
	// db := configs.AppDB
	err := db.Select(
		&events,
		`SELECT
			name,id,address,date,description,category,price,
			number_of_tickets,available_tickets,
			organizer, ST_X(location) AS lat,ST_Y(location) AS lang,
			JSON_ARRAYAGG(i.img_path) AS images
		FROM 
			event 
		LEFT JOIN
			event_images i on i.event_id = event.id
		GROUP BY
			event.id
		`,
	)
	if err != nil {
		panic(err)
	}
	return events
}

func AddEvent(
	event models.AddEventBody,
	images []*multipart.FileHeader,
	saveFunc func(f *multipart.FileHeader, p string) error,
) error {
	db := configs.AppDB
	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		panic(fmt.Sprintf("unable to add the events becaause %s", err.Error()))
	}
	res, err := tx.Exec(`INSERT INTO event
		(name, date, location, address, category, description, number_of_tickets, available_tickets, organizer,price)
                VALUES 
		(?, ?, ST_GeomFromText('POINT(40.7128 -74.0060)'), ?, ?, ?, ?, ?, ?,?);
`, event.Name,
		event.Date,
		event.Address,
		event.Category,
		event.Description,
		event.NumberOfTickets,
		event.NumberOfTickets,
		event.Organizer,
		event.TicketPrice)
	if err != nil {
		panic(fmt.Sprintf("unable to add the events becaause %s", err.Error()))
	}
	eventId, err := res.LastInsertId()
	if err != nil {
		panic(fmt.Sprintf("unable to add the events becaause %s", err.Error()))
	}
	for _, img := range images {
		imgPath := fmt.Sprintf("/uploads/%s", img.Filename)
		err = saveFunc(img, fmt.Sprintf("./storage/uploads/%s", img.Filename))
		if err != nil {
			panic(fmt.Sprintf("unable to add the events becaause %s", err.Error()))
		}

		_, err = tx.Exec(`INSERT INTO event_images 
			(event_id,img_path)
			VALUES 
			(?,?)`, eventId, imgPath)
	}
	if err != nil {
		panic(fmt.Sprintf("unable to add the events becaause %s", err.Error()))
	}
	if err = tx.Commit(); err != nil {
		panic(fmt.Sprintf("unable to add the events becaause %s", err.Error()))
	}
	return nil
}

func GetEvent(id string) models.Event {
	events := models.Event{}
	db := configs.AppDB
	err := db.Get(&events, `SELECT
			name,id,address,date,description,category,price,
			number_of_tickets,available_tickets,
			organizer, ST_X(location) AS lat,ST_Y(location) AS lang,
			JSON_ARRAYAGG(i.img_path) AS images
		FROM 
			event 
		LEFT JOIN
			event_images i on i.event_id = event.id
		WHERE
		 event.id=?
		GROUP BY
			event.id
		`, id)
	if err != nil {
		// panic(fiber.NewError(404, fmt.Sprintf("Evnet with id %s does not exsist", id)))
		fmt.Println(err)
		panic(errors.NewNotFoundError(fmt.Sprintf("Event with id %s does not exist",
			id)))
		// panic(errors.NotFoundError{Code: 404, Message: "this event does not exist"})
	}
	return events
}

func DeleteEvent(id string) (int64, error) {
	db := configs.AppDB
	res, err := db.Exec("DELETE  FROM event WHERE id=?", id)
	if err != nil {
		// panic(fiber.NewError(404, fmt.Sprintf("Evnet with id %s does not exsist", id)))
		panic(errors.NewNotFoundError(fmt.Sprintf("Event with id %s does not exist",
			id)))
		// panic(errors.NotFoundError{Code: 404, Message: "this event does not exist"})
	}
	return res.RowsAffected()
}

func UpdateEvent(
	id string,
	event models.UpdateEventBody,
	images []*multipart.FileHeader,
	saveFunc func(f *multipart.FileHeader, p string) error,
) (models.Event, error) {
	GetEvent(id)
	db := configs.AppDB
	tx, err := db.Begin()
	if err != nil {
		panic(fmt.Sprintf("unable to add the events becaause %s", err.Error()))
	}
	_, err = db.Exec(`UPDATE event 
		SET
		name=?,
		date=?,
		description=?,
		address=?,
		category=?,
		number_of_tickets=?,
		available_tickets=?,
		organizer=?,
		price=? 
		WHERE id=?;`,
		event.Name,
		event.Date,
		event.Description,
		event.Address,
		event.Category,
		event.NumberOfTickets,
		event.AvialableTickets,
		event.Organizer,
		event.TicketPrice,
		id)
	if err != nil {
		panic(fmt.Sprintf("unable to update the events becaause %s", err.Error()))
	}
	// if images == nil {
	// 	panic("unable to update the events becaause no images passed")
	// }
	fmt.Println(event.DeletedImages)
	for _, v := range event.DeletedImages {
		fmt.Println(v)
		_, err = tx.Exec(`DELETE	FROM 
			event_images 
		WHERE
			event_id=? AND img_path=?
		`, id, v)
		if err != nil {
			panic(fmt.Sprintf("unable to update the events becaause %s", err.Error()))
		}
	}
	for _, img := range images {
		imgPath := fmt.Sprintf("/uploads/%s", img.Filename)
		err = saveFunc(img, fmt.Sprintf("./storage/uploads/%s", img.Filename))
		if err != nil {
			panic(fmt.Sprintf("unable to update the events becaause %s", err.Error()))
		}

		_, err = tx.Exec(`INSERT INTO event_images 
			(event_id,img_path)
			VALUES 
			(?,?)`, id, imgPath)
	}
	if err != nil {
		panic(fmt.Sprintf("unable to update the events becaause %s", err.Error()))
	}
	if err = tx.Commit(); err != nil {
		panic(fmt.Sprintf("unable to update the events becaause %s", err.Error()))
	}
	res := GetEvent(id)
	return res, err
}

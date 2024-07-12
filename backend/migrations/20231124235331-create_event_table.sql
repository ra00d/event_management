
-- +migrate Up
CREATE TABLE IF NOT EXISTS event (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date DATETIME NOT NULL,
    location POINT NOT NULL,
    address VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
	price INT,
    description TEXT,
    number_of_tickets INT NOT NULL,
    available_tickets INT NOT NULL,
    organizer VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS event_images (
	event_id INT NOT NULL,
	img_path VARCHAR(255) NOT NULL,
	 PRIMARY KEY(event_id,img_path),

	FOREIGN KEY(event_id) REFERENCES event(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +migrate Down
DROP TABLE IF EXISTS event_images;
DROP TABLE IF EXISTS event;


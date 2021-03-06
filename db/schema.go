package db

import (
	"log"

	"github.com/jinzhu/gorm"
)

var schema = `
	PRAGMA foreign_keys = ON;

	CREATE TABLE IF NOT EXISTS sources (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL
	);

	CREATE TABLE IF NOT EXISTS rss_items (
		id INTEGER PRIMARY KEY,
		source_id INTEGER,
		title TEXT NOT NULL,
		link TEXT NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL,
		FOREIGN KEY (source_id) REFERENCES sources (id)
	);

	CREATE UNIQUE INDEX IF NOT EXISTS rss_items_link_idx ON rss_items(link);
`

// Migrate :nodoc:
func Migrate(db *gorm.DB) {
	err := db.Exec(schema).Error
	if err != nil {
		log.Fatal(err)
	}
}

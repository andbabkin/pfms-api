package base

import (
	"database/sql"
	"time"
)

// AutoIncr defines a standard for autoincremented primary key field
type AutoIncr struct {
	ID uint64
}

// Created defines a standard for a field which stores a record creation date
type Created struct {
	CreatedAt time.Time `db:"created_at"`
}

// Updated defines a standard for a field which stores a date of last record change
type Updated struct {
	UpdatedAt time.Time `db:"updated_at"`
}

// Deleted defines a standard for a field which stores a deletion date of soft-deletable record
type Deleted struct {
	DeletedAt sql.NullTime `db:"deleted_at"`
}

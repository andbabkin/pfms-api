package storage

import (
	"strconv"
)

// FindByIDSQL generates SQL string to find a record in the table by its id.
func FindByIDSQL(id uint64, table string) string {
	return "SELECT * FROM `" + table + "` WHERE id = " + strconv.FormatUint(id, 10) + " LIMIT 1"
}

// FindByFieldSQL prepares SQL string with named variable to find a record in
// the table by a field value.
func FindByFieldSQL(field string, table string) string {
	return "SELECT * FROM `" + table + "` WHERE `" + field + "` = :" + field
}

// UpdateFieldSQL prepares SQL with named variables to update a field value
// in a record with specified id
func UpdateFieldSQL(field string, table string) string {
	return "UPDATE `" + table + "` SET `" + field + "`=:" + field + ", `updated_at`=:updated_at WHERE id = :id"
}

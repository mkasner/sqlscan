// generated by sqlscan -type=Transaction -version=2; DO NOT EDIT

package main

import "database/sql"

// Returns all field names from Transaction
func (t *Transaction) Fields() []string {
	return []string{"HELLO", "WORLD", "TODAY"}
}

// // Scans to Transaction
func (t *Transaction) Scan(rows *sql.Rows) (Transaction, error) {
	var r Transaction
	err := rows.Scan(
		&r.Hello,
		&r.World,
		&r.Today,
	)
	return r, err
}

// // Scans to Transaction
func (t *Transaction) ScanRow(row *sql.Row) (Transaction, error) {
	var r Transaction
	err := row.Scan(
		&r.Hello,
		&r.World,
		&r.Today,
	)
	return r, err
}

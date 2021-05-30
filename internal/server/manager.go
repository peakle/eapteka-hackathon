package server

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/peakle/eapteka-hackathon/internal"
	sg "github.com/wakeapp/go-sql-generator"
)

func addSchedule(m *internal.SQLManager, userID, drug string) error {
	ins := &sg.InsertData{
		TableName: "Schedule",
		IsIgnore:  true,
		Fields:    []string{"userID", "drug", "date", "createdAt", "updatedAt"},
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	ins.Add([]string{userID, drug, date, date, date})

	if _, err := m.Insert(ins); err != nil {
		return err
	}

	return nil
}

func lastSchedule(m *internal.SQLManager, userID, drug string) (string, error) {
	const never = "никогда"
	var query = `
		SELECT date
		FROM Schedule
		WHERE 1
			AND userID = ?
			AND drug = ?
		ORDER BY date DESC
		LIMIT 1
	`
	var rows *sql.Rows
	var stmt, err = m.GetConnection().Prepare(query)
	if err != nil {
		return "", fmt.Errorf("on lastSchedule: %s", err)
	}
	defer stmt.Close()

	rows, err = stmt.Query(userID, drug)
	if err != nil {
		return never, fmt.Errorf("on lastSchedule: %s", err)
	}
	defer rows.Close()

	var date string
	for rows.Next() {
		err = rows.Scan(
			&date,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return never, nil
			}
			return "", fmt.Errorf("on lastSchedule: %s", err)
		}
	}

	if date == "" {
		return never, nil
	}

	return date, nil
}

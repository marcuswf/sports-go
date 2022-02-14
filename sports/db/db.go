package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (e *eventsRepo) seed() error {
	statement, err := e.db.Prepare(`CREATE TABLE IF NOT EXISTS sport_events (id INTEGER PRIMARY KEY, name TEXT, type INTEGER, location TEXT, visible INTEGER, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = e.db.Prepare(`INSERT OR IGNORE INTO sport_events(id, name, type, location, visible, advertised_start_time) VALUES (?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Team().Name(),
				faker.Number().Between(0, 3),
				faker.Address().City(),
				faker.Number().Between(0, 1),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}

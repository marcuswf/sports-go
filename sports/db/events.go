package db

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"

	"git.neds.sh/matty/entain/sports/proto/sports"
)

// EventsRepo provides repository access to events.
type EventsRepo interface {
	// Init will initialise our events repository.
	Init() error

	// List will return a list of events.
	List(filter *sports.ListEventsRequestFilter, sortBy *sports.ListEventsRequestSortBy) ([]*sports.Event, error)

	// GetEvent will return a single event by its id
	Get(eventId int64) (*sports.Event, error)
}

type eventsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewEventsRepo creates a new events repository.
func NewEventsRepo(db *sql.DB) EventsRepo {
	return &eventsRepo{db: db}
}

// Init prepares the event repository dummy data.
func (e *eventsRepo) Init() error {
	var err error

	e.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy events.
		err = e.seed()
	})

	return err
}

func (e *eventsRepo) Get(eventId int64) (*sports.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getEventQueries()[eventsList]
	query += " WHERE id = ?"
	args = append(args, eventId)

	rows, err := e.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	events, err := e.scanEvents(rows)
	if len(events) == 0 {
		return nil, errors.New("Event not found")
	}
	return events[0], err
}

func (e *eventsRepo) List(filter *sports.ListEventsRequestFilter, sortBy *sports.ListEventsRequestSortBy) ([]*sports.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getEventQueries()[eventsList]

	query, args = e.applyFilter(query, filter)

	query, err = e.applySortBy(query, sortBy)
	if err != nil {
		return nil, err
	}

	rows, err := e.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return e.scanEvents(rows)
}

func (e *eventsRepo) applySortBy(query string, sortBy *sports.ListEventsRequestSortBy) (string, error) {
	if sortBy == nil {
		return query, nil
	}

	var column = strings.Trim(sortBy.PropertyName, " ")
	if column != "" {
		valid := regexp.MustCompile("^[A-Za-z0-9_]+$")
		if !valid.MatchString(sortBy.PropertyName) { //avoid sql injection
			return query, errors.New("Invalid sortBy / property_name")
		}
		query += " ORDER BY " + column
		if sortBy.Descending == true {
			query += " DESC"
		}
	}

	return query, nil
}

func (e *eventsRepo) applyFilter(query string, filter *sports.ListEventsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.Type) > 0 {

		clauses = append(clauses, "type IN ("+strings.Repeat("?,", len(filter.Type)-1)+"?)")

		for _, sType := range filter.Type {

			args = append(args, sType)
		}
	}

	if filter.Visible != nil {
		clauses = append(clauses, "visible=?")
		args = append(args, filter.Visible)
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

func (e *eventsRepo) scanEvents(
	rows *sql.Rows,
) ([]*sports.Event, error) {
	var events []*sports.Event

	for rows.Next() {
		var event sports.Event
		var advertisedStart time.Time
		var sType int

		if err := rows.Scan(&event.Id, &event.Name, &sType, &event.Location, &event.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		//cast from int to SportType
		event.Type = sports.SportType(sType)

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		event.AdvertisedStartTime = ts

		events = append(events, &event)
	}

	return events, nil
}

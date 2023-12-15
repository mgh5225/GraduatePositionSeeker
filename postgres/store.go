package postgres

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/mgh5225/gps"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		UniversityStore:    NewUniversityStore(db),
		FacultyStore:       NewFacultyStore(db),
		InterestStore:      NewInterestStore(db),
		FacultyMemberStore: NewFacultyMemberStore(db),
		TemplateStore:      NewTemplateStore(db),
	}, nil
}

type Store struct {
	gps.UniversityStore
	gps.FacultyStore
	gps.InterestStore
	gps.FacultyMemberStore
	gps.TemplateStore
}

package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	*UniversityStore
	*FacultyStore
	*InterestStore
	*FacultyMemberStore
	*TemplateStore
}

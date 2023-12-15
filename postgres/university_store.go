package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mgh5225/gps"
)

func NewUniversityStore(db *sqlx.DB) *UniversityStore {
	return &UniversityStore{
		DB: db,
	}
}

type UniversityStore struct {
	*sqlx.DB
}

func (s *UniversityStore) University(id uuid.UUID) (gps.University, error) {
	var u gps.University

	err := s.Get(&u, `select * from universities where id = $1`, id)

	if err != nil {
		return gps.University{}, fmt.Errorf("error getting university: %w", err)
	}

	return u, nil
}

func (s *UniversityStore) Universities() ([]gps.University, error) {
	var u []gps.University

	err := s.Select(&u, `select * from universities`)

	if err != nil {
		return []gps.University{}, fmt.Errorf("error getting universities: %w", err)
	}

	return u, nil
}

func (s *UniversityStore) CreateUniversity(u *gps.University) error {
	err := s.Get(u, "insert into universities values ($1, $2) returning *", u.ID, u.Name)

	if err != nil {
		return fmt.Errorf("error creating university: %w", err)
	}

	return nil
}

func (s *UniversityStore) UpdateUniversity(u *gps.University) error {
	err := s.Get(u, "update universities set name=$1 where id=$2 returning *", u.Name, u.ID)

	if err != nil {
		return fmt.Errorf("error updating university: %w", err)
	}

	return nil
}

func (s *UniversityStore) DeleteUniversity(id uuid.UUID) error {
	_, err := s.Exec(`delete from universities where id=$1`, id)

	if err != nil {
		return fmt.Errorf("error deleting university: %w", err)
	}

	return nil
}

package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mgh5225/gps"
)

func NewFacultyStore(db *sqlx.DB) *FacultyStore {
	return &FacultyStore{
		DB: db,
	}
}

type FacultyStore struct {
	*sqlx.DB
}

func (s *FacultyStore) Faculty(id uuid.UUID) (gps.Faculty, error) {
	var f gps.Faculty

	err := s.Get(&f, `select * from faculties where id = $1`, id)

	if err != nil {
		return gps.Faculty{}, fmt.Errorf("error getting faculty: %w", err)
	}

	return f, nil
}

func (s *FacultyStore) Faculties(u_id uuid.UUID) ([]gps.Faculty, error) {
	var f []gps.Faculty

	err := s.Select(&f, `select * from faculties where university_id=$1`, u_id)

	if err != nil {
		return []gps.Faculty{}, fmt.Errorf("error getting faculties: %w", err)
	}

	return f, nil
}

func (s *FacultyStore) CreateFaculty(f *gps.Faculty) error {
	err := s.Get(f, "insert into faculties values ($1, $2, $3) returning *", f.ID, f.UniversityID, f.Name)

	if err != nil {
		return fmt.Errorf("error creating faculty: %w", err)
	}

	return nil
}

func (s *FacultyStore) UpdateFaculty(f *gps.Faculty) error {
	err := s.Get(f, "update faculties set name=$1, university_id=$2 where id=$3 returning *", f.Name, f.UniversityID, f.ID)

	if err != nil {
		return fmt.Errorf("error updating faculty: %w", err)
	}

	return nil
}

func (s *FacultyStore) DeleteFaculty(id uuid.UUID) error {
	_, err := s.Exec(`delete from faculties where id=$1`, id)

	if err != nil {
		return fmt.Errorf("error deleting faculty: %w", err)
	}

	return nil
}

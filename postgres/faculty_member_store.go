package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mgh5225/gps"
)

func NewFacultyMemberStore(db *sqlx.DB) *FacultyMemberStore {
	return &FacultyMemberStore{
		DB: db,
	}
}

type FacultyMemberStore struct {
	*sqlx.DB
}

func (s *FacultyMemberStore) FacultyMember(id uuid.UUID) (gps.FacultyMember, error) {
	var fm gps.FacultyMember

	err := s.Get(&fm, `select * from faculty_members where id = $1`, id)

	if err != nil {
		return gps.FacultyMember{}, fmt.Errorf("error getting faculty member: %w", err)
	}

	return fm, nil
}

func (s *FacultyMemberStore) FacultyMembers(f_id uuid.UUID) ([]gps.FacultyMember, error) {
	var fm []gps.FacultyMember

	err := s.Get(&fm, `select * from faculty_members where faculty_id=$1`, f_id)

	if err != nil {
		return []gps.FacultyMember{}, fmt.Errorf("error getting faculty members: %w", err)
	}

	return fm, nil
}

func (s *FacultyMemberStore) CreateFacultyMember(fm *gps.FacultyMember) error {
	err := s.Get(fm, "insert into faculty_members values ($1, $2, $3, $4, $5) returning *", fm.ID, fm.FacultyID, fm.FirstName, fm.LastName, fm.Email)

	if err != nil {
		return fmt.Errorf("error creating faculty member: %w", err)
	}

	return nil
}

func (s *FacultyMemberStore) AddInterestToFacultyMember(fm_id uuid.UUID, i_id uuid.UUID) error {
	_, err := s.Exec("insert into faculty_member_interests values ($1, $2)", fm_id, i_id)

	if err != nil {
		return fmt.Errorf("error adding interest for faculty member: %w", err)
	}

	return nil
}

func (s *FacultyMemberStore) DeleteInterestFromFacultyMember(fm_id uuid.UUID, i_id uuid.UUID) error {
	_, err := s.Exec("delete from faculty_member_interests where faculty_id=$1 and interest_id=$2", fm_id, i_id)

	if err != nil {
		return fmt.Errorf("error deleting interest from faculty member: %w", err)
	}

	return nil
}

func (s *FacultyMemberStore) UpdateFacultyMember(fm *gps.FacultyMember) error {
	err := s.Get(fm,
		"update faculty_members set faculty_id=$1, first_name=$2, last_name=$3, email=$4 where id=$5 returning *",
		fm.FacultyID, fm.FirstName, fm.LastName, fm.Email, fm.ID)

	if err != nil {
		return fmt.Errorf("error updating faculty member: %w", err)
	}

	return nil
}

func (s *FacultyMemberStore) DeleteFacultyMember(id uuid.UUID) error {
	_, err := s.Exec(`delete from faculty_members where id=$1`, id)

	if err != nil {
		return fmt.Errorf("error deleting faculty member: %w", err)
	}

	return nil
}

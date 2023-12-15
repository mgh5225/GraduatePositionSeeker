package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mgh5225/gps"
)

func NewInterestStore(db *sqlx.DB) *InterestStore {
	return &InterestStore{
		DB: db,
	}
}

type InterestStore struct {
	*sqlx.DB
}

func (s *InterestStore) Interest(id uuid.UUID) (gps.Interest, error) {
	var i gps.Interest

	err := s.Get(&i, `select * from interests where id = $1`, id)

	if err != nil {
		return gps.Interest{}, fmt.Errorf("error getting interest: %w", err)
	}

	return i, nil
}

func (s *InterestStore) Interests() ([]gps.Interest, error) {
	var i []gps.Interest

	err := s.Get(&i, `select * from interests`)

	if err != nil {
		return []gps.Interest{}, fmt.Errorf("error getting interests: %w", err)
	}

	return i, nil
}

func (s *InterestStore) CreateInterest(i *gps.Interest) error {
	err := s.Get(i, "insert into interests values ($1, $2) returning *", i.ID, i.Name)

	if err != nil {
		return fmt.Errorf("error creating interest: %w", err)
	}

	return nil
}

func (s *InterestStore) UpdateInterest(i *gps.Interest) error {
	err := s.Get(i, "update interests set name=$1 where id=$2 returning *", i.Name, i.ID)

	if err != nil {
		return fmt.Errorf("error updating interest: %w", err)
	}

	return nil
}

func (s *InterestStore) DeleteInterest(id uuid.UUID) error {
	_, err := s.Exec(`delete from interests where id=$1`, id)

	if err != nil {
		return fmt.Errorf("error deleting interest: %w", err)
	}

	return nil
}

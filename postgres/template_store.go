package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mgh5225/gps"
)

func NewTemplateStore(db *sqlx.DB) *TemplateStore {
	return &TemplateStore{
		DB: db,
	}
}

type TemplateStore struct {
	*sqlx.DB
}

func (s *TemplateStore) Template(id uuid.UUID) (gps.Template, error) {
	var t gps.Template

	err := s.Get(&t, `select * from templates where id = $1`, id)

	if err != nil {
		return gps.Template{}, fmt.Errorf("error getting template: %w", err)
	}

	return t, nil
}

func (s *TemplateStore) Templates() ([]gps.Template, error) {
	var t []gps.Template

	err := s.Get(&t, `select * from templates`)

	if err != nil {
		return []gps.Template{}, fmt.Errorf("error getting templates: %w", err)
	}

	return t, nil
}

func (s *TemplateStore) CreateTemplate(t *gps.Template) error {
	err := s.Get(t, "insert into templates values ($1, $2, $3) returning *", t.ID, t.Subject, t.Body)

	if err != nil {
		return fmt.Errorf("error creating template: %w", err)
	}

	return nil
}

func (s *TemplateStore) UpdateTemplate(t *gps.Template) error {
	err := s.Get(t, "update templates set subject=$1, body=$2 where id=$3 returning *", t.Subject, t.Body, t.ID)

	if err != nil {
		return fmt.Errorf("error updating template: %w", err)
	}

	return nil
}

func (s *TemplateStore) DeleteTemplate(id uuid.UUID) error {
	_, err := s.Exec(`delete from templates where id=$1`, id)

	if err != nil {
		return fmt.Errorf("error deleting template: %w", err)
	}

	return nil
}

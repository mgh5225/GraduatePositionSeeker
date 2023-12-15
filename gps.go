package gps

import "github.com/google/uuid"

type University struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

type Faculty struct {
	ID           uuid.UUID `db:"id"`
	UniversityID uuid.UUID `db:"university_id"`
	Name         string    `db:"name"`
}

type Interest struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

type FacultyMemeber struct {
	ID        uuid.UUID `db:"id"`
	FacultyID uuid.UUID `db:"faculty_id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
}

type FacultyMemeberInterest struct {
	FacultyID  uuid.UUID `db:"faculty_id"`
	InterestID uuid.UUID `db:"interest_id"`
}

type Template struct {
	ID      uuid.UUID `db:"id"`
	Subject string    `db:"subject"`
	Body    string    `db:"body"`
}

type UniversityStore interface {
	University(id uuid.UUID) (University, error)
	Universities() ([]University, error)
	CreateUniversity(u *University) error
	UpdateUniversity(u *University) error
	DeleteUniversity(id uuid.UUID) error
}

type FacultyStore interface {
	Faculty(id uuid.UUID) (Faculty, error)
	Faculties(u_id uuid.UUID) ([]Faculty, error)
	CreateFaculty(f *Faculty) error
	UpdateFaculty(f *Faculty) error
	DeleteFaculty(id uuid.UUID) error
}

type InterestStore interface {
	Interest(id uuid.UUID) (Interest, error)
	Interests() ([]Interest, error)
	CreateInterest(i *Interest) error
	UpdateInterest(i *Interest) error
	DeleteInterest(id uuid.UUID) error
}

type FacultyMemeberStore interface {
	FacultyMemeber(id uuid.UUID) (FacultyMemeber, error)
	FacultyMemebers(f_id uuid.UUID) ([]FacultyMemeber, error)
	CreateFacultyMember(fm *FacultyMemeber) error
	UpdateFacultyMember(fm *FacultyMemeber) error
	DeleteFacultyMember(id uuid.UUID) error
}

type TemplateStore interface {
	Template(id uuid.UUID) (Template, error)
	Templates() ([]Template, error)
	CreateTemplate(t *Template) error
	UpdateTemplate(t *Template) error
	DeleteTemplate(id uuid.UUID) error
}

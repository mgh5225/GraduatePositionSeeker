package gps

import "github.com/google/uuid"

type University struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

type Faculty struct {
	ID           uuid.UUID `db:"id" json:"id"`
	UniversityID uuid.UUID `db:"university_id" json:"university_id"`
	Name         string    `db:"name" json:"name"`
}

type Interest struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

type FacultyMember struct {
	ID        uuid.UUID `db:"id" json:"id"`
	FacultyID uuid.UUID `db:"faculty_id" json:"faculty_id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email" json:"email"`
}

type FacultyMemberInterest struct {
	FacultyID  uuid.UUID `db:"faculty_id" json:"faculty_id"`
	InterestID uuid.UUID `db:"interest_id" json:"interest_id"`
}

type Template struct {
	ID      uuid.UUID `db:"id" json:"id"`
	Subject string    `db:"subject" json:"subject"`
	Body    string    `db:"body" json:"body"`
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

type FacultyMemberStore interface {
	FacultyMember(id uuid.UUID) (FacultyMember, error)
	FacultyMembers(f_id uuid.UUID) ([]FacultyMember, error)
	CreateFacultyMember(fm *FacultyMember) error
	AddInterestToFacultyMember(fm_id uuid.UUID, i_id uuid.UUID) error
	DeleteInterestFromFacultyMember(fm_id uuid.UUID, i_id uuid.UUID) error
	UpdateFacultyMember(fm *FacultyMember) error
	DeleteFacultyMember(id uuid.UUID) error
}

type TemplateStore interface {
	Template(id uuid.UUID) (Template, error)
	Templates() ([]Template, error)
	CreateTemplate(t *Template) error
	UpdateTemplate(t *Template) error
	DeleteTemplate(id uuid.UUID) error
}

type Store interface {
	UniversityStore
	FacultyStore
	InterestStore
	FacultyMemberStore
	TemplateStore
}

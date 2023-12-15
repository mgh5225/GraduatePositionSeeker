CREATE TABLE universities (
  id UUID PRIMARY KEY,
  name  TEXT NOT NULL
);

CREATE TABLE faculties (
  id UUID PRIMARY KEY,
  university_id UUID NOT NULL REFERENCES universities (id) ON DELETE CASCADE,
  name TEXT NOT NULL
);

CREATE TABLE interests (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE faculty_members (
  id UUID PRIMARY KEY,
  faculty_id UUID NOT NULL REFERENCES faculties (id) ON DELETE CASCADE,
  first_name  TEXT NOT NULL,
  last_name TEXT NOT NULL,
  email TEXT NOT NULL
);

CREATE TABLE faculty_member_interests (
  faculty_id UUID NOT NULL REFERENCES faculties (id),
  interest_id UUID NOT NULL REFERENCES interests (id)
);

CREATE TABLE templates (
  id UUID PRIMARY KEY,
  subject TEXT NOT NULL,
  body TEXT NOT NULL
);

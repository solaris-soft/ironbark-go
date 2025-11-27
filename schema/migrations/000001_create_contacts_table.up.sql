CREATE TABLE IF NOT EXISTS contacts (
  id         uuid      PRIMARY KEY DEFAULT gen_random_uuid(),
  first_name text      NOT NULL,
  last_name  text      NOT NULL,
  email      text      NOT NULL,
  phone      text      NOT NULL,
  address    text      NOT NULL,
  city       text      NOT NULL,
  state      text      NOT NULL,
  zip        text      NOT NULL,
  country    text      NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)
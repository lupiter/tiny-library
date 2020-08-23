BEGIN;
CREATE TABLE IF NOT EXISTS books (
  id serial primary key,
  title text,
  author text,
  isbn text,
  year text,
  publisher text,
  tags text[],
  max_loan_days int,
  location text,
  format text
);

CREATE TABLE IF NOT EXISTS patrons (
  id serial primary key,
  card_number text,
  name text,
  dob text,
  address text,
  active boolean,
  max_loan_days int
);

CREATE TABLE IF NOT EXISTS loans (
  id serial primary key,
  book integer,
  patron integer,
  lent timestamp with time zone,
  due_back timestamp with time zone,
  returned timestamp with time zone,
  constraint fk_book foreign key (book) references books(id),
  constraint fk_patron foreign key (patron) references patrons(id)
);
END;

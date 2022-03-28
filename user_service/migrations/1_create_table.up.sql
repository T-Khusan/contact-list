CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  id uuid primary key DEFAULT (uuid_generate_v4()),
  name varchar(150) not null,
  lastname varchar(150),
  password varchar(150) not null
);

CREATE TABLE IF NOT EXISTS contact (
  id uuid primary key DEFAULT (uuid_generate_v4()),
  name varchar(150) not null,
  phone varchar(150) not null,
  user_id uuid not null,
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) on delete cascade
);

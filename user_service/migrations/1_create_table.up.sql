CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE `user` (
  id uuid primary key DEFAULT (uuid_generate_v4()),
  name varchar(150) not null,
  lastname varchar(150),
  password varchar(150) not null
);

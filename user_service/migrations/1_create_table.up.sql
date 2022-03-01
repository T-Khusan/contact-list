CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE contact (
  id uuid primary key DEFAULT uuid_generate_v4(),
  name varchar(150) not null,
  phone varchar(150) not null
);

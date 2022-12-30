create table if not exists users (
  id text primary key,
  username text not null,
  password text not null
);
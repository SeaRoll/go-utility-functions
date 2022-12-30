-- dialect: postgresql
create table if not exists application_user (
  id VARCHAR(255) primary key,
  username VARCHAR(255) not null,
  password VARCHAR(255) not null,
  deleted boolean not null default false
);

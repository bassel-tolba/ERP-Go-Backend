-- +goose Up
create table companies (
  id serial primary key,
  name varchar(255) unique not null,
  created_at timestamptz not null default now()
);
-- +goose Down
drop table companies;

-- name: CreateCompany :one
insert into companies (name)
values ($1) returning *;


-- name: GetCompanyByID :one
select * from companies where id = $1;

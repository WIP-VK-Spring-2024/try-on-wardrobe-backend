// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package sqlc

import (
	"context"

	"try-on/internal/pkg/domain"
	"try-on/internal/pkg/utils"
)

const createUser = `-- name: CreateUser :one
insert into users(
    name,
    email,
    password,
    gender
) values ($1, $2, $3, $4)
returning id
`

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
	Gender   domain.Gender
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (utils.UUID, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Gender,
	)
	var id utils.UUID
	err := row.Scan(&id)
	return id, err
}

const getSubscribedToUsers = `-- name: GetSubscribedToUsers :many
select users.id, users.created_at, users.updated_at, users.name, users.email, users.password, users.gender, users.privacy, users.avatar
from users
join subs on subs.subscriber_id = $1
     and subs.user_id = users.id
`

func (q *Queries) GetSubscribedToUsers(ctx context.Context, subscriberID utils.UUID) ([]User, error) {
	rows, err := q.db.Query(ctx, getSubscribedToUsers, subscriberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Gender,
			&i.Privacy,
			&i.Avatar,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
select id, created_at, updated_at, name, email, password, gender, privacy, avatar from users
where lower(email) = lower($1)
`

func (q *Queries) GetUserByEmail(ctx context.Context, lower string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, lower)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Gender,
		&i.Privacy,
		&i.Avatar,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
select id, created_at, updated_at, name, email, password, gender, privacy, avatar from users
where id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id utils.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Gender,
		&i.Privacy,
		&i.Avatar,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
select id, created_at, updated_at, name, email, password, gender, privacy, avatar from users
where name = $1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Gender,
		&i.Privacy,
		&i.Avatar,
	)
	return i, err
}

const searchUsers = `-- name: SearchUsers :many
select users.id, users.created_at, users.updated_at, users.name, users.email, users.password, users.gender, users.privacy, users.avatar
from users
left join subs on subs.subscriber_id = $1
     and subs.user_id = users.id
where lower(name) like lower($3)
      and lower(name) > $4
      and subs.user_id is null
      and users.id <> $1
order by lower(name)
limit $2
`

type SearchUsersParams struct {
	SubscriberID utils.UUID
	Limit        int32
	Name         string
	Since        string
}

func (q *Queries) SearchUsers(ctx context.Context, arg SearchUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, searchUsers,
		arg.SubscriberID,
		arg.Limit,
		arg.Name,
		arg.Since,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Gender,
			&i.Privacy,
			&i.Avatar,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
update users
set name = case when $4::text = '' then name
                else $4::text end,
    gender = coalesce($2, gender),
    privacy = coalesce($3, privacy),
    avatar = case when $5::text = '' then avatar
                  else $5::text end,
    updated_at = now()
where id = $1
`

type UpdateUserParams struct {
	ID      utils.UUID
	Gender  domain.Gender
	Privacy domain.Privacy
	Name    string
	Avatar  string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Gender,
		arg.Privacy,
		arg.Name,
		arg.Avatar,
	)
	return err
}

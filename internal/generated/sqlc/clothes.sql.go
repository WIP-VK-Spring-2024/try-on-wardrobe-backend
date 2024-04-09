// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: clothes.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"try-on/internal/pkg/domain"
	"try-on/internal/pkg/utils"
)

const createClothes = `-- name: CreateClothes :one
insert into clothes(
    name,
    user_id,
    image
)
values ($1, $2, $3)
returning id
`

type CreateClothesParams struct {
	Name   string
	UserID utils.UUID
	Image  string
}

func (q *Queries) CreateClothes(ctx context.Context, arg CreateClothesParams) (utils.UUID, error) {
	row := q.db.QueryRow(ctx, createClothes, arg.Name, arg.UserID, arg.Image)
	var id utils.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteClothes = `-- name: DeleteClothes :exec
delete from clothes
where id = $1
`

func (q *Queries) DeleteClothes(ctx context.Context, id utils.UUID) error {
	_, err := q.db.Exec(ctx, deleteClothes, id)
	return err
}

const getClothesById = `-- name: GetClothesById :one
select
    clothes.id, clothes.created_at, clothes.updated_at, clothes.name, clothes.note, clothes.user_id, clothes.style_id, clothes.type_id, clothes.subtype_id, clothes.color, clothes.seasons, clothes.image, clothes.status,
    types.name as type,
    coalesce(types.tryonable, false) as tryonable,
    subtypes.name as subtype,
    styles.name as style,
    array_remove(array_agg(tags.name), null)::text[] as tags
from clothes
left join types on types.id = clothes.type_id
left join subtypes on subtypes.id = clothes.subtype_id
left join styles on styles.id = clothes.style_id
left join clothes_tags on clothes.id = clothes_tags.clothes_id
left join tags on clothes_tags.tag_id = tags.id
where clothes.id = $1
group by
    clothes.id,
    clothes.user_id,
    clothes.name,
    clothes.note,
    clothes.image,
    clothes.type_id,
    clothes.subtype_id,
    clothes.style_id,
    clothes.status,
    clothes.color,
    clothes.seasons,
    clothes.created_at,
    clothes.updated_at,
    tryonable,
    type,
    subtype,
    style
`

type GetClothesByIdRow struct {
	ID        utils.UUID
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	Name      string
	Note      pgtype.Text
	UserID    utils.UUID
	StyleID   utils.UUID
	TypeID    utils.UUID
	SubtypeID utils.UUID
	Color     pgtype.Text
	Seasons   []domain.Season
	Image     string
	Status    NullStatus
	Type      pgtype.Text
	Tryonable bool
	Subtype   pgtype.Text
	Style     pgtype.Text
	Tags      []string
}

func (q *Queries) GetClothesById(ctx context.Context, id utils.UUID) (GetClothesByIdRow, error) {
	row := q.db.QueryRow(ctx, getClothesById, id)
	var i GetClothesByIdRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Note,
		&i.UserID,
		&i.StyleID,
		&i.TypeID,
		&i.SubtypeID,
		&i.Color,
		&i.Seasons,
		&i.Image,
		&i.Status,
		&i.Type,
		&i.Tryonable,
		&i.Subtype,
		&i.Style,
		&i.Tags,
	)
	return i, err
}

const getClothesByUser = `-- name: GetClothesByUser :many
select
    clothes.id, clothes.created_at, clothes.updated_at, clothes.name, clothes.note, clothes.user_id, clothes.style_id, clothes.type_id, clothes.subtype_id, clothes.color, clothes.seasons, clothes.image, clothes.status,
    types.name as type,
    coalesce(types.tryonable, false) as tryonable,
    subtypes.name as subtype,
    styles.name as style,
    array_remove(array_agg(tags.name), null)::text[] as tags
from clothes
left join types on types.id = clothes.type_id
left join subtypes on subtypes.id = clothes.subtype_id
left join styles on styles.id = clothes.style_id
left join clothes_tags on clothes.id = clothes_tags.clothes_id
left join tags on clothes_tags.tag_id = tags.id
where clothes.user_id = $1
group by
    clothes.id,
    clothes.user_id,
    clothes.name,
    clothes.note,
    clothes.image,
    clothes.type_id,
    clothes.subtype_id,
    clothes.style_id,
    clothes.status,
    clothes.color,
    clothes.seasons,
    clothes.created_at,
    clothes.updated_at,
    tryonable,
    type,
    subtype,
    style
order by clothes.created_at desc
`

type GetClothesByUserRow struct {
	ID        utils.UUID
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	Name      string
	Note      pgtype.Text
	UserID    utils.UUID
	StyleID   utils.UUID
	TypeID    utils.UUID
	SubtypeID utils.UUID
	Color     pgtype.Text
	Seasons   []domain.Season
	Image     string
	Status    NullStatus
	Type      pgtype.Text
	Tryonable bool
	Subtype   pgtype.Text
	Style     pgtype.Text
	Tags      []string
}

func (q *Queries) GetClothesByUser(ctx context.Context, userID utils.UUID) ([]GetClothesByUserRow, error) {
	rows, err := q.db.Query(ctx, getClothesByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetClothesByUserRow
	for rows.Next() {
		var i GetClothesByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Note,
			&i.UserID,
			&i.StyleID,
			&i.TypeID,
			&i.SubtypeID,
			&i.Color,
			&i.Seasons,
			&i.Image,
			&i.Status,
			&i.Type,
			&i.Tryonable,
			&i.Subtype,
			&i.Style,
			&i.Tags,
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

const getClothesIdByOutfit = `-- name: GetClothesIdByOutfit :many
select c.id
from clothes c
join outfits o on o.transforms ? c.id
where o.id = $1
`

func (q *Queries) GetClothesIdByOutfit(ctx context.Context, id utils.UUID) ([]utils.UUID, error) {
	rows, err := q.db.Query(ctx, getClothesIdByOutfit, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []utils.UUID
	for rows.Next() {
		var id utils.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getClothesTryOnInfo = `-- name: GetClothesTryOnInfo :many
select
    clothes.id,
    try_on_type(type.name) as category
from clothes
join types on types.id = clothes.type_id
where clothes.id = any($1::uuid[]) and category <> ''
`

type GetClothesTryOnInfoRow struct {
	ID       utils.UUID
	Category string
}

func (q *Queries) GetClothesTryOnInfo(ctx context.Context, ids []utils.UUID) ([]GetClothesTryOnInfoRow, error) {
	rows, err := q.db.Query(ctx, getClothesTryOnInfo, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetClothesTryOnInfoRow
	for rows.Next() {
		var i GetClothesTryOnInfoRow
		if err := rows.Scan(&i.ID, &i.Category); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setClothesImage = `-- name: SetClothesImage :exec
update clothes
set image = $2
where id = $1
`

func (q *Queries) SetClothesImage(ctx context.Context, iD utils.UUID, image string) error {
	_, err := q.db.Exec(ctx, setClothesImage, iD, image)
	return err
}

const updateClothes = `-- name: UpdateClothes :exec
update clothes
set name = coalesce($2, name),
    note = coalesce($3, note),
    type_id = coalesce($4, type_id),
    subtype_id = coalesce($5, subtype_id),
    style_id = coalesce($6, style_id),
    color = coalesce($7, color),
    seasons = coalesce($8, seasons)::season[],
    updated_at = now()
where id = $1
`

type UpdateClothesParams struct {
	ID        utils.UUID
	Name      string
	Note      pgtype.Text
	TypeID    utils.UUID
	SubtypeID utils.UUID
	StyleID   utils.UUID
	Color     pgtype.Text
	Seasons   []domain.Season
}

func (q *Queries) UpdateClothes(ctx context.Context, arg UpdateClothesParams) error {
	_, err := q.db.Exec(ctx, updateClothes,
		arg.ID,
		arg.Name,
		arg.Note,
		arg.TypeID,
		arg.SubtypeID,
		arg.StyleID,
		arg.Color,
		arg.Seasons,
	)
	return err
}

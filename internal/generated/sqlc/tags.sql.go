// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tags.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"try-on/internal/pkg/utils"
)

const createClothesTagLinks = `-- name: CreateClothesTagLinks :exec
insert into clothes_tags (clothes_id, tag_id)
    select $1, id
    from tags where name = any($2::text[])
`

func (q *Queries) CreateClothesTagLinks(ctx context.Context, clothesID utils.UUID, tags []string) error {
	_, err := q.db.Exec(ctx, createClothesTagLinks, clothesID, tags)
	return err
}

const createOutfitTagLinks = `-- name: CreateOutfitTagLinks :exec
insert into outfits_tags(outfit_id, tag_id)
    select $1, id
    from tags where name = any($2::text[])
`

func (q *Queries) CreateOutfitTagLinks(ctx context.Context, outfitID utils.UUID, tags []string) error {
	_, err := q.db.Exec(ctx, createOutfitTagLinks, outfitID, tags)
	return err
}

const createTags = `-- name: CreateTags :exec
insert into tags (name) values (  
  unnest($1::varchar[])
) on conflict do nothing
`

func (q *Queries) CreateTags(ctx context.Context, names []string) error {
	_, err := q.db.Exec(ctx, createTags, names)
	return err
}

const createTagsWithEng = `-- name: CreateTagsWithEng :exec
insert into tags (name, eng_name) values ($1, $2::text)
`

func (q *Queries) CreateTagsWithEng(ctx context.Context, name string, column2 string) error {
	_, err := q.db.Exec(ctx, createTagsWithEng, name, column2)
	return err
}

const deleteClothesTagLinks = `-- name: DeleteClothesTagLinks :exec
delete from clothes_tags
where clothes_id = $1 and
    tag_id not in (
        select id from tags
        where name = any($2::text[])
    )
`

func (q *Queries) DeleteClothesTagLinks(ctx context.Context, clothesID utils.UUID, tags []string) error {
	_, err := q.db.Exec(ctx, deleteClothesTagLinks, clothesID, tags)
	return err
}

const deleteOutfitTagLinks = `-- name: DeleteOutfitTagLinks :exec
delete from outfits_tags
where outfit_id = $1 and
    tag_id not in (
        select id from tags
        where name = any($2::text[])
    )
`

func (q *Queries) DeleteOutfitTagLinks(ctx context.Context, outfitID utils.UUID, tags []string) error {
	_, err := q.db.Exec(ctx, deleteOutfitTagLinks, outfitID, tags)
	return err
}

const getNotCreatedTags = `-- name: GetNotCreatedTags :many
select name::text from unnest($1::text[]) as t(name)
    where name not in (select name from tags)
`

func (q *Queries) GetNotCreatedTags(ctx context.Context, names []string) ([]string, error) {
	rows, err := q.db.Query(ctx, getNotCreatedTags, names)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTagEngNames = `-- name: GetTagEngNames :many
select eng_name
from tags
where eng_name is not null
order by use_count desc
limit $1 offset $2
`

func (q *Queries) GetTagEngNames(ctx context.Context, limit int32, offset int32) ([]pgtype.Text, error) {
	rows, err := q.db.Query(ctx, getTagEngNames, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Text
	for rows.Next() {
		var eng_name pgtype.Text
		if err := rows.Scan(&eng_name); err != nil {
			return nil, err
		}
		items = append(items, eng_name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTags = `-- name: GetTags :many
select id, created_at, updated_at, name, use_count, eng_name
from tags
order by use_count desc
limit $1 offset $2
`

func (q *Queries) GetTags(ctx context.Context, limit int32, offset int32) ([]Tag, error) {
	rows, err := q.db.Query(ctx, getTags, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tag
	for rows.Next() {
		var i Tag
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.UseCount,
			&i.EngName,
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

const getTagsByEngName = `-- name: GetTagsByEngName :many
select tags.name
from tags
join unnest($1::text[])
    with ordinality t(eng_name, ord)
    on tags.eng_name = t.eng_name
order by t.ord
`

func (q *Queries) GetTagsByEngName(ctx context.Context, engNames []string) ([]string, error) {
	rows, err := q.db.Query(ctx, getTagsByEngName, engNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setTagEngName = `-- name: SetTagEngName :exec
update tags
set eng_name = $2::text
where name = $1
`

func (q *Queries) SetTagEngName(ctx context.Context, name string, engName string) error {
	_, err := q.db.Exec(ctx, setTagEngName, name, engName)
	return err
}

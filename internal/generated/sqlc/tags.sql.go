// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tags.sql

package sqlc

import (
	"context"

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

const createTags = `-- name: CreateTags :exec
insert into tags (name) values (  
  unnest($1::varchar[])
) on conflict do nothing
`

func (q *Queries) CreateTags(ctx context.Context, names []string) error {
	_, err := q.db.Exec(ctx, createTags, names)
	return err
}

const getTags = `-- name: GetTags :many
select
    id, 
    name,
    use_count
from tags
order by use_count desc
limit $1 offset $2
`

type GetTagsRow struct {
	ID       utils.UUID
	Name     string
	UseCount int32
}

func (q *Queries) GetTags(ctx context.Context, limit int32, offset int32) ([]GetTagsRow, error) {
	rows, err := q.db.Query(ctx, getTags, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTagsRow
	for rows.Next() {
		var i GetTagsRow
		if err := rows.Scan(&i.ID, &i.Name, &i.UseCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
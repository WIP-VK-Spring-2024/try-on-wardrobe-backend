-- name: CreateTags :exec
insert into tags (name) values (  
  unnest(sqlc.arg(names)::varchar[])
) on conflict do nothing;

-- name: CreateClothesTagLinks :exec
insert into clothes_tags (clothes_id, tag_id)
    select sqlc.arg(clothes_id), id
    from tags where name in (sqlc.arg(tags)::text[]);

-- name: DeleteClothesTagLinks :exec
delete from clothes_tags
where clothes_id = $1 and
    tag_id not in (
        select id from tags
        where name in (sqlc.arg(tags)::text[])
    );

-- name: CreateOutfitTagLinks :exec
insert into outfits_tags(outfit_id, tag_id)
    select sqlc.arg(outfit_id), id
    from tags where name in (sqlc.arg(tags)::text[]);

-- name: DeleteOutfitTagLinks :exec
delete from outfits_tags
where outfit_id = $1 and
    tag_id not in (
        select id from tags
        where name in (sqlc.arg(tags)::text[])
    );

-- name: GetTags :many
select *
from tags
order by use_count desc
limit $1 offset $2;

-- name: GetTagEngNames :many
select eng_name
from tags
where eng_name is not null
order by use_count desc
limit $1 offset $2;

-- name: GetTagsByEngName :many
select tags.name
from tags
join unnest(sqlc.arg(eng_names)::text[])
    with ordinality t(eng_name, ord)
    on tags.eng_name = t.eng_name
order by t.ord;

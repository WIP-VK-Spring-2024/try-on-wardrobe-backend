-- name: GetTypes :many
select
    types.*,
    array_agg(subtypes.id order by subtypes.name)::uuid[] as subtype_ids,
    array_agg(subtypes.name order by subtypes.name)::text[] as subtype_names
from types
left join subtypes on types.id = subtypes.type_id
group by
    types.id,
    types.name,
    types.created_at,
    types.updated_at
order by types.created_at, types.name;

-- name: GetSubtypes :many
select * from subtypes;

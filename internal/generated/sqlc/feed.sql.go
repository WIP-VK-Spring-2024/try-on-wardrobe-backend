// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feed.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"try-on/internal/pkg/utils"
)

const createComment = `-- name: CreateComment :one
insert into post_comments(post_id, user_id, body)
    values($1, $2, $3)
    returning id
`

type CreateCommentParams struct {
	PostID utils.UUID
	UserID utils.UUID
	Body   string
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (utils.UUID, error) {
	row := q.db.QueryRow(ctx, createComment, arg.PostID, arg.UserID, arg.Body)
	var id utils.UUID
	err := row.Scan(&id)
	return id, err
}

const getComments = `-- name: GetComments :many
select
    post_comments.id,
    post_comments.created_at,
    post_comments.updated_at,
    post_comments.user_id,
    post_comments.body,
    post_comments.rating,
    users.avatar as user_image,
    coalesce(post_comment_ratings.value, 0) as user_rating
from post_comments
join users on users.id = post_comments.user_id
left join post_comment_ratings on post_comment_ratings.user_id = $1
where post_comments.post_id = $2
  and post_comments.created_at < $4::timestamp
order by post_comments.created_at
limit $3
`

type GetCommentsParams struct {
	UserID utils.UUID
	PostID utils.UUID
	Limit  int32
	Since  utils.Time
}

type GetCommentsRow struct {
	ID         utils.UUID
	CreatedAt  utils.Time
	UpdatedAt  utils.Time
	UserID     utils.UUID
	Body       string
	Rating     int32
	UserImage  string
	UserRating int32
}

func (q *Queries) GetComments(ctx context.Context, arg GetCommentsParams) ([]GetCommentsRow, error) {
	rows, err := q.db.Query(ctx, getComments,
		arg.UserID,
		arg.PostID,
		arg.Limit,
		arg.Since,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentsRow
	for rows.Next() {
		var i GetCommentsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Body,
			&i.Rating,
			&i.UserImage,
			&i.UserRating,
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

const getLikedPosts = `-- name: GetLikedPosts :many
select
    posts.id,
    posts.created_at,
    posts.updated_at,
    posts.outfit_id,
    outfits.user_id,
    outfits.image as outfit_image,
    users.avatar as user_image,
    posts.rating,
    post_ratings.value as user_rating,
    coalesce(try_on_results.image, '') as try_on_image,
    coalesce(try_on_results.id, uuid_nil()) as try_on_id
from posts
join outfits on outfits.id = posts.outfit_id
join users on users.id = outfits.user_id
join post_ratings on post_ratings.user_id = $1
left join try_on_results on try_on_results.id = outfits.try_on_result_id
where posts.created_at < $3::timestamp
    and post_ratings.value = 1
order by posts.created_at
limit $2
`

type GetLikedPostsParams struct {
	UserID utils.UUID
	Limit  int32
	Since  utils.Time
}

type GetLikedPostsRow struct {
	ID          utils.UUID
	CreatedAt   utils.Time
	UpdatedAt   utils.Time
	OutfitID    utils.UUID
	UserID      utils.UUID
	OutfitImage pgtype.Text
	UserImage   string
	Rating      int32
	UserRating  int32
	TryOnImage  string
	TryOnID     utils.UUID
}

func (q *Queries) GetLikedPosts(ctx context.Context, arg GetLikedPostsParams) ([]GetLikedPostsRow, error) {
	rows, err := q.db.Query(ctx, getLikedPosts, arg.UserID, arg.Limit, arg.Since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLikedPostsRow
	for rows.Next() {
		var i GetLikedPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.OutfitID,
			&i.UserID,
			&i.OutfitImage,
			&i.UserImage,
			&i.Rating,
			&i.UserRating,
			&i.TryOnImage,
			&i.TryOnID,
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

const getPosts = `-- name: GetPosts :many
select
    posts.id,
    posts.created_at,
    posts.updated_at,
    posts.outfit_id,
    outfits.user_id,
    outfits.image as outfit_image,
    users.avatar as user_image,
    posts.rating,
    coalesce(post_ratings.value, 0) as user_rating,
    coalesce(try_on_results.image, '') as try_on_image,
    coalesce(try_on_results.id, uuid_nil()) as try_on_id
from posts
join outfits on outfits.id = posts.outfit_id
join users on users.id = outfits.user_id
left join post_ratings on post_ratings.user_id = $1
left join try_on_results on try_on_results.id = outfits.try_on_result_id
where posts.created_at < $3::timestamp
order by posts.created_at
limit $2
`

type GetPostsParams struct {
	UserID utils.UUID
	Limit  int32
	Since  utils.Time
}

type GetPostsRow struct {
	ID          utils.UUID
	CreatedAt   utils.Time
	UpdatedAt   utils.Time
	OutfitID    utils.UUID
	UserID      utils.UUID
	OutfitImage pgtype.Text
	UserImage   string
	Rating      int32
	UserRating  int32
	TryOnImage  string
	TryOnID     utils.UUID
}

func (q *Queries) GetPosts(ctx context.Context, arg GetPostsParams) ([]GetPostsRow, error) {
	rows, err := q.db.Query(ctx, getPosts, arg.UserID, arg.Limit, arg.Since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsRow
	for rows.Next() {
		var i GetPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.OutfitID,
			&i.UserID,
			&i.OutfitImage,
			&i.UserImage,
			&i.Rating,
			&i.UserRating,
			&i.TryOnImage,
			&i.TryOnID,
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

const getSubscriptionPosts = `-- name: GetSubscriptionPosts :many
select
    posts.id,
    posts.created_at,
    posts.updated_at,
    posts.outfit_id,
    outfits.user_id,
    outfits.image as outfit_image,
    users.avatar as user_image,
    posts.rating,
    coalesce(post_ratings.value, 0) as user_rating,
    coalesce(try_on_results.image, '') as try_on_image,
    coalesce(try_on_results.id, uuid_nil()) as try_on_id
from posts
join outfits on outfits.id = posts.outfit_id
join users on users.id = outfits.user_id
join subs on subs.user_id = users.id and subs.subscriber_id = $1
left join post_ratings on post_ratings.user_id = $1
left join try_on_results on try_on_results.id = outfits.try_on_result_id
where posts.created_at < $3::timestamp
order by posts.created_at
limit $2
`

type GetSubscriptionPostsParams struct {
	SubscriberID utils.UUID
	Limit        int32
	Since        utils.Time
}

type GetSubscriptionPostsRow struct {
	ID          utils.UUID
	CreatedAt   utils.Time
	UpdatedAt   utils.Time
	OutfitID    utils.UUID
	UserID      utils.UUID
	OutfitImage pgtype.Text
	UserImage   string
	Rating      int32
	UserRating  int32
	TryOnImage  string
	TryOnID     utils.UUID
}

func (q *Queries) GetSubscriptionPosts(ctx context.Context, arg GetSubscriptionPostsParams) ([]GetSubscriptionPostsRow, error) {
	rows, err := q.db.Query(ctx, getSubscriptionPosts, arg.SubscriberID, arg.Limit, arg.Since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSubscriptionPostsRow
	for rows.Next() {
		var i GetSubscriptionPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.OutfitID,
			&i.UserID,
			&i.OutfitImage,
			&i.UserImage,
			&i.Rating,
			&i.UserRating,
			&i.TryOnImage,
			&i.TryOnID,
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

const rateComment = `-- name: RateComment :exec
insert into post_comment_ratings(comment_id, user_id, value)
    values($1, $2, $3)
    on conflict (comment_id, user_id) do update
    set value = excluded.value
`

type RateCommentParams struct {
	CommentID utils.UUID
	UserID    utils.UUID
	Value     int32
}

func (q *Queries) RateComment(ctx context.Context, arg RateCommentParams) error {
	_, err := q.db.Exec(ctx, rateComment, arg.CommentID, arg.UserID, arg.Value)
	return err
}

const ratePost = `-- name: RatePost :exec
insert into post_ratings(post_id, user_id, value)
    values($1, $2, $3)
    on conflict (post_id, user_id) do update
    set value = excluded.value
`

type RatePostParams struct {
	PostID utils.UUID
	UserID utils.UUID
	Value  int32
}

func (q *Queries) RatePost(ctx context.Context, arg RatePostParams) error {
	_, err := q.db.Exec(ctx, ratePost, arg.PostID, arg.UserID, arg.Value)
	return err
}

const subscribe = `-- name: Subscribe :exec
insert into subs(subscriber_id, user_id)
    values($1, $2)
`

func (q *Queries) Subscribe(ctx context.Context, subscriberID utils.UUID, userID utils.UUID) error {
	_, err := q.db.Exec(ctx, subscribe, subscriberID, userID)
	return err
}

const unsubscribe = `-- name: Unsubscribe :exec
delete from subs
where subscriber_id = $1 and user_id = $2
`

func (q *Queries) Unsubscribe(ctx context.Context, subscriberID utils.UUID, userID utils.UUID) error {
	_, err := q.db.Exec(ctx, unsubscribe, subscriberID, userID)
	return err
}
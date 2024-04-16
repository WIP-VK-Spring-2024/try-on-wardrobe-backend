package domain

import (
	"try-on/internal/pkg/utils"
)

//easyjson:json
type Post struct {
	Model

	OutfitID    utils.UUID
	OutfitImage string

	TryOnID    utils.UUID
	TryOnImage string

	UserID    utils.UUID
	UserImage string

	Rating     int
	UserRating int
}

//easyjson:json
type Comment struct {
	Model
	CommentModel

	UserImage string

	Rating     int
	UserRating int
}

//easyjson:json
type CommentModel struct {
	UserID utils.UUID
	Body   string
}

type GetPostsOpts struct {
	RequestingUserID utils.UUID
	Limit            int32      `query:"limit"`
	Since            utils.Time `query:"since"`
}

type GetCommentsOpts struct {
	PostID           utils.UUID
	RequestingUserID utils.UUID
	Limit            int32      `query:"limit"`
	Since            utils.Time `query:"since"`
}

type FeedRepository interface {
	GetPosts(opts GetPostsOpts) ([]Post, error)
	GetLikedPosts(opts GetPostsOpts) ([]Post, error)
	GetSubscriptionPosts(opts GetPostsOpts) ([]Post, error)

	Subscribe(subscriberId, userId utils.UUID) error
	Unsubscribe(subscriberId, userId utils.UUID) error

	GetPost(postId utils.UUID) (*Post, error)
	GetComments(opts GetCommentsOpts) ([]Comment, error)

	RatePost(userId, postId utils.UUID, rating int) error
	RateComment(userId, commentId utils.UUID, rating int) error

	Comment(postId utils.UUID, comment CommentModel) error
}
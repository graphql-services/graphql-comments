package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateComment     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Comment, err error)
	UpdateComment     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Comment, err error)
	DeleteComment     func(ctx context.Context, r *GeneratedResolver, id string) (item *Comment, err error)
	DeleteAllComments func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryComment      func(ctx context.Context, r *GeneratedResolver, opts QueryCommentHandlerOptions) (*Comment, error)
	QueryComments     func(ctx context.Context, r *GeneratedResolver, opts QueryCommentsHandlerOptions) (*CommentResultType, error)

	CommentCreatedByUser func(ctx context.Context, r *GeneratedResolver, obj *Comment) (res *User, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateComment:     CreateCommentHandler,
		UpdateComment:     UpdateCommentHandler,
		DeleteComment:     DeleteCommentHandler,
		DeleteAllComments: DeleteAllCommentsHandler,
		QueryComment:      QueryCommentHandler,
		QueryComments:     QueryCommentsHandler,

		CommentCreatedByUser: CommentCreatedByUserHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.DB.Query()
	}
	return db
}
